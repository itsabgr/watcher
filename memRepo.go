package watcher

import (
	"math/big"
	"sync"
)

type memRepo struct {
	_map  map[string][]Tx
	mutex sync.RWMutex
}

func (m *memRepo) StoreTxsByBlockID(tx Tx) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m._map[tx.Block().String()]; !ok {
		m._map[tx.Block().String()] = []Tx{}
	}
	m._map[tx.Block().String()] = append(m._map[tx.Block().String()], tx)
	return nil
}

func (m *memRepo) FindTxsByBlockID(blockID *big.Int) ([]Tx, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	txs, ok := m._map[blockID.String()]
	if !ok {
		return []Tx{}, nil
	}
	return txs, nil
}
func (m *memRepo) PurgeTxsByBlockID(blockID *big.Int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m._map, blockID.String())
	return nil
}

func NewMemRepo() (Repo, error) {
	repo := new(memRepo)
	repo._map = make(map[string][]Tx)
	return repo, nil
}
