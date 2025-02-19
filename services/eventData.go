package services

import "aheadPMP/global"

func GetEventData() []global.TransferEventData {
	return global.GetTransferEvents()
}
