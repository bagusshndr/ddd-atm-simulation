package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	assert.Equal(t, "", UndefinedFlag.String())
	assert.Equal(t, "Deposit", Deposit.String())
	assert.Equal(t, "Withdraw", Withdram.String())
	assert.Equal(t, "Transfer", Transfer.String())
}
