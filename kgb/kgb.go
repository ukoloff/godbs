package main

import (
	"os"

	"github.com/ukoloff/godbs/dbs"
)

func main() {
	var dbs dbs.DBS
	dbs.MakeCircle(1)
	f, _ := os.Create("kgb.dbs")
	defer f.Close()
	dbs.Save(f)
}
