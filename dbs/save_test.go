package dbs

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	dbs := DBS{
		Part{
			Name: "Circle",
			Paths: []Path{
				Path{
					Node{Point{1, 0}, 1},
					Node{Point{-1, 0}, 0},
				},
			},
		},
	}
	var out bytes.Buffer

	dbs.Save(&out)
	fmt.Println(out.Bytes())
	f, _ := os.Create("x.dbs")
	dbs.Save(f)
}
