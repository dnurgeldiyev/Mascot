package entity

type Balance struct {
	PlayerName     string
	Balance        int
	FreeRoundsLeft int
}

type BalanceDTO struct {
	CallerId             int    `json:"callerId"`
	PlayerName           string `json:"playerName"`
	Currency             string `json:"currency"`
	GameId               int    `json:"gameId,omitempty"`
	SessionId            string `json:"sessionId,omitempty"`
	SessionAlternativeId string `json:"sessionAlternativeId,omitempty"`
	BonusId              string `json:"bonusId,omitempty"`
}
