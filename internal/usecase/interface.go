package usecase

import (
	"dovran/mascot/internal/entity"
)

type (
	UseCase interface {
		GetBalance(playerName string) (item *entity.Balance, err error)
		WithDrawAndDeposit(drawDeposit entity.DrawAndDepositDTO) (item *entity.Balance, err error)
		RollBackTransaction(rollback entity.RollBackTransactionDTO) (err error)
	}
	Balance interface {
		GetBalance(playerName string) (item *entity.Balance, err error)
		UpdateBalance(playerName string, balance, freeRoundsLeft int) (item *entity.Balance, err error)
	}
	Transaction interface {
		GetTransaction(transactionRef, sessionId string) (item *entity.Transaction, err error)
		AddTransaction(playerName, transactionRef, sessionId string, withdraw, deposit int, rollBackStatus bool) (item entity.Transaction, err error)
		UpdateStatusTransaction(transactionRef, sessionId string, transactionStatus bool) (item *entity.Transaction, err error)
	}
)
