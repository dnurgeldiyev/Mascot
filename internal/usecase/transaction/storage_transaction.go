package transaction

import (
	"dovran/mascot/internal/entity"
	"errors"
	"sync"
)

type StorageTransaction struct {
	m   map[entity.TransactionKey]entity.Transaction
	mtx sync.RWMutex
}

func NewStorageTransaction() *StorageTransaction {
	var s StorageTransaction
	s.m = make(map[entity.TransactionKey]entity.Transaction)

	return &s
}

func (s *StorageTransaction) GetTransaction(transactionRef, sessionId string) (item *entity.Transaction, err error) {

	item = &entity.Transaction{}

	key := entity.TransactionKey{
		TransactionRef: transactionRef,
		SessionID:      sessionId,
	}
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	v, ok := s.m[key]
	if !ok {
		err = errors.New("transaction not found")
		return
	}

	item = &v

	return
}

func (s *StorageTransaction) AddTransaction(playerName, transactionRef, sessionId string, withdraw, deposit int, rollBackStatus bool) (item entity.Transaction, err error) {

	item = entity.Transaction{
		PlayerName:     playerName,
		WithDraw:       withdraw,
		Deposit:        deposit,
		TransactionRef: transactionRef,
		SessionID:      sessionId,
		RollBackStatus: rollBackStatus,
	}

	key := entity.TransactionKey{
		TransactionRef: transactionRef,
		SessionID:      sessionId,
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.m[key] = item

	return
}

func (s *StorageTransaction) UpdateStatusTransaction(transactionRef, sessionId string, transactionStatus bool) (item *entity.Transaction, err error) {

	item = &entity.Transaction{}

	key := entity.TransactionKey{
		TransactionRef: transactionRef,
		SessionID:      sessionId,
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	v, ok := s.m[key]
	if !ok {
		err = errors.New("transaction not found")
		return
	}

	v.RollBackStatus = transactionStatus

	s.m[key] = v

	item = &v

	return
}
