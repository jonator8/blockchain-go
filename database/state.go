package database

import (
	"fmt"
	"os"
)

type State struct {
	Balances  map[Account]uint
	txMemPool []Tx

	dbFile *os.File
}

func NewStateFromDisk() (*State, error) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fmt.Println(currentDirectory)

	return nil, nil
}
