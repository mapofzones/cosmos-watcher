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
