package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type account struct {
	login    string
	password string
	url      string
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_=!@";

func promptData(text string) string {
	var res string

	fmt.Print(text);
  fmt.Scan(&res);

  return res
}

func (acc *account) outputAccount()  {
  fmt.Println(acc.login, acc.password, acc.url);
}

func (acc *account) generatePassword(length int) error {
  if length < 0 {
    return errors.New("PASSWORD_LENGTH");
  }

  newPass := make([]rune, length);

  for index := range length {
    randomIndex := rand.Intn(len(letters));
    item := letters[randomIndex];
    newPass[index] = rune(item);
  }

  acc.password = string(newPass);

  return nil;
}

func main() {
  login := promptData("Введите логин: ");
  password := promptData("Введите пароль: ");
  url := promptData("Введите URL: ");

  acc := account{
    login: login,
    password: password,
    url: url,
  }

  if len(acc.password) == 0 {
    acc.generatePassword(10);
  }

  acc.outputAccount();
}