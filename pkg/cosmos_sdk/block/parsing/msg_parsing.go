package cosmos

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connection "github.com/cosmos/cosmos-sdk/x/ibc/03-connection"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	tendermint "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
)

func parseMsg(msg sdk.Msg) ([]watcher.Message, error) {
	switch msg := msg.(type) {

	// client creation
	case tendermint.MsgCreateClient:
		return []watcher.Message{
			watcher.CreateClient{
				ChainID:    msg.Header.ChainID,
				ClientID:   msg.ClientID,
				ClientType: msg.GetClientType(),
			},
		}, nil

	// connection creation
	case connection.MsgConnectionOpenInit:
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: msg.ConnectionID,
				ClientID:     msg.ClientID,
			},
		}, nil

	case connection.MsgConnectionOpenTry:
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: msg.ConnectionID,
				ClientID:     msg.ClientID,
			},
		}, nil

	// channel creation
	case channel.MsgChannelOpenInit:
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    msg.ChannelID,
				PortID:       msg.PortID,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}, nil

	case channel.MsgChannelOpenTry:
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    msg.ChannelID,
				PortID:       msg.PortID,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}, nil

	// channel opening/closing
	case channel.MsgChannelOpenAck:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelID,
			},
		}, nil

	case channel.MsgChannelOpenConfirm:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelID,
			},
		}, nil

	case channel.MsgChannelCloseInit:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelID,
			},
		}, nil

	case channel.MsgChannelCloseConfirm:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelID,
			},
		}, nil

	// ibc transfer messages
	case transfer.MsgTransfer:
		return []watcher.Message{
			watcher.IBCTransfer{
				ChannelID: msg.SourceChannel,
				Sender:    msg.Sender.String(),
				Recipient: msg.Receiver,
				Amount:    sdkCoinsToStruct(msg.Amount),
				Source:    true,
			},
		}, nil
	case channel.MsgPacket:
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
				Amount:    sdkCoinsToStruct(data.Amount),
				Source:    false,
			},
		}, nil

	}

	return []watcher.Message{}, nil
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
