package v1

type Balance struct {
	CallerId             int          `json:"callerId"`
	PlayerName           string       `json:"playerName"`
	Currency             CurrencyType `json:"currency"`
	GameId               int          `json:"gameId,omitempty"`
	SessionId            string       `json:"sessionId,omitempty"`
	SessionAlternativeId string       `json:"sessionAlternativeId,omitempty"`
	BonusId              string       `json:"bonusId,omitempty"`
}

type BalanceResponse struct {
	Balance        int `json:"balance"`
	FreeRoundsLeft int `json:"freeRoundsLeft,omitempty"`
}

type DrawAndDeposit struct {
	CallerId             int          `json:"callerId"`
	PlayerName           string       `json:"playerName"`
	WithDraw             int          `json:"withDraw"`
	Deposit              int          `json:"deposit"`
	Currency             CurrencyType `json:"currency"`
	TransactionRef       string       `json:"transactionRef"`
	GameRoundRef         string       `json:"gameRoundRef,omitempty"`
	GameId               int          `json:"gameId,omitempty"`
	Source               string       `json:"source,omitempty"`
	Reason               Reason       `json:"reason,omitempty"`
	SessionId            string       `json:"sessionId,omitempty"`
	SessionAlternativeId string       `json:"sessionAlternativeId,omitempty"`
	SpinDetails          SpinDetail   `json:"spinDetails,omitempty"`
	BonusId              string       `json:"bonusId,omitempty"`
	ChargeFreeRounds     int          `json:"chargeFreeRounds,omitempty"`
}

type DrawAndDepositResponse struct {
	NewBalance     int    `json:"newBalance"`
	TransactionId  string `json:"transactionId"`
	FreeRoundsLeft int    `json:"freeRoundsLeft,omitempty"`
}

type RollBackTransaction struct {
	CallerId             int    `json:"callerId"`
	PlayerName           string `json:"playerName"`
	TransactionRef       string `json:"transactionRef"`
	GameId               int    `json:"gameId,omitempty"`
	SessionId            string `json:"sessionId,omitempty"`
	SessionAlternativeId string `json:"sessionAlternativeId,omitempty"`
	RoundId              string `json:"roundId,omitempty"`
}

type CurrencyType string

const (
	EuroCurrency CurrencyType = "EUR"
)

type SpinDetail struct {
	BetType string `json:"betType,omitempty"`
	WinType string `json:"winType,omitempty"`
}

type Reason string

const (
	GamePlay      Reason = "GAME_PLAY"
	GamePlayFinal Reason = "GAME_PLAY_FINAL"
)
