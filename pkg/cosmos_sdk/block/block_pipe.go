package cosmos

import (
	"context"
	codec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	blockcodec "github.com/mapofzones/cosmos-watcher/pkg/codec"
	crawler "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/crawler"
	parsing "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/parsing"
	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	websocket "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/websocket"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	"log"
)

// BlockStream returns channel of ordered blocks
func BlockStream(ctx context.Context, client *http.HTTP, startHeight int64) <-chan watcher.Block {
	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 1")
	return toInterface(
		ctx, decodedStream(
			ctx, ordered(
				ctx, crawlerToWebsocket(
					ctx, client, startHeight), startHeight)))
}

// decodedStream is used to convert decode proto-encoded blocks with cosmos-sdk codec
// as well as to convert the data in blocks to messages specified by the app
func decodedStream(ctx context.Context, stream <-chan block.Block) <-chan block.ProcessedBlock {
	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 2")
	processedStream := make(chan block.ProcessedBlock)
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	blockcodec.RegisterInterfacesAndImpls(interfaceRegistry)

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 2")
	go func() {
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 3")
		defer close(processedStream)
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 4")
		for {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 5")
			select {
			case block := <-stream:
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 6")
				decoded, err := parsing.DecodeBlock(cdc, block)
				if err != nil {
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 7")
					log.Fatal(err)
					return
				}
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 8")
				select {
				case processedStream <- decoded:
				case <-ctx.Done():
					return
				}
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 9")
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 10")
				return
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 11")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 12")
	return processedStream
}

func toInterface(ctx context.Context, stream <-chan block.ProcessedBlock) <-chan watcher.Block {
	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 13")
	out := make(chan watcher.Block)

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 14")
	go func() {
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 15")
		defer close(out)
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 16")
		for {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 17")
			select {
			case block := <-stream:
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 18")
				select {
				case out <- block:
				case <-ctx.Done():
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 19")
					return
				}
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 20")
				return
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 21")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 22")
	return out
}

// normalizedStream is used to format websocket data correspondingly to our exported block structure
func normalizedStream(ctx context.Context, blockWithTxs <-chan block.WithTxs) <-chan block.Block {
	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 23")
	blockStream := make(chan block.Block)

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 24")
	go func() {
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 25")
		defer close(blockStream)
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 26")
		for {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 27")
			select {
			case Block, ok := <-blockWithTxs:
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 28")
				if !ok {
					return
				}
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 29")
				select {
				case blockStream <- block.Normalize(Block):
				case <-ctx.Done():
				}
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 30")
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 31")
				return
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 32")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 33")
	return blockStream
}

func ordered(ctx context.Context, stream <-chan block.Block, startHeight int64) <-chan block.Block {
	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 34")
	orderedStream := make(chan block.Block)

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 35")
	go func() {
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 36")
		defer close(orderedStream)
		expectedHeight := startHeight

		blocks := map[int64]block.Block{}
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 37")
		for {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 38")
			for {
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 39")
				// go through the local cache
				if block, ok := blocks[expectedHeight]; ok {
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 40")
					select {
					case orderedStream <- block:
						log.Println("pkg.cosmos_sdk.block.block_pipe.go - 41")
						delete(blocks, expectedHeight)
						expectedHeight++
						continue
					case <-ctx.Done():
						log.Println("pkg.cosmos_sdk.block.block_pipe.go - 42")
						return
					}
				}
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 43")
				break
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 44")
			select {
			case block, ok := <-stream:
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 45")
				// return since the source channel is closed
				if !ok {
					return
				}

				// should not happen, but worth checking in case of a bug somewhere in the pipe
				if block.Height < expectedHeight {
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 46")
					break
				}

				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 47")
				// if we got the block we expected to get
				if block.Height == expectedHeight {
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 48")
					select {
					case orderedStream <- block:
						log.Println("pkg.cosmos_sdk.block.block_pipe.go - 49")
						expectedHeight++
					case <-ctx.Done():
						log.Println("pkg.cosmos_sdk.block.block_pipe.go - 50")
						return
					}
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 51")
					// if this is not the block we need
				} else {
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 52")
					// put it in local cache
					blocks[block.Height] = block
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 53")
					// if this happened, might as well kill the watcher
					// because blocks cache should never be this large
					if len(blocks) > 16 {
						log.Println("pkg.cosmos_sdk.block.block_pipe.go - 54")
						log.Fatal("ordered blocks buffer overflow")
						return
					}
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 55")
				}
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 56")
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 57")
				return
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 58")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 59")
	return orderedStream
}

// crawler to websocket check blockchain height and decides whether to run crawler or we can listen on websocket
// if we caught up with blockchain, this method switches from crawler stream to websocket stream
func crawlerToWebsocket(ctx context.Context, client *http.HTTP, startHeight int64) <-chan block.Block {
	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 60")
	blockStream := make(chan block.Block)

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 61")
	go func() {
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 62")
		defer close(blockStream)
		status, err := client.Status(ctx)
		if err != nil {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 63")
			log.Println(err)
			return
		}
		currentHeight := status.SyncInfo.LatestBlockHeight
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 64")
		// if we have some catching up to do
		for currentHeight-startHeight > 1 {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 65")
			blocks := crawler.BlockRange(ctx, client, startHeight, currentHeight)
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 66")
			for block := range blocks {
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 67")
				select {
				case blockStream <- block:
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 68")
					startHeight++
				case <-ctx.Done():
					log.Println("pkg.cosmos_sdk.block.block_pipe.go - 69")
					return
				}
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 70")
			}

			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 71")
			// update current latest blockchain height to know if we need another iteration of this loop
			status, err := client.Status(ctx)
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 72")
			if err != nil {
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 73")
				log.Println(err)
				return
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 74")
			currentHeight = status.SyncInfo.LatestBlockHeight
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 75")
		}

		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 76")
		// alias variable to make stuff less confusing
		// since we synced up, start height == lastProcessedBlock
		lastProcessedBlock := startHeight

		// if we reached this point in the function, we can start listening on the websocket in realtime
		websocketStream := normalizedStream(ctx, websocket.BlockStream(ctx, client))

		// get one block from websocket to know how many blocks we have missed while switching from
		// crawler to websocket block stream
		var pivotBlock block.Block
		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 77")
		select {
		case block, ok := <-websocketStream:
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 78")
			if !ok {
				return
			}
			pivotBlock = block
		case <-ctx.Done():
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 79")
			return
		}

		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 80")
		// last time we catch up with node
		for block := range crawler.BlockRange(ctx, client, lastProcessedBlock, pivotBlock.Height) {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 81")
			select {
			case blockStream <- block:
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 82")
				return
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 83")
		}

		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 84")
		// send pivot block
		select {
		case blockStream <- pivotBlock:
		case <-ctx.Done():
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 85")
			return
		}

		log.Println("pkg.cosmos_sdk.block.block_pipe.go - 86")
		// switch to websocket at last
		for {
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 87")
			select {
			case blockStream <- <-websocketStream:
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.block_pipe.go - 88")
				return
			}
			log.Println("pkg.cosmos_sdk.block.block_pipe.go - 89")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.block_pipe.go - 90")
	return blockStream
}
