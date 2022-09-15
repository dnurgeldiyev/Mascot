package usecase

import (
	"dovran/mascot/internal/controller/v1"
	"dovran/mascot/internal/usecase/balance"
	"dovran/mascot/internal/usecase/transaction"
	"github.com/google/uuid"
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

func (uc UCase) GetBalance(playerName string) (item *v1.BalanceResponse, err error) {

	item = &v1.BalanceResponse{}

	getBalance, err := uc.Balance.GetBalance(playerName)
	if err != nil {
		err = ErrNotFound
		return
	}

	item.Balance = getBalance.Balance
	item.FreeRoundsLeft = getBalance.FreeRoundsLeft

	return
}

func (uc UCase) WithDrawAndDeposit(drawDeposit v1.DrawAndDeposit) (item *v1.DrawAndDepositResponse, err error) {

	item = &v1.DrawAndDepositResponse{}

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

	updateBalance, err := uc.Balance.UpdateBalance(getBalance.PlayerName, getBalance.Balance-drawDeposit.WithDraw+drawDeposit.Deposit, drawDeposit.ChargeFreeRounds)
	if err != nil {
		err = ErrInternalServerError

		if _, err = uc.Transaction.UpdateStatusTransaction(drawDeposit.TransactionRef, drawDeposit.SessionId, true); err != nil {
			err = ErrInternalServerError
			return
		}

		return
	}

	item.NewBalance = updateBalance.Balance
	item.TransactionId = uuid.NewString()
	item.FreeRoundsLeft = updateBalance.FreeRoundsLeft

	return
}

func (uc UCase) RollBackTransaction(rollback v1.RollBackTransaction) (err error) {

	getTransaction, err := uc.Transaction.GetTransaction(rollback.TransactionRef, rollback.SessionId)
	if err != nil {
		err = ErrForbidden
		return
	}

	if getTransaction == nil {
		_, err = uc.Transaction.AddTransaction(rollback.PlayerName, rollback.TransactionRef, rollback.SessionId, 0, 0, true)
		if err != nil {
			err = ErrForbidden
			return
		}
		return
	}

	_, err = uc.Transaction.UpdateStatusTransaction(rollback.TransactionRef, rollback.SessionId, true)
	if err != nil {
		err = ErrForbidden
		return
	}

	return
}
