package global

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Account struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var AddressToNameMap map[string]string

func LoadAccountsFromFile(filePath string) error {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var accounts []Account

	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return err
	}

	AddressToNameMap = make(map[string]string)

	for _, account := range accounts {
		AddressToNameMap[strings.ToLower(account.Address)] = account.Name
	}

	return nil
}
