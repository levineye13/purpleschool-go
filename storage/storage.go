package storage

import (
	"bin/bins"
	"bin/files"
	"encoding/json"
	"errors"
	"time"

	"github.com/fatih/color"
)

type Storage struct {
  Bins bins.BinList `json:"bins"`
  UpdateAt time.Time `json:"updateAt"`
}

const binsFileName = "bins.json"

func GetStorage() (*Storage, error) {
  isJson := files.CheckJsonExt(binsFileName)

  if !isJson {
    color.Red("Необходим файл в формате JSON")
    return nil, errors.New("INVALID_EXT")
  }

  _, err := files.CreateFile(binsFileName)

  if err != nil {
    color.Red("Ошибка создания файла")
    return nil, err
  }

  fileBytes, err := files.ReadFile(binsFileName)

  if err != nil {
    return &Storage{
      Bins: bins.BinList{},
      UpdateAt: time.Now(),
    }, nil
  }

  var storage *Storage

  err = json.Unmarshal(fileBytes, &storage)

  if err != nil {
    return &Storage{
      Bins: bins.BinList{},
      UpdateAt: time.Now(),
    }, nil
  }

  return storage, nil
}

func (storage *Storage) AddBin(bin bins.Bin) error {
  isJson := files.CheckJsonExt(binsFileName)

  if !isJson {
    color.Red("Необходим файл в формате JSON")
    return errors.New("INVALID_EXT")
  }

  fileBytes, err := files.ReadFile(binsFileName)

  if err != nil {
    color.Red("Ошибка чтения файла")
    return err
  }

  err = json.Unmarshal(fileBytes, &storage)

  if err != nil && len(fileBytes) > 0 {
    color.Red("Ошибка преобразования из json")
    return err
  }

  storage.Bins = append(storage.Bins, bin)
  storage.UpdateAt = time.Now()

  newBins, err := json.Marshal(storage)

  if err != nil {
    color.Red("Ошибка преобразования в json")
    return err
  }

  err = files.WriteFile(binsFileName, newBins)

  if err != nil {
    color.Red("Ошибка записи файла")
    return err
  }

  return nil
}

func (storage *Storage) GetBins() (*bins.BinList, error) {
  isJson := files.CheckJsonExt(binsFileName)

  if !isJson {
    color.Red("Необходим файл в формате JSON")
    return nil, errors.New("INVALID_EXT")
  }

  fileBytes, err := files.ReadFile(binsFileName)

  if err != nil {
    color.Red("Ошибка чтения файла")
    return nil, err
  }

  err = json.Unmarshal(fileBytes, &storage)

  if err != nil && len(fileBytes) > 0 {
    color.Red("Ошибка преобразования из json")
    return nil, err
  }

  return &storage.Bins, nil
}