package Model

import (
	"htSample/global"
	"testing"
)

func ConnectDB(t *testing.T) error {
	err := global.InitDB("dev.vincent78.top", 15432, "htsample", "htsample", "htsample")
	if err != nil {
		t.Errorf("\nconnect the db failure : %v", err.Error())
	} else {
		t.Logf("\nconnect the db success")
	}
	return err
}
