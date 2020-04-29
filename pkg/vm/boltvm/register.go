package boltvm

import (
	"fmt"
)

type BoltContract struct {
	// enable/disable bolt contract
	Enabled bool
	// contract name
	Name string
	// contract address
	Address string
	// Contract is contract object
	Contract Contract
}

// register contract
func Register(contracts []*BoltContract) map[string]Contract {
	boltRegister := make(map[string]Contract)
	for _, c := range contracts {
		if _, ok := boltRegister[c.Address]; ok {
			panic("duplicate bolt contract address")
		} else {
			boltRegister[c.Address] = c.Contract
		}
	}
	return boltRegister
}

func GetBoltContract(address string, boltRegister map[string]Contract) (contract Contract, err error) {
	var ok bool
	if contract, ok = boltRegister[address]; !ok {
		return nil, fmt.Errorf("the address %v is not a bolt contract", address)
	}
	return contract, nil
}
