package main

import (
	"bin/bins"
	"bin/files"
	"bin/storage"
	"encoding/json"
	"fmt"
)

const binsFileName = "bins.json"

func main() {
  bin1, _ := bins.NewBin("uuid1", "binName", true);
  bin2, _ := bins.NewBin("uuid2", "binName", false);

  jsonDb := files.CreateJsonDb("bins.json")

  binStorage, _ := storage.GetStorage(jsonDb)

  binStorage.AddBin(*bin1)
  binStorage.AddBin(*bin2)
  bins, _ := binStorage.GetBins()

  value, _ := json.MarshalIndent(bins, "", " ")
  fmt.Println(string(value))
}