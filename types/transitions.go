package watcher

// interface compile-time checks
var (
	_ StateTransition = Transaction{}
	_ StateTransition = Transfer{}
)

// Transaction is a special type of transition because it can
// contain multiple other transitions inside of it
type Transaction struct {
	Hash        string
	Accepted    bool
	Transitions []StateTransition
}

func (t Transaction) Type() string {
	return "transaction"
}

type Transfer struct {
	Sender, Recipient string
	Amount            int64
	Coin              string
}

func (t Transfer) Type() string {
	return "transfer"
}

// IBC protocol transitions

// client-related transitions
// https://github.com/cosmos/ics/tree/master/spec/ics-002-client-semantics

type CreateClient struct {
	// client ID
	ID string
	// type of lite-client
	ClientType string
	// chain-ID of the blockchain to which client belongs to
	ChainID string
}

func (t CreateClient) Type() string {
	return "create_client"
}

// connection-related transitions
// https://github.com/cosmos/ics/tree/master/spec/ics-003-connection-semantics
// There are 4 parts of connection handshake between chain A and B:
//  openInit(A), openTry(B), openAck(A), openConfirm(B)

//  Initialize connection with chain some other chain
// this could be openInit on chain A or openTry on chain B
type CreateConnection struct {
	ID       string
	ClientID string
}

func (t CreateConnection) Type() string {
	return "create_connection"
}

// channel-related transitions
// https://github.com/cosmos/ics/tree/master/spec/ics-004-channel-and-packet-semantics
// There are 4 parts of channel handshake between chain A and B:
//  openInit(A), openTry(B), openAck(A), openConfirm(B)
// There are also two methods responsible for closing channels:
// closeInit(A), closeConfirm(B)

// this transition covers openInit and openTry
type CreateChannel struct {
	ID           string
	connectionID string
}

func (t CreateChannel) Type() string {
	return "create_channel"
}

// this transition covers openAck and openConfirm
type OpenChannel struct {
	ID string
}

func (t OpenChannel) Type() string {
	return "open_channel"
}

// this transition covers closeInit and closeConfirm
type CloseChannel struct {
	ID string
}

func (t CloseChannel) Type() string {
	return "close_channel"
}
