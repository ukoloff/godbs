package dbs

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	dbs := NewCircle(1)
	var out bytes.Buffer

	dbs.Save(&out)
	fmt.Println(out.Bytes())
	f, _ := os.Create("testdata/.!.dbs")
	dbs.Save(f)
}
