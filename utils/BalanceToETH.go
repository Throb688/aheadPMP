package utils

import (
	"math/big"
	"strings"
)

func FormatBalanceToETH(balance *big.Int) string {
	// 1 ETH = 10^18 Wei
	ethValue := new(big.Rat).SetInt(balance)

	// 1 ETH = 10^18 Wei -> 计算 ETH = Wei / 10^18
	weiPerETH := new(big.Rat).SetInt(big.NewInt(1e18)) // 10^18
	ethValue.Quo(ethValue, weiPerETH)                  // 计算 ethValue = wei / 1e18

	// 转换为字符串，保留 18 位小数
	ethStr := ethValue.FloatString(18)

	// 如果小数部分全是 0，去掉小数部分
	if strings.Contains(ethStr, ".") {
		// 去除尾部无效的零
		ethStr = strings.TrimRight(ethStr, "0")
		// 如果小数部分全是零，去掉小数点
		if strings.HasSuffix(ethStr, ".") {
			ethStr = strings.TrimSuffix(ethStr, ".")
		}
	}

	return ethStr
}
