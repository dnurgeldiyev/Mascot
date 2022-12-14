package balance

import (
	"dovran/mascot/internal/entity"
	"errors"
	"fmt"
	"sync"
)

type StorageBalance struct {
	m   map[string]entity.Balance
	mtx sync.RWMutex
}

func NewStorageBalance() *StorageBalance {
	var s StorageBalance
	s.m = make(map[string]entity.Balance)

	return &s
}

func (s *StorageBalance) GetBalance(playerName string) (item *entity.Balance, err error) {

	item = &entity.Balance{}

	s.mtx.RLock()
	defer s.mtx.RUnlock()

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

func (s *StorageBalance) UpdateBalance(playerName string, balance, freeRoundsLeft int) (item *entity.Balance, err error) {
	item = &entity.Balance{}

	s.mtx.Lock()
	defer s.mtx.Unlock()

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

func (s *StorageBalance) AddPlayerAndBalance(playerName string, balance int) (err error) {

	s.mtx.Lock()
	defer s.mtx.Unlock()

	_, ok := s.m[playerName]
	if ok {
		err = errors.New(fmt.Sprintf("player with name: %s exist", playerName))
		return
	}

	s.m[playerName] = entity.Balance{
		PlayerName:     playerName,
		Balance:        balance,
		FreeRoundsLeft: 0,
	}

	return
}
