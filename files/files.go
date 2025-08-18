package files

import (
	"errors"
	"os"
	"strings"
)

type JsonDb struct {
  filename string
}

func CreateJsonDb(filename string) *JsonDb {
  return &JsonDb{
    filename: filename,
  }
}

func (db *JsonDb) CheckJsonExt(name string) bool {
  splitedName := strings.Split(name, ".")

  if len(splitedName) < 2 {
    return false;
  }

  ext := splitedName[1]

  return ext == "json"
}

func (db *JsonDb) Read() ([]byte, error) {
	file, err := os.ReadFile(db.filename)

  if err != nil {
    return nil, errors.New("CANT_READ_FILE")
  }

  return file, nil
}

func (db *JsonDb) Write(data []byte) error {
  file, err := os.Create(db.filename)

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

func (db *JsonDb) CreateFile() (*os.File, error) {
  _, err := os.Stat(db.filename)
  isNotExist := errors.Is(err, os.ErrNotExist)

  if isNotExist {
    file, err := os.Create(db.filename)

    if err != nil {
      defer file.Close()
      return file, err
    }
  }

  file, err := os.Open(db.filename)

  if err != nil {
    return nil, err
  }
  
  defer file.Close()


  return file, nil
}