package utils

import (
	"aheadPMP/config"
	"aheadPMP/global"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

func FindOldEvent() {

	blockNumber, err := global.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get latest block number: %v", err)
	}

	startBlock := big.NewInt(9000)             // 起始区块
	endBlock := big.NewInt(int64(blockNumber)) // 结束区块

	// 创建查询过滤器
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{common.HexToAddress(config.ContractAddress)},
	}

	logs, err := global.Client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	// 解析日志
	for _, vLog := range logs {
		// 使用生成的合约绑定代码解析日志
		event, err := global.AheadPMPContract.ParseTransfer(vLog)
		if err != nil {
			log.Printf("Failed to parse Transfer event: %v", err)
			continue
		}

		// 获取交易哈希
		txHash := vLog.TxHash

		// 获取区块高度
		blockNumber := vLog.BlockNumber

		// 获取区块时间戳
		block, err := global.Client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
		if err != nil {
			log.Printf("Failed to get block: %v", err)
			continue
		}
		timestamp := time.Unix(int64(block.Time()), 0)

		// 保存事件数据到全局变量
		transferEvent := global.TransferEventData{
			From:        event.From,
			To:          event.To,
			Value:       event.Value,
			TxHash:      txHash,
			BlockNumber: blockNumber,
			Timestamp:   timestamp,
		}

		global.TransferEvents = append(global.TransferEvents, transferEvent)

	}
}

// ListenEvent 改为实时监听 Transfer 事件
func ListenEvent() {
	// 创建查询过滤器，用于订阅 Transfer 事件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(config.ContractAddress)}, // 监听指定合约地址
	}

	// 创建事件订阅
	logsChannel := make(chan types.Log)

	sub, err := global.Client.SubscribeFilterLogs(context.Background(), query, logsChannel)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	// 持续监听事件
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Error while subscribing: %v", err)
		case vLog := <-logsChannel:
			// 使用合约生成的绑定代码解析日志
			event, err := global.AheadPMPContract.ParseTransfer(vLog)
			if err != nil {
				log.Printf("Failed to parse Transfer event: %v", err)
				continue
			}

			// 获取交易哈希
			txHash := vLog.TxHash

			// 获取区块高度
			blockNumber := vLog.BlockNumber

			// 获取区块时间戳
			block, err := global.Client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to get block: %v", err)
				continue
			}
			timestamp := time.Unix(int64(block.Time()), 0)

			// 保存事件数据到全局变量
			transferEvent := global.TransferEventData{
				From:        event.From,
				To:          event.To,
				Value:       event.Value,
				TxHash:      txHash,
				BlockNumber: blockNumber,
				Timestamp:   timestamp,
			}

			// 保存到全局事件列表
			global.TransferEvents = append(global.TransferEvents, transferEvent)

			//// 输出日志
			//fmt.Printf("Transfer Event: From=%s, To=%s, Value=%s, TxHash=%s, BlockNumber=%d, Timestamp=%s\n",
			//	event.From.Hex(),
			//	event.To.Hex(),
			//	event.Value.String(),
			//	txHash.Hex(),
			//	blockNumber,
			//	timestamp.String(),
			//)
		}
	}
}
