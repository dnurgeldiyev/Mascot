package entity

type Transaction struct {
	PlayerName     string
	WithDraw       int
	Deposit        int
	TransactionRef string
	SessionID      string
	RollBackStatus bool
}

type TransactionKey struct {
	TransactionRef string
	SessionID      string
}