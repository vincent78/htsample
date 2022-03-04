package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPaymentProcessPO(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	_, err = FindPaymentProcessList()
	assert.NoError(t, err)
}

func TestFindPaymentProcessListByToken(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	_, err = FindPaymentProcessListByToken("5839ceacb00b089")
	assert.NoError(t, err)
}
