package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

type Account struct {
	login    string
	password string
	url      string
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_=!@"

func (acc *Account) OutputAccount() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(length int) error {
	if length < 0 {
		return errors.New("PASSWORD_LENGTH")
	}

	newPass := make([]rune, length)

	for index := range length {
		randomIndex := rand.Intn(len(letters))
		item := letters[randomIndex]
		newPass[index] = rune(item)
	}

	acc.password = string(newPass)

	return nil
}

func NewAccount(login, password, urlValue string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, urlErr := url.ParseRequestURI(urlValue)

	if urlErr != nil {
		return nil, errors.New("INVALID_URL")
	}

	acc := Account{
		login:    login,
		password: password,
		url:      urlValue,
	}

	if len(acc.password) <= 4 {
		acc.generatePassword(10)
	}

	return &acc, nil
}