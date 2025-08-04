package files

import (
	"errors"
	"os"
	"strings"
)

func checkJsonExt(name string) bool {
  splitedName := strings.Split(name, ".")
  ext := splitedName[1]

  return ext == "json"
}

func ReadFile(name string) ([]byte, error) {
  isJson := checkJsonExt(name)

  if !isJson {
    return nil, errors.New("INVALID_EXT")
  }

	file, err := os.ReadFile(name)

  if err != nil {
    return nil, errors.New("CANT_READ_FILE")
  }

  return file, nil
}

func WriteFile(name string, data []byte) error {
  isJson := checkJsonExt(name)

  if !isJson {
    return errors.New("INVALID_EXT")
  }

  file, err := os.Create(name)

  if err != nil {
    return errors.New("CANT_CREATE_FILE")
  }

  defer file.Close()

  _, err = file.Write(data)

  if err != nil {
    return errors.New("CANT_WRITE_FILE")
  }

  return nil
}
