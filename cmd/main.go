package main

import (
	"fmt"

	"github.com/haquenafeem/rio"
)

func main() {
	fmt.Printf("%v %v %v %v %v %v %v %v",
		rio.Ignore("Ignore"),
		rio.Error("Error"),
		rio.Success("Success"),
		rio.Warn("Warn"),
		rio.Log("Log"),
		rio.Info("Info"),
		rio.Question("Question"),
		rio.Quote(`"Quote"`))
	fmt.Println()
}
