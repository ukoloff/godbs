package dbs

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	dbs := NewCircle(1)
	if j, e := json.Marshal(dbs); e != nil {
		t.Error(e)
	} else if string(j) != `[{"partid":"Circle","paths":[[[1,0,-1],[-1,0,-1],[1,0,0]]]}]` {
		t.Errorf("Invalid JSON")
	}
}

func TestUnmarshal(t *testing.T) {
	js := `[{"partid":"Circle","paths":[[[1,0,1],[-1,0,0]]]}]`
	var dbs DBS
	json.Unmarshal([]byte(js), &dbs)
	if len(dbs) != 1 {
		t.Errorf("Should have 1 part")
	}
	part := dbs[0]
	if part.Name != "Circle" {
		t.Errorf("Invalid part name")
	}
	if len(part.Paths) != 1 {
		t.Errorf("Should have 1 path")
	}
	path := part.Paths[0]
	if len(path) != 2 {
		t.Errorf("Should have 2 nodes")
	}
}
