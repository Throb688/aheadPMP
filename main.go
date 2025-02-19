package main

import (
	"aheadPMP/config"
	"aheadPMP/global"
	"aheadPMP/routes"
	"aheadPMP/utils"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func init() {

	err := global.ConnectToNode(config.Grpc)
	if err != nil {
		log.Fatalf("连接节点失败 err:%v", err)
	}
	log.Println("连接节点成功")

	PMPTokenAddress := common.HexToAddress(config.ContractAddress)
	err = global.LoadContract(PMPTokenAddress)
	if err != nil {
		log.Fatalf("加载合约失败 err:%v", err)
	}
	log.Println("加载合约成功")

	err = global.LoadAccountsFromFile(config.FilePath)
	if err != nil {
		log.Fatalf("加载账户信息失败 err:%v", err)
	}
	log.Println("加载账户信息成功")

}

func main() {
	r := routes.Router()
	utils.FindOldEvent()
	go utils.ListenEvent()

	r.Run(":8080")
}
