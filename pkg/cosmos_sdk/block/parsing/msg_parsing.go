package cosmos

import (
	"encoding/json"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	solomachine "github.com/cosmos/ibc-go/v3/modules/light-clients/06-solomachine/types"
	types7 "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	types6 "github.com/tendermint/tendermint/abci/types"
	"math/big"

	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"log"
)

type attributeFiler struct {
	key   string
	value string
}

func parseMsg(msg sdk.Msg, txResult *types6.ResponseDeliverTx, errCode uint32) ([]watcher.Message, error) {
	switch msg := msg.(type) {

	// send creation
	case *types.MsgSend:
		return []watcher.Message{
			watcher.Transfer{
				Sender:    (*msg).FromAddress,
				Recipient: (*msg).ToAddress,
				Amount:    sdkCoinsToStruct((*msg).Amount),
			},
		}, nil

	// client creation
	case *clienttypes.MsgCreateClient:
		value := msg.ClientState.GetCachedValue()
		var chainId string
		switch client := value.(type) {
		case *types7.ClientState:
			chainId = client.ChainId
		case *solomachine.ClientState:
			pubKey, _ := client.ConsensusState.GetPubKey()
			chainId = pubKey.String()
		}
		clientId := ""
		clientId = ParseClientIDFromResults(txResult, clientId)
		messages := []watcher.Message{
			watcher.CreateClient{
				ChainID:    chainId,
				ClientID:   clientId,
				ClientType: "",
			},
		}
		if clientId == "" && errCode == 0 {
			return messages, errors.New("clientID not found")
		}
		return messages, nil

	// connection creation
	case *connectiontypes.MsgConnectionOpenInit:
		if errCode != 0 {
			return []watcher.Message{}, nil
		}
		expectedEvents := []string{connectiontypes.EventTypeConnectionOpenInit}
		attributeKeys := []string{connectiontypes.AttributeKeyConnectionID}
		attrFiler := attributeFiler{clienttypes.AttributeKeyClientID, msg.ClientId}
		connectionIDs := ParseIDsFromResults(txResult, expectedEvents, attributeKeys,
			attrFiler, attributeFiler{}, attributeFiler{}, attributeFiler{})
		if (len(connectionIDs) != 1 || len(connectionIDs[0]) == 0) && errCode == 0 {
			return nil, errors.New("connectionID not found")
		}
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: connectionIDs[0],
				ClientID:     msg.ClientId,
			},
		}, nil

	case *connectiontypes.MsgConnectionOpenTry:
		if errCode != 0 {
			return []watcher.Message{}, nil
		}
		expectedEvents := []string{connectiontypes.EventTypeConnectionOpenTry}
		attributeKeys := []string{connectiontypes.AttributeKeyConnectionID}
		attrFiler := attributeFiler{clienttypes.AttributeKeyClientID, msg.ClientId}
		connectionIDs := ParseIDsFromResults(txResult, expectedEvents, attributeKeys,
			attrFiler, attributeFiler{}, attributeFiler{}, attributeFiler{})
		if (len(connectionIDs) != 1 || len(connectionIDs[0]) == 0) && errCode == 0 {
			return nil, errors.New("connectionID not found")
		}
		return []watcher.Message{
			watcher.CreateConnection{
				ConnectionID: connectionIDs[0],
				ClientID:     msg.ClientId,
			},
		}, nil

	// channel creation
	case *channeltypes.MsgChannelOpenInit:
		if errCode != 0 {
			return []watcher.Message{}, nil
		}
		expectedEvents := []string{channeltypes.EventTypeChannelOpenInit}
		attributeKeys := []string{channeltypes.AttributeKeyChannelID}

		attrFiler1 := attributeFiler{channeltypes.AttributeKeyPortID, msg.PortId}
		attrFiler2 := attributeFiler{"counterparty_port_id", msg.Channel.Counterparty.PortId}
		attrFiler3 := attributeFiler{"counterparty_channel_id", msg.Channel.Counterparty.ChannelId}
		attrFiler4 := attributeFiler{connectiontypes.AttributeKeyConnectionID, msg.Channel.ConnectionHops[0]}

		channelIDs := ParseIDsFromResults(txResult, expectedEvents, attributeKeys,
			attrFiler1, attrFiler2, attrFiler3, attrFiler4)
		if (len(channelIDs) != 1 || len(channelIDs[0]) == 0) && errCode == 0 {
			return nil, errors.New("MsgChannelOpenInit channelID not found")
		}
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    channelIDs[0],
				PortID:       msg.PortId,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}, nil

	case *channeltypes.MsgChannelOpenTry:
		if errCode != 0 {
			return []watcher.Message{}, nil
		}
		expectedEvents := []string{channeltypes.EventTypeChannelOpenTry}
		attributeKeys := []string{channeltypes.AttributeKeyChannelID}
		attrFiler1 := attributeFiler{channeltypes.AttributeKeyPortID, msg.PortId}
		attrFiler2 := attributeFiler{"counterparty_port_id", msg.Channel.Counterparty.PortId}
		attrFiler3 := attributeFiler{"counterparty_channel_id", msg.Channel.Counterparty.ChannelId}
		attrFiler4 := attributeFiler{connectiontypes.AttributeKeyConnectionID, msg.Channel.ConnectionHops[0]}

		channelIDs := ParseIDsFromResults(txResult, expectedEvents, attributeKeys, attrFiler1, attrFiler2, attrFiler3, attrFiler4)
		if len(channelIDs) != 1 || len(channelIDs[0]) == 0 {
			return nil, errors.New("MsgChannelOpenTry channelID not found")
		}
		return []watcher.Message{
			watcher.CreateChannel{
				ChannelID:    channelIDs[0],
				PortID:       msg.PortId,
				ConnectionID: msg.Channel.ConnectionHops[0],
			},
		}, nil

	// channel opening/closing
	case *channeltypes.MsgChannelOpenAck:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	case *channeltypes.MsgChannelOpenConfirm:
		return []watcher.Message{
			watcher.OpenChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	case *channeltypes.MsgChannelCloseInit:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	case *channeltypes.MsgChannelCloseConfirm:
		return []watcher.Message{
			watcher.CloseChannel{
				ChannelID: msg.ChannelId,
			},
		}, nil

	// ibc transfer messages
	case *transfer.MsgTransfer:
		return []watcher.Message{
			watcher.IBCTransfer{
				ChannelID: msg.SourceChannel,
				Sender:    msg.Sender,
				Recipient: msg.Receiver,
				Amount:    sdkCoinsToStruct([]sdk.Coin{msg.Token}),
				Source:    true,
			},
		}, nil

	case *channeltypes.MsgRecvPacket:
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

func ParseClientIDFromResults(txResult *types6.ResponseDeliverTx, clientId string) string {
	if txResult != nil {
		for _, event := range txResult.Events {
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

func ParseIDsFromResults(txResult *types6.ResponseDeliverTx, expectedEvents []string, attributeKeys []string, attrFiler1 attributeFiler, attrFiler2 attributeFiler, attrFiler3 attributeFiler, attrFiler4 attributeFiler) []string {
	var attributesValues []string
	if txResult != nil {
		for _, event := range txResult.Events {
			for _, expected := range expectedEvents {
				if event.Type == expected {
					isCorrect := false
					if attrFiler1 == (attributeFiler{}) {
						isCorrect = true
					} else {
						isCorrect1 := false
						isCorrect2 := false
						isCorrect3 := false
						isCorrect4 := false
						for _, attr := range event.Attributes {
							if attrFiler1 != (attributeFiler{}) &&
								attrFiler1.key == string(attr.Key) &&
								attrFiler1.value == string(attr.Value) {
								isCorrect1 = true
							}
							if attrFiler2 != (attributeFiler{}) &&
								attrFiler2.key == string(attr.Key) &&
								attrFiler2.value == string(attr.Value) {
								isCorrect2 = true
							}
							if attrFiler3 != (attributeFiler{}) &&
								attrFiler3.key == string(attr.Key) &&
								attrFiler3.value == string(attr.Value) {
								isCorrect3 = true
							}
							if attrFiler4 != (attributeFiler{}) &&
								attrFiler4.key == string(attr.Key) &&
								attrFiler4.value == string(attr.Value) {
								isCorrect4 = true
							}
							if isCorrect1 == true && isCorrect2 == true && isCorrect3 == true && isCorrect4 == true {
								isCorrect = true
							}
						}
					}
					for _, attr := range event.Attributes {
						var values []string
						for _, expectedKey := range attributeKeys {
							if string(attr.Key) == expectedKey && isCorrect {
								values = append(values, string(attr.Value))
								log.Println(expectedKey, " attr.Value:", string(attr.Value))
							}
						}
						for _, value := range values {
							attributesValues = append(attributesValues, value)
						}
					}
				}
			}
		}
	}
	return attributesValues
}

func sdkCoinsToStruct(data []sdk.Coin) []struct {
	Amount *big.Int
	Coin   string
} {
	transformed := make([]struct {
		Amount *big.Int
		Coin   string
	}, len(data))

	for i, sdkCoin := range data {
		n := new(big.Int)
		base := 10
		amount, ok := n.SetString(sdkCoin.Amount.String(), base)
		if !ok {
			log.Fatalf("Cannot unmarshal %s to bigint: error", sdkCoin.Amount)
		}

		transformed[i] = struct {
			Amount *big.Int
			Coin   string
		}{
			Coin:   sdkCoin.Denom,
			Amount: amount,
		}
	}
	return transformed
}

func packetToStruct(data transfer.FungibleTokenPacketData) []struct {
	Amount *big.Int
	Coin   string
} {
	transformed := make([]struct {
		Amount *big.Int
		Coin   string
	}, 1)

	n := new(big.Int)
	base := 10
	amountString := "0"
	if len(data.Amount) > 0 {
		amountString = data.Amount
	}
	amount, ok := n.SetString(amountString, base)
	if !ok {
		log.Fatalf("Cannot unmarshal %s to bigint: error", data.Amount)
	}

	transformed[0] = struct {
		Amount *big.Int
		Coin   string
	}{
		Coin:   data.Denom,
		Amount: amount,
	}
	return transformed
}
