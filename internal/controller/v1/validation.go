package v1

import "errors"

func (b PlayerWithBalance) isValid() error {
	if len(b.PlayerName) < 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	switch b.Currency {
	case EuroCurrency:
	default:
		return ErrIllegalCurrencyCode
	}

	if b.Balance < 0 {
		return errors.New("player new balance can't less than zero")
	}

	return nil
}

func (b Balance) isValid() error {

	if b.CallerId < 1 {
		return errors.New("CallerId can't be less 1")
	}

	if len(b.PlayerName) < 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	switch b.Currency {
	case EuroCurrency:
	default:
		return ErrIllegalCurrencyCode
	}

	return nil
}
func (b DrawAndDeposit) isValid() error {

	if b.CallerId < 1 {
		return errors.New("CallerId can't be less 1")
	}

	if len(b.PlayerName) < 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	if b.WithDraw < 0 {
		return ErrNegativeWithdrawalCode
	}

	if b.Deposit < 0 {
		return ErrNegativeDepositCode
	}

	switch b.Currency {
	case EuroCurrency:
	default:
		return ErrIllegalCurrencyCode
	}

	if b.TransactionRef == "" {
		return errors.New("TransactionRef can't be empty string")
	}

	return nil
}
func (b RollBackTransaction) isValid() error {

	if b.CallerId < 1 {
		return errors.New("CallerId can't be less 1")
	}

	if len(b.PlayerName) < 1 {
		return errors.New("PlayerName length can't be less 2")
	}

	if b.TransactionRef == "" {
		return errors.New("TransactionRef can't be empty string")
	}

	return nil
}
