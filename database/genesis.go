package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis() (genesis, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return genesis{}, err
	}

	genesisFilePath := filepath.Join(cwd, db, genFileName)
	content, err := ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return genesis{}, err
	}

	var loadedGenesis genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return genesis{}, err
	}

	return loadedGenesis, nil
}
