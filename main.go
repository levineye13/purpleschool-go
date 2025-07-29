package main

import (
	"bin/bins"
	"fmt"
)

func main() {
  bin, _ := bins.NewBin("uuid", "binName", true);
  fmt.Println(bin)
}