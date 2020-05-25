package watcher

// interface compile-time checks
var (
	_ Message = Transaction{}
	_ Message = Transfer{}
)

// Transaction is a special type of message because it can
// contain multiple other messages inside of it
// this is done in order to know what messages belong to what tx
// or to know even that they have happened outside of tx
type Transaction struct {
	Hash     string
	Accepted bool
	Messages []Message
}

func (t Transaction) Type() string {
	return "transaction"
}

type Transfer struct {
	Sender, Recipient string
	Amount            []struct {
		Amount int64
		Coin   string
	}
}

func (t Transfer) Type() string {
	return "transfer"
}

// IBC protocol messages

// client-related messages
// https://github.com/cosmos/ics/tree/master/spec/ics-002-client-semantics

type CreateClient struct {
	// client ID
	ClientID string
	// type of lite-client
	ClientType string
	// chain-ID of the blockchain to which client belongs to
	ChainID string
}

func (t CreateClient) Type() string {
	return "create_client"
}

// connection-related messages
// https://github.com/cosmos/ics/tree/master/spec/ics-003-connection-semantics
// There are 4 parts of connection handshake between chain A and B:
//  openInit(A), openTry(B), openAck(A), openConfirm(B)

//  Initialize connection with chain some other chain
// this could be openInit on chain A or openTry on chain B
type CreateConnection struct {
	ConnectionID string
	ClientID     string
}

func (t CreateConnection) Type() string {
	return "create_connection"
}

// channel-related messages
// https://github.com/cosmos/ics/tree/master/spec/ics-004-channel-and-packet-semantics
// There are 4 parts of channel handshake between chain A and B:
//  openInit(A), openTry(B), openAck(A), openConfirm(B)
// There are also two methods responsible for closing channels:
// closeInit(A), closeConfirm(B)

// this message covers openInit and openTry
type CreateChannel struct {
	ChannelID    string
	ConnectionID string
	PortID       string
}

func (t CreateChannel) Type() string {
	return "create_channel"
}

// this message covers openAck and openConfirm
type OpenChannel struct {
	ChannelID string
}

func (t OpenChannel) Type() string {
	return "open_channel"
}

// this message covers closeInit and closeConfirm
type CloseChannel struct {
	ChannelID string
}

func (t CloseChannel) Type() string {
	return "close_channel"
}

// messages related to IBC token transfer
// https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer
// there are two actions performed on open channels responsible for ibc token transfer:
// SendPacket(transfer.MsgTransfer) and receivePacket(channel.MsgPackget)

// same as regular transfer, but with channel specified
type IBCTransfer struct {
	ChannelID         string
	Sender, Recipient string
	Amount            []struct {
		Amount int64
		Coin   string
	}
}

func (t IBCTransfer) Type() string {
	return "ibc_transfer"
}
