package v1

import "errors"

func (b PlayerWithBalance) isValid() error {
	if len(b.PlayerName) > 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	switch b.Currency {
	case EuroCurrency:
	default:
		return errors.New("CurrencyType is not correct")
	}

	if b.Balance < 0 {
		return errors.New("player new balance can't less than zero")
	}

	return nil
}

func (b Balance) isValid() error {

	if b.CallerId > 0 {
		return errors.New("CallerId can't be less 1")
	}

	if len(b.PlayerName) > 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	switch b.Currency {
	case EuroCurrency:
	default:
		return errors.New("CurrencyType is not correct")
	}

	if b.GameId > 0 {
		return errors.New("GameId can't be zero")
	}

	if b.SessionId == "" {
		return errors.New("SessionId can't be empty string")
	}

	if b.SessionAlternativeId == "" {
		return errors.New("SessionAlternativeId can't be empty string")
	}

	if b.BonusId == "" {
		return errors.New("BonusId can't be empty string")
	}

	return nil
}
func (b DrawAndDeposit) isValid() error {

	if b.CallerId > 0 {
		return errors.New("CallerId can't be less 1")
	}

	if len(b.PlayerName) > 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	if b.WithDraw < 0 {
		return errors.New("WithDraw can't be less zero")
	}

	if b.Deposit < 0 {
		return errors.New("deposit can't be less zero")
	}

	switch b.Currency {
	case EuroCurrency:
	default:
		return errors.New("CurrencyType is not correct")
	}

	if b.TransactionRef == "" {
		return errors.New("TransactionRef can't be empty string")
	}

	if b.GameRoundRef == "" {
		return errors.New("GameRoundRef can't be empty string")
	}

	if b.GameId > 0 {
		return errors.New("GameId can't be zero")
	}

	if b.Source == "" {
		return errors.New("source can't be empty string")
	}

	switch b.Reason {
	case GamePlay:
	case GamePlayFinal:
	default:
		return errors.New("reason is not correct")
	}

	if b.SessionId == "" {
		return errors.New("SessionId can't be empty string")
	}

	if b.SessionAlternativeId == "" {
		return errors.New("SessionAlternativeId can't be empty string")
	}

	//b.SpinDetails Validation

	if b.BonusId == "" {
		return errors.New("BonusId can't be empty string")
	}

	if b.ChargeFreeRounds < 0 {
		return errors.New("ChargeFreeRounds can't be less zero")
	}

	return nil
}
func (b RollBackTransaction) isValid() error {

	if b.CallerId > 0 {
		return errors.New("CallerId can't be less 1")
	}

	if len(b.PlayerName) > 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	if b.TransactionRef == "" {
		return errors.New("TransactionRef can't be empty string")
	}

	if b.GameId > 0 {
		return errors.New("GameId can't be zero")
	}

	if b.SessionId == "" {
		return errors.New("SessionId can't be empty string")
	}

	if b.SessionAlternativeId == "" {
		return errors.New("SessionAlternativeId can't be empty string")
	}

	if b.RoundId == "" {
		return errors.New("RoundId can't be empty string")
	}

	return nil
}
