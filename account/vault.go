package account

import (
	"encoding/json"
	"password/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt  time.Time `json:"updateAt"`
}

const dataFileName string = "data.json"

func NewVault() *Vault {
  dataBytes, err := files.ReadFile(dataFileName)

  if err != nil {
    return &Vault{
      Accounts: []Account{},
      UpdateAt: time.Now(),
    }
  }

  var vault Vault

  err = json.Unmarshal(dataBytes, &vault)

  if err != nil {
    color.Red(err.Error())
  }

  return &vault
}

func (vault *Vault) ToBytes() ([]byte, error) {
  file, err := json.Marshal(&vault)

  if err != nil {
    return nil, err;
  }

  return file, nil
}

func (vault *Vault) AddAccount(acc Account) {
  vault.Accounts = append(vault.Accounts, acc)
  vault.UpdateAt = time.Now()

  dataBytes, err := vault.ToBytes()

  if err != nil {
    color.Red(err.Error())
  }

  files.WriteFile(dataFileName, dataBytes)
}

func FindAccountsByURL(url string) ([]Account, error) {
  file, err := files.ReadFile(dataFileName)

  if err != nil {
    return nil, err;
  }

  var vault Vault

  err = json.Unmarshal(file, &vault)

  if err != nil {
    color.Red(err.Error())
  }

  accounts := []Account{}

  for _, account := range vault.Accounts {
    if strings.Contains(account.Url, url) {
      accounts = append(accounts, account)
    }
  }

  return accounts, nil
}