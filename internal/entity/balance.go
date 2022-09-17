package entity

type Balance struct {
	PlayerName     string
	Balance        int
	FreeRoundsLeft int
}

type BalanceDTO struct {
	CallerId             int
	PlayerName           string
	Currency             string
	GameId               int
	SessionId            string
	SessionAlternativeId string
	BonusId              string
}
