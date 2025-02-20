package services

import (
	"aheadPMP/global"
	"aheadPMP/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tealeg/xlsx"
	"strings"
)

type ExportType struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Balance string `json:"balance"`
}

// ExportExcel 将用户和项目数据导出为 Excel 文件
func ExportExcel() (*xlsx.File, error) {
	var typeData []ExportType // 统一使用一个切片

	// 遍历 global.AddressToNameMap 生成数据
	for address, name := range global.AddressToNameMap {
		balance, err := global.AheadPMPContract.BalanceOf(nil, common.HexToAddress(address))
		if err != nil {
			// 处理错误，保持代码稳定，跳过有问题的地址
			continue
		}

		// 创建数据对象
		exportData := ExportType{
			Name:    name,
			Address: address,
			Balance: utils.FormatBalanceToETH(balance),
		}

		typeData = append(typeData, exportData)
	}

	// 创建一个新的 Excel 文件
	file := xlsx.NewFile()

	// 创建工作表
	userSheet, err := file.AddSheet("用户数据")
	if err != nil {
		return nil, err
	}
	projectSheet, err := file.AddSheet("项目信息")
	if err != nil {
		return nil, err
	}

	// 设置表头
	userHeader := userSheet.AddRow()
	userHeader.AddCell().SetString("姓名")
	userHeader.AddCell().SetString("账户地址")
	userHeader.AddCell().SetString("余额（PMP）")

	projectHeader := projectSheet.AddRow()
	projectHeader.AddCell().SetString("项目名称")
	projectHeader.AddCell().SetString("项目地址")
	projectHeader.AddCell().SetString("余额（PMP）")

	// 填充用户数据
	for _, user := range typeData {
		row := userSheet.AddRow()
		row.AddCell().SetString(user.Name)
		row.AddCell().SetString(user.Address)
		row.AddCell().SetString(user.Balance)
	}

	// 项目数据过滤并填充
	for _, project := range typeData {
		if strings.Contains(project.Name, "项目") {
			row := projectSheet.AddRow()
			row.AddCell().SetString(project.Name)
			row.AddCell().SetString(project.Address)
			row.AddCell().SetString(project.Balance)
		}
	}

	return file, nil
}
