package global

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type TransferEventData struct {
	From        common.Address // 转账发起方
	To          common.Address // 转账接收方
	Value       *big.Int       // 转账金额
	TxHash      common.Hash    // 交易哈希
	BlockNumber uint64         // 区块高度
	Timestamp   time.Time      // 区块时间戳
}

var TransferEvents []TransferEventData

func GetTransferEvents() []TransferEventData {
	return TransferEvents
}

func PrintEventLogs() {
	for _, eventLog := range TransferEvents {
		fmt.Printf("From=%s, To=%s, Value=%s, TxHash=%s, BlockNumber=%d, Timestamp=%s\n",
			eventLog.From.Hex(),
			eventLog.To.Hex(),
			eventLog.Value,
			eventLog.TxHash,
			eventLog.BlockNumber,
			eventLog.Timestamp,
		)
	}
}
