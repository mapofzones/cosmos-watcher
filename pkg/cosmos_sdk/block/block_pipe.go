package cosmos

import (
	"context"
	codec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	types3 "github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	types4 "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/07-tendermint/types"
	crawler "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/crawler"
	parsing "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/parsing"
	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	websocket "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/websocket"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	"log"
	blockcodec "github.com/mapofzones/cosmos-watcher/pkg/codec"
)

// BlockStream returns channel of ordered blocks
func BlockStream(ctx context.Context, client *http.HTTP, startHeight int64) <-chan watcher.Block {
	return toInterface(
		ctx, decodedStream(
			ctx, ordered(
				ctx, crawlerToWebsocket(
					ctx, client, startHeight), startHeight)))
}

// decodedStream is used to convert decode proto-encoded blocks with cosmos-sdk codec
// as well as to convert the data in blocks to messages specified by the app
func decodedStream(ctx context.Context, stream <-chan block.Block) <-chan block.ProcessedBlock {
	processedStream := make(chan block.ProcessedBlock)
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	simapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
	blockcodec.RegisterMessagesImplementations(interfaceRegistry)

	//interfaceRegistry.RegisterInterface("tendermint.crypto.PubKey", (*crypto.PubKey)(nil))
	//interfaceRegistry.RegisterImplementations((*crypto.PubKey)(nil), &ed25519.PubKey{})
	//interfaceRegistry.RegisterImplementations((*crypto.PubKey)(nil), &secp256k1.PubKey{})
	//interfaceRegistry.RegisterImplementations((*crypto.PubKey)(nil), &multisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterImplementations((*types3.ClientState)(nil), &types4.ClientState{})
	interfaceRegistry.RegisterImplementations((*types3.ConsensusState)(nil), &types4.ConsensusState{})
	interfaceRegistry.RegisterImplementations((*types3.Header)(nil), &types4.Header{})
	interfaceRegistry.RegisterImplementations((*types3.Misbehaviour)(nil), &types4.Misbehaviour{})

	go func() {
		defer close(processedStream)
		for {
			select {
			case block := <-stream:
				decoded, err := parsing.DecodeBlock(cdc, block)
				if err != nil {
					log.Fatal(err)
					return
				}
				select {
				case processedStream <- decoded:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return processedStream
}

func toInterface(ctx context.Context, stream <-chan block.ProcessedBlock) <-chan watcher.Block {
	out := make(chan watcher.Block)

	go func() {
		defer close(out)
		for {
			select {
			case block := <-stream:
				select {
				case out <- block:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

// normalizedStream is used to format websocket data correspondingly to our exported block structure
func normalizedStream(ctx context.Context, blockWithTxs <-chan block.WithTxs) <-chan block.Block {
	blockStream := make(chan block.Block)

	go func() {
		defer close(blockStream)
		for {
			select {
			case Block, ok := <-blockWithTxs:
				if !ok {
					return
				}
				select {
				case blockStream <- block.Normalize(Block):
				case <-ctx.Done():
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockStream
}

func ordered(ctx context.Context, stream <-chan block.Block, startHeight int64) <-chan block.Block {
	orderedStream := make(chan block.Block)

	go func() {
		defer close(orderedStream)
		expectedHeight := startHeight

		blocks := map[int64]block.Block{}
		for {
			for {
				// go through the local cache
				if block, ok := blocks[expectedHeight]; ok {
					select {
					case orderedStream <- block:
						delete(blocks, expectedHeight)
						expectedHeight++
						continue
					case <-ctx.Done():
						return
					}
				}
				break
			}
			select {
			case block, ok := <-stream:
				// return since the source channel is closed
				if !ok {
					return
				}

				// should not happen, but worth checking in case of a bug somewhere in the pipe
				if block.Height < expectedHeight {
					break
				}

				// if we got the block we expected to get
				if block.Height == expectedHeight {
					select {
					case orderedStream <- block:
						expectedHeight++
					case <-ctx.Done():
						return
					}
					// if this is not the block we need
				} else {
					// put it in local cache
					blocks[block.Height] = block
					// if this happened, might as well kill the watcher
					// because blocks cache should never be this large
					if len(blocks) > 16 {
						log.Fatal("ordered blocks buffer overflow")
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return orderedStream
}

// crawler to websocket check blockchain height and decides whether to run crawler or we can listen on websocket
// if we caught up with blockchain, this method switches from crawler stream to websocket stream
func crawlerToWebsocket(ctx context.Context, client *http.HTTP, startHeight int64) <-chan block.Block {
	blockStream := make(chan block.Block)

	go func() {
		defer close(blockStream)
		status, err := client.Status(ctx)
		if err != nil {
			log.Println(err)
			return
		}
		currentHeight := status.SyncInfo.LatestBlockHeight
		// if we have some catching up to do
		for currentHeight-startHeight > 1 {
			blocks := crawler.BlockRange(ctx, client, startHeight, currentHeight)
			for block := range blocks {
				select {
				case blockStream <- block:
					startHeight++
				case <-ctx.Done():
					return
				}
			}

			// update current latest blockchain height to know if we need another iteration of this loop
			status, err := client.Status(ctx)
			if err != nil {
				log.Println(err)
				return
			}
			currentHeight = status.SyncInfo.LatestBlockHeight
		}

		// alias variable to make stuff less confusing
		// since we synced up, start height == lastProcessedBlock
		lastProcessedBlock := startHeight

		// if we reached this point in the function, we can start listening on the websocket in realtime
		websocketStream := normalizedStream(ctx, websocket.BlockStream(ctx, client))

		// get one block from websocket to know how many blocks we have missed while switching from
		// crawler to websocket block stream
		var pivotBlock block.Block
		select {
		case block, ok := <-websocketStream:
			if !ok {
				return
			}
			pivotBlock = block
		case <-ctx.Done():
			return
		}

		// last time we catch up with node
		for block := range crawler.BlockRange(ctx, client, lastProcessedBlock, pivotBlock.Height) {
			select {
			case blockStream <- block:
			case <-ctx.Done():
				return
			}
		}

		// send pivot block
		select {
		case blockStream <- pivotBlock:
		case <-ctx.Done():
			return
		}

		// switch to websocket at last
		for {
			select {
			case blockStream <- <-websocketStream:
			case <-ctx.Done():
				return
			}
		}

	}()

	return blockStream
}
