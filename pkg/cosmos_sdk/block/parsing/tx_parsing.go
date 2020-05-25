package cosmos

import (
	"encoding/json"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	connection "github.com/cosmos/cosmos-sdk/x/ibc/03-connection"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	tendermint "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
)

func parseTx(tx auth.StdTx) watcher.Message {
	return watcher.Transaction{}
}

func parseMsg(msg sdk.Msg) []watcher.Message {
	switch msg := msg.(type) {

	// client creation
	case tendermint.MsgCreateClient:
		return []watcher.Message{
			watcher.CreateClient{
				ChainID:    msg.Header.ChainID,
				ClientID:   msg.ClientID,
				ClientType: msg.GetClientType(),
			},
		}

	// connection creation
	case connection.MsgConnectionOpenInit:
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: msg.ConnectionID,
				ClientID:     msg.ClientID,
			},
		}
	case connection.MsgConnectionOpenTry:
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: msg.ConnectionID,
				ClientID:     msg.ClientID,
			},
		}

	// channel creation
	case channel.MsgChannelOpenInit:
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    msg.ChannelID,
				PortID:       msg.PortID,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}
	case channel.MsgChannelOpenTry:
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    msg.ChannelID,
				PortID:       msg.PortID,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}

	// channel opening/closing
	case channel.MsgChannelOpenAck:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelID,
			},
		}
	case channel.MsgChannelOpenConfirm:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelID,
			},
		}
	case channel.MsgChannelCloseInit:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelID,
			},
		}
	case channel.MsgChannelCloseConfirm:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelID,
			},
		}

	// ibc transfer messages
	case transfer.MsgTransfer:
		return []watcher.Message{
			watcher.IBCTransfer{
				ChannelID: msg.SourceChannel,
				Sender:    msg.Sender.String(),
				Recipient: msg.Receiver,
				Amount:    sdkCoinsToStruct(msg.Amount),
			},
		}
	case channel.MsgPacket:
		data := transfer.FungibleTokenPacketData{}
		err := json.Unmarshal(msg.Packet.Data, &data)
		if err != nil {
			log.Println(err)
			break
		}
		return []watcher.Message{
			watcher.IBCTransfer{
				ChannelID: msg.Packet.SourceChannel,
				Sender:    data.Sender,
				Recipient: data.Receiver,
				Amount:    sdkCoinsToStruct(data.Amount),
			},
		}

	}

	return []watcher.Message{}
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
