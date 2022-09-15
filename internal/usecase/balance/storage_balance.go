package balance

import (
	"dovran/mascot/internal/entity"
	"errors"
	"sync"
)

type StorageBalance struct {
	m   map[string]entity.Balance
	mtx sync.Mutex
}

func NewStorageBalance() *StorageBalance {
	var s StorageBalance
	s.m = make(map[string]entity.Balance)

	return &s
}

func (s StorageBalance) GetBalance(playerName string) (item *entity.Balance, err error) {

	item = &entity.Balance{}

	v, ok := s.m[playerName]
	if !ok {
		err = errors.New("player not found")
		return
	}

	item.PlayerName = v.PlayerName
	item.Balance = v.Balance
	item.FreeRoundsLeft = v.FreeRoundsLeft

	return
}

func (s StorageBalance) UpdateBalance(playerName string, balance, freeRoundsLeft int) (item *entity.Balance, err error) {
	item = &entity.Balance{}

	_, ok := s.m[playerName]
	if !ok {
		err = errors.New("player not found")
		return
	}

	s.m[playerName] = entity.Balance{
		PlayerName:     playerName,
		Balance:        balance,
		FreeRoundsLeft: freeRoundsLeft,
	}

	item.PlayerName = playerName
	item.Balance = balance
	item.FreeRoundsLeft = freeRoundsLeft

	return
}
