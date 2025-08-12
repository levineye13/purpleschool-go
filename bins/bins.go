package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Private   bool `json:"private"`
	CreatedAt time.Time `json:"createAt"`
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
		Id:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}, nil
}