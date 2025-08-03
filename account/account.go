package account

import (
	"errors"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_=!@"

func (acc *Account) OutputAccount() {
  output := strings.Join([]string{
    acc.Login, 
    acc.Password, 
    acc.Url,
    }, " ")

  color.Red(output)
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

	acc.Password = string(newPass)

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
		Login:    login,
		Password: password,
		Url:      urlValue,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
	}

	if len(acc.Password) <= 4 {
		acc.generatePassword(10)
	}

	return &acc, nil
}

