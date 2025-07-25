package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	name      string
	private   bool
	createdAt time.Time
}

type BinList = []Bin;

func newBin(id, name string, private bool) (*Bin, error) {
  if id == "" {
    return nil, errors.New("INVALID_ID")
  }

  if name == "" {
    return nil, errors.New("INVALID_NAME")
  }

  return &Bin{
    id: id,
    name: name,
    private: private,
    createdAt: time.Now(),
  }, nil
}

func main() {
  bin, binError := newBin("uuid", "binName", true);
  
  if binError != nil {
    switch binError.Error() {
      case "INVALID_ID":
        fmt.Println("Невалидный идентификатор")

      case "INVALID_NAME":
        fmt.Println("Невалидное наименование")
    }
  } else {
    fmt.Println(bin.id, bin.name, bin.private, bin.createdAt)
  }
}