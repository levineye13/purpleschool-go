package main

import (
	"bin/bins"
	"bin/storage"
	"encoding/json"
	"fmt"
)

func main() {
  bin1, _ := bins.NewBin("uuid1", "binName", true);
  bin2, _ := bins.NewBin("uuid2", "binName", false);
  binStorage, _ := storage.GetStorage()

  binStorage.AddBin(*bin1)
  binStorage.AddBin(*bin2)
  bins, _ := binStorage.GetBins()

  value, _ := json.MarshalIndent(bins, "", " ")
  fmt.Println(string(value))
}