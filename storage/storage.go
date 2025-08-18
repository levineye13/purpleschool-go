package storage

import (
	"bin/bins"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Db interface {
  Read() ([]byte, error)
  Write([]byte) error
}

type Storage struct {
  Bins bins.BinList `json:"bins"`
  UpdateAt time.Time `json:"updateAt"`
}

type StorageWithDb struct {
  Storage
  db Db
}

func GetStorage(db Db) (*StorageWithDb, error) {
  fileBytes, err := db.Read()

  if err != nil {
    return &StorageWithDb{
      Storage: Storage{
        Bins: bins.BinList{},
        UpdateAt: time.Now(),
      },
      db: db,
    }, nil
  }

  var storage *Storage

  err = json.Unmarshal(fileBytes, &storage)

  if err != nil {
    return &StorageWithDb{
      Storage: Storage{
        Bins: bins.BinList{},
        UpdateAt: time.Now(),
      },
      db: db,
    }, nil
  }

  return &StorageWithDb{
    Storage: *storage,
    db: db,
  }, nil
}

func (storage *StorageWithDb) AddBin(bin bins.Bin) error {
  data, err := storage.db.Read()

  if err != nil {
    return err
  }

  err = json.Unmarshal(data, &storage)

  if err != nil && len(data) > 0 {
    return err
  }

  storage.Bins = append(storage.Bins, bin)
  storage.UpdateAt = time.Now()

  newBins, err := json.Marshal(storage)

  if err != nil {
    return err
  }

  err = storage.db.Write(newBins)

  if err != nil {
    color.Red("Ошибка записи файла")
    return err
  }

  return nil
}

func (storage *StorageWithDb) GetBins() (*bins.BinList, error) {
  data, err := storage.db.Read()

  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(data, &storage)

  if err != nil && len(data) > 0 {
    return nil, err
  }

  return &storage.Bins, nil
}