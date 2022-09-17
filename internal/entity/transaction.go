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

type DrawAndDepositDTO struct {
	CallerId             int
	PlayerName           string
	WithDraw             int
	Deposit              int
	Currency             string
	TransactionRef       string
	GameRoundRef         string
	GameId               int
	Source               string
	Reason               string
	SessionId            string
	SessionAlternativeId string
	BonusId              string
	ChargeFreeRounds     int
}

type RollBackTransactionDTO struct {
	CallerId             int
	PlayerName           string
	TransactionRef       string
	GameId               int
	SessionId            string
	SessionAlternativeId string
	RoundId              string
}
