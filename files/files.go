package files

import (
	"errors"
	"os"
	"strings"
)

func CheckJsonExt(name string) bool {
  splitedName := strings.Split(name, ".")

  if len(splitedName) < 2 {
    return false;
  }

  ext := splitedName[1]

  return ext == "json"
}

func ReadFile(name string) ([]byte, error) {
  

	file, err := os.ReadFile(name)

  if err != nil {
    return nil, errors.New("CANT_READ_FILE")
  }

  return file, nil
}

func WriteFile(name string, data []byte) error {
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

func CreateFile(name string) (*os.File, error) {
  _, err := os.Stat(name)
  isNotExist := errors.Is(err, os.ErrNotExist)

  if isNotExist {
    file, err := os.Create(name)

    if err != nil {
      defer file.Close()
      return file, err
    }
  }

  file, err := os.Open(name)

  if err != nil {
    return nil, err
  }
  
  defer file.Close()


  return file, nil
}