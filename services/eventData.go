package services

import (
	"aheadPMP/global"
	"strings"
)

func GetEventData() []global.TransferEventData {
	return global.GetTransferEvents()
}

func SearchForEventData(_data string) (string, []global.TransferEventData, []global.TransferEventData) {
	var expensesResult, incomeResult []global.TransferEventData
	address := _data

	if len(_data) < 10 {
		for key, value := range global.AddressToNameMap {
			if value == _data {
				address = key
				break
			}
		}
	}
	address = strings.ToLower(address)

	if len(address) == 66 {
		for _, v := range global.GetTransferEvents() {
			if strings.ToLower(v.TxHash.Hex()) == address {
				expensesResult = append(expensesResult, v)
			}
		}
		return "", expensesResult, incomeResult
	}

	for _, v := range global.GetTransferEvents() {
		if strings.ToLower(v.From.Hex()) == address {
			expensesResult = append(expensesResult, v)
		}
		if strings.ToLower(v.To.Hex()) == address {
			incomeResult = append(incomeResult, v)
		}
	}

	return global.AddressToNameMap[address], expensesResult, incomeResult
}
