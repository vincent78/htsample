package utils

import "testing"

func TestGetUUID(t *testing.T) {

	for i := 0; i < 10; i++ {
		t.Logf("the uuid[%v] : %v", i, GetUUID())
	}
}
