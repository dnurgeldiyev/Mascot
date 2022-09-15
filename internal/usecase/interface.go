package usecase

import (
	"dovran/mascot/internal/controller/v1"
	"dovran/mascot/internal/entity"
)

type (
	UseCase interface {
		GetBalance(playerName string) (item *v1.BalanceResponse, err error)
		WithDrawAndDeposit(drawDeposit v1.DrawAndDeposit) (item *v1.DrawAndDepositResponse, err error)
		RollBackTransaction(rollback v1.RollBackTransaction) (err error)
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
