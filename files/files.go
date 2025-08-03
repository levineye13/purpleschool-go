package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	file, err := os.ReadFile(name)

  if err != nil {
    color.Red("Не удалось прочитать файл")
    return nil, err;
  }

	return file, nil
}

func WriteFile(name string, data []byte) error {
	file, err := os.Create(name)

  if err != nil {
    color.Red("Не удалось создать файл")
    return err;
  }

  _, err = file.Write(data)

  if err != nil {
    color.Red("Не удалось записать данные в файл")
    return err;
  }

  file.Close()

	return nil
}