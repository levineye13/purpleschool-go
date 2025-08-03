package menu

import (
	"encoding/json"
	"fmt"
	"password/account"
	"slices"

	"github.com/fatih/color"
)

func promptData(text string) string {
	var res string

	fmt.Print(text);
  fmt.Scanln(&res);

  return res
}

func createAccount()  {
  login := promptData("Введите логин: ")
  password := promptData("Введите пароль: ")
  url := promptData("Введите URL: ")

  newAccount, _ := account.NewAccount(login, password, url)
  vault := account.NewVault()
  vault.AddAccount(*newAccount)
}

func findAccount()  {
  url := promptData("Введите URL: ")
  accounts, _ := account.FindAccountsByURL(url)

  data, _ := json.MarshalIndent(accounts, "", " ");

  color.Cyan("Результат поиска: ")
  color.Cyan(string(data))
}

func GetMenu() {
  items := []string{"1", "2", "3", "4"}
  var menuItem string;

  Menu:
    for {
      fmt.Println(`Выберите вариант:
1: Создать аккаунт
2: Найти аккаунт
3: Удалить аккаунт
4: Выход
      `)

      _, menuItemErr := fmt.Scanln(&menuItem);

      if menuItemErr != nil {
        fmt.Println("Некорректный ввод")
        continue Menu;
      }

      if !slices.Contains(items, menuItem) {
        fmt.Println("Введите вариант из указанных")
        continue Menu;
      }

      switch menuItem {
        case "1":
          createAccount()
        case "2":
          findAccount()
        case "3":
          account.DeleteAccount()
        case "4":
          break Menu
      }
    }


}