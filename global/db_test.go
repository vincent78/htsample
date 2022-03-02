package global

import (
	"testing"
)

func Test_DB(t *testing.T) {
	err := InitDB("dev.vincent78.top", 15432, "htsample", "htsample", "htsample")
	if err != nil {
		t.Errorf("connect the db failure : %v", err.Error())
	} else {
		t.Log("connect the db success")
	}
}
