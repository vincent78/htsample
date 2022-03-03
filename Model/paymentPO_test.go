package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentPO_Create(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	p := NewPaymentPO("a001", "incoming", "usd", 100, 101, "tokenTest", "admin")
	e := p.Create()
	assert.Empty(t, e)
}
