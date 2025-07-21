package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func promptData(text string) string {
	var res string

	fmt.Print(text);
  fmt.Scan(&res);

  return res
}

func outputAccount(acc account)  {
  fmt.Println(acc.login, acc.password, acc.url);
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

  outputAccount(acc);
}