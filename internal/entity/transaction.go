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
	CallerId             int    `json:"callerId"`
	PlayerName           string `json:"playerName"`
	WithDraw             int    `json:"withDraw"`
	Deposit              int    `json:"deposit"`
	Currency             string `json:"currency"`
	TransactionRef       string `json:"transactionRef"`
	GameRoundRef         string `json:"gameRoundRef,omitempty"`
	GameId               int    `json:"gameId,omitempty"`
	Source               string `json:"source,omitempty"`
	Reason               string `json:"reason,omitempty"`
	SessionId            string `json:"sessionId,omitempty"`
	SessionAlternativeId string `json:"sessionAlternativeId,omitempty"`
	BonusId              string `json:"bonusId,omitempty"`
	ChargeFreeRounds     int    `json:"chargeFreeRounds,omitempty"`
}

type RollBackTransactionDTO struct {
	CallerId             int    `json:"callerId"`
	PlayerName           string `json:"playerName"`
	TransactionRef       string `json:"transactionRef"`
	GameId               int    `json:"gameId,omitempty"`
	SessionId            string `json:"sessionId,omitempty"`
	SessionAlternativeId string `json:"sessionAlternativeId,omitempty"`
	RoundId              string `json:"roundId,omitempty"`
}
