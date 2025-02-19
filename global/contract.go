package global

import (
	"aheadPMP/contract"
	"github.com/ethereum/go-ethereum/common"
)

var (
	AheadPMPContract *contract.Contract
)

func LoadContract(contractAddress common.Address) error {
	var err error
	AheadPMPContract, err = contract.NewContract(contractAddress, Client)
	if err != nil {
		return err
	}
	return nil
}
