package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type State struct {
	Balances  map[Account]uint
	txMemPool []Tx

	dbFile *os.File
}

func NewStateFromDisk() (*State, error) {
	gen, err := loadGenesis()
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(filepath.Join(cwd, db, txFileName), os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	state := &State{balances, make([]Tx, 0), f}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var tx Tx
		err = json.Unmarshal(scanner.Bytes(), &tx)
		if err != nil {
			return nil, err
		}

		if err := state.apply(tx); err != nil {
			return nil, err
		}
	}

	return state, nil
}

func (s *State) apply(tx Tx) error {
	if tx.IsReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	if s.Balances[tx.From] < tx.Value {
		return fmt.Errorf("not enough balance")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value

	return nil
}

func (s *State) Add(tx Tx) error {
	if err := s.apply(tx); err != nil {
		return err
	}

	s.txMemPool = append(s.txMemPool, tx)

	return nil
}

func (s *State) Persist() error {
	memPool := make([]Tx, len(s.txMemPool))
	copy(memPool, s.txMemPool)

	for i := 0; i < len(memPool); i++ {
		txJson, err := json.Marshal(memPool[i])
		if err != nil {
			return err
		}

		if _, err = s.dbFile.Write(append(txJson, '\n')); err != nil {
			return err
		}
		s.txMemPool = s.txMemPool[1:]
	}

	return nil
}

func (s *State) Close() {
	s.dbFile.Close()
}
