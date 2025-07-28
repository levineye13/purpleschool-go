package main

import (
	"fmt"
	"password/account"
)

func promptData(text string) string {
	var res string

	fmt.Print(text);
  fmt.Scanln(&res);

  return res
}

func main() {
  login := promptData("Введите логин: ");
  password := promptData("Введите пароль: ");
  url := promptData("Введите URL: ");

  acc, accErr := account.NewAccount(login, password, url);

  if accErr != nil {
    fmt.Println(accErr);
    return;
  }

  acc.OutputAccount();
}