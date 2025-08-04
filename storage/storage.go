package storage

import (
	"bin/bins"
)

type Storage struct {
  Bins bins.BinList `json:"bins"`
}

