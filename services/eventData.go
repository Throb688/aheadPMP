package services

import (
	"aheadPMP/global"
	"strings"
)

func GetEventData() []global.TransferEventData {
	return global.GetTransferEvents()
}

func SearchForEventData(_data string) ([]global.TransferEventData, []global.TransferEventData) {
	var expensesResult, incomeResult []global.TransferEventData
	address := _data

	if len(_data) != 42 {
		for key, value := range global.AddressToNameMap {
			if value == _data {
				address = key
				break
			}
		}
	}
	address = strings.ToLower(address)

	for _, v := range global.GetTransferEvents() {
		if strings.ToLower(v.From.Hex()) == address {
			expensesResult = append(expensesResult, v)
		}
		if strings.ToLower(v.To.Hex()) == address {
			incomeResult = append(incomeResult, v)
		}
	}

	return expensesResult, incomeResult
}
