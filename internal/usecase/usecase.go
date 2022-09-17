package usecase

import (
	"dovran/mascot/internal/entity"
	"dovran/mascot/internal/usecase/balance"
	"dovran/mascot/internal/usecase/transaction"
)

type UCase struct {
	Balance     *balance.StorageBalance
	Transaction *transaction.StorageTransaction
}

func NewUCase(b *balance.StorageBalance, t *transaction.StorageTransaction) *UCase {

	return &UCase{
		Balance:     b,
		Transaction: t,
	}

}

func (uc UCase) AddPlayerWithBalance(playerName string, balance int) (err error) {

	err = uc.Balance.AddPlayerAndBalance(playerName, balance)

	if err != nil {
		err = ErrBadRequest
		return
	}

	return
}

func (uc UCase) GetBalance(playerName string) (item *entity.Balance, err error) {

	item = &entity.Balance{}

	item, err = uc.Balance.GetBalance(playerName)

	if err != nil {
		err = ErrNotFound
		return
	}

	return
}

func (uc UCase) WithDrawAndDeposit(drawDeposit entity.DrawAndDepositDTO) (item *entity.Balance, err error) {

	item = &entity.Balance{}

	getBalance, err := uc.Balance.GetBalance(drawDeposit.PlayerName)
	if err != nil {
		err = ErrNotFound
		return
	}

	_, err = uc.Transaction.GetTransaction(drawDeposit.TransactionRef, drawDeposit.SessionId)
	if err == nil {
		err = ErrBadRequest
		return
	}

	if drawDeposit.WithDraw > getBalance.Balance {
		err = ErrBadRequest
		return
	}

	_, err = uc.Transaction.AddTransaction(drawDeposit.PlayerName, drawDeposit.TransactionRef, drawDeposit.SessionId, drawDeposit.WithDraw, drawDeposit.Deposit, false)

	if err != nil {
		err = ErrForbidden
		return
	}

	item, err = uc.Balance.UpdateBalance(getBalance.PlayerName, getBalance.Balance-drawDeposit.WithDraw+drawDeposit.Deposit, drawDeposit.ChargeFreeRounds)
	if err != nil {
		err = ErrInternalServerError

		if _, err = uc.Transaction.UpdateStatusTransaction(drawDeposit.TransactionRef, drawDeposit.SessionId, true); err != nil {
			err = ErrInternalServerError
			return
		}

		return
	}

	return
}

func (uc UCase) RollBackTransaction(rollback entity.RollBackTransactionDTO) (err error) {

	getTransaction, err := uc.Transaction.GetTransaction(rollback.TransactionRef, rollback.SessionId)

	if err != nil {
		_, err = uc.Transaction.AddTransaction(rollback.PlayerName, rollback.TransactionRef, rollback.SessionId, 0, 0, true)
		if err != nil {
			err = ErrForbidden
			return
		}
		return
	}

	if getTransaction.RollBackStatus {
		err = ErrConflict
		return
	}

	_, err = uc.Transaction.UpdateStatusTransaction(rollback.TransactionRef, rollback.SessionId, true)
	if err != nil {
		err = ErrForbidden
		return
	}

	getBalance, err := uc.Balance.GetBalance(getTransaction.PlayerName)
	if err != nil {
		err = ErrForbidden
		return
	}

	_, err = uc.Balance.UpdateBalance(getTransaction.PlayerName, (getBalance.Balance-getTransaction.Deposit)+getTransaction.WithDraw, getBalance.FreeRoundsLeft)
	if err != nil {
		err = ErrForbidden
		return
	}

	return
}
