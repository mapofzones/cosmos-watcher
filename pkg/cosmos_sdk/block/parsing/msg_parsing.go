package cosmos

import (
	"encoding/json"
	"errors"
	types6 "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	types5 "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	types2 "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	types3 "github.com/cosmos/cosmos-sdk/x/ibc/core/03-connection/types"
	types4 "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	types7 "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/07-tendermint/types"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"log"
)

func parseMsg(msg sdk.Msg, results []*types6.ResponseDeliverTx) ([]watcher.Message, error) {
	log.Println("parseMsg")
	switch msg := msg.(type) {

	// send creation
	case *types.MsgSend:
		return []watcher.Message{
			watcher.Transfer{
				Sender: (*msg).FromAddress,
				Recipient: (*msg).ToAddress,
				Amount: sdkCoinsToStruct((*msg).Amount),
			},
		}, nil

	// client creation
	case *types2.MsgCreateClient:
		value := msg.ClientState.GetCachedValue()
		chainId := value.(*types7.ClientState).ChainId
		clientId := ""
		clientId = ParseClientIDFromResults(results, clientId)
		messages := []watcher.Message{
			watcher.CreateClient{
				ChainID:    chainId,
				ClientID:   clientId,
				ClientType: "",
			},
		}
		if clientId == "" {
			return messages, errors.New("clientId not found")
		}
		return messages, nil

	// connection creation
	case *types3.MsgConnectionOpenInit:
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: msg.Counterparty.ConnectionId,
				ClientID:     msg.ClientId,
			},
		}, nil

	case *types3.MsgConnectionOpenTry:
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: msg.Counterparty.ConnectionId,
				ClientID:     msg.ClientId,
			},
		}, nil

	// channel creation
	case *types4.MsgChannelOpenInit:
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    msg.Channel.Counterparty.ChannelId,
				PortID:       msg.PortId,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}, nil

	case *types4.MsgChannelOpenTry:
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    msg.Channel.Counterparty.ChannelId,
				PortID:       msg.PortId,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}, nil

	// channel opening/closing
	case *types4.MsgChannelOpenAck:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	case *types4.MsgChannelOpenConfirm:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	case *types4.MsgChannelCloseInit:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	case *types4.MsgChannelCloseConfirm:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	// ibc transfer messages
	case *types5.MsgTransfer:
		return []watcher.Message{
			watcher.IBCTransfer{
				ChannelID: msg.SourceChannel,
				Sender:    msg.Sender,
				Recipient: msg.Receiver,
				Amount:    sdkCoinsToStruct([]sdk.Coin{msg.Token}),
				Source:    true,
			},
		}, nil

	case *types4.MsgRecvPacket:
		data := transfer.FungibleTokenPacketData{}
		err := json.Unmarshal(msg.Packet.Data, &data)
		if err != nil {
			return nil, err
		}
		return []watcher.Message{
			watcher.IBCTransfer{
				ChannelID: msg.Packet.DestinationChannel,
				Sender:    data.Sender,
				Recipient: data.Receiver,
				Amount:    packetToStruct(data),
				Source:    false,
			},
		}, nil
	}

	return []watcher.Message{}, nil
}

func ParseClientIDFromResults(results []*types6.ResponseDeliverTx, clientId string) string {
	for _, res := range results {
		for _, event := range res.Events {
			if event.Type == clienttypes.EventTypeCreateClient {
				for _, attr := range event.Attributes {
					if string(attr.Key) == clienttypes.AttributeKeyClientID {
						clientId = string(attr.Value)
						log.Println("client attr.Value:", string(attr.Value))
					}
				}
			}
		}
	}
	return clientId
}

func sdkCoinsToStruct(data []sdk.Coin) []struct {
	Amount int64
	Coin   string
} {
	transformed := make([]struct {
		Amount int64
		Coin   string
	}, len(data))

	for i, sdkCoin := range data {
		transformed[i] = struct {
			Amount int64
			Coin   string
		}{
			Coin:   sdkCoin.Denom,
			Amount: sdkCoin.Amount.Int64(),
		}
	}
	return transformed
}

func packetToStruct(data transfer.FungibleTokenPacketData) []struct {
	Amount int64
	Coin   string
} {
	transformed := make([]struct {
		Amount int64
		Coin   string
	}, 1)

	transformed[0] = struct {
		Amount int64
		Coin   string
	}{
		Coin:   data.Denom,
		Amount: int64(data.Amount),
	}
	return transformed
}