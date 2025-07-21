package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TBookmarks = map[string]string;

func showMenu() {
	fmt.Println(`
	1. Посмотреть закладки
	2. Добавить закладку
	3. Удалить закладку
	4. Выход
	`);
}

func getMenuItemInput() (int, error) {
	var input string;

	fmt.Scan(&input);

	str := strings.TrimSpace(input);
	num, numErr := strconv.Atoi(str);

	if numErr != nil {
		fmt.Println("Введено некорректное число");
		return 0, numErr
	}

	if num < 1 || num > 4 {
		return 0, errors.New("MENU_RANGE_ERROR");
	}

	return num, nil;
}

func showItems(items TBookmarks) string {
	if len(items) == 0 {
		return "У вас пока нет сохраненных закладок";
	}

	itemsStr := ``;

	for name, text := range items {
		itemsStr += fmt.Sprintln(name, text);
	}

	return itemsStr;
}

func addItem(items TBookmarks, key, link string) error {
  item := items[key];

  if item != "" {
		return errors.New("MAP_UNIQUE_ERROR");
  }

	items[key] = link;

	return nil;
}

func getCreateItemInput() (string, string, error) {
	var key string;
	var link string;

	fmt.Println("Введите название закладки:");

	_, keyErr := fmt.Scan(&key);

	if keyErr != nil {
		return "", "", keyErr;
	}

	fmt.Println("Введите ссылку:");

	_, linkErr := fmt.Scan(&link);

	if linkErr != nil {
		return "", "", linkErr;
	}

	return key, link, nil;
}

func getDeleteItemInput() (string, error) {
	var key string;

	fmt.Println("Введите название закладки");
	_, keyErr := fmt.Scan(&key);

	if keyErr != nil {
		return "", keyErr;
	}

	return key, nil;
}


func main() {
	fmt.Println("Приложение - Утилита закладок");

	items := TBookmarks{};

  LoopLabel:
    for {
      showMenu();
      menuItem, menuRangeErr := getMenuItemInput();

      if menuRangeErr != nil {
        fmt.Println("Число выходит за пределы допустимых");
      }

      fmt.Print("\n");

      switch menuItem {
        case 1:
          fmt.Println("Ваши закладки:");
          itemsStr := showItems(items);
          fmt.Println(itemsStr);

        case 2: 
          fmt.Println("Создание закладки:");
          key, link, _ := getCreateItemInput();
          createItemErr := addItem(items, key, link);

          if createItemErr != nil {
            fmt.Println("Ошибка: закладка с таким названием уже существует");
          }

        case 3:
          fmt.Println("Удаление закладки:");
          key, _ := getDeleteItemInput();
          delete(items, key);

        case 4:
          break LoopLabel;
      }
    }
}