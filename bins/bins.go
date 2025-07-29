package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	name      string
	private   bool
	createdAt time.Time
}

type BinList = []Bin

func NewBin(id, name string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	return &Bin{
		id:        id,
		name:      name,
		private:   private,
		createdAt: time.Now(),
	}, nil
}