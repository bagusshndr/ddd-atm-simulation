package aggregate

import (
	"ddd-atm-simulation/internal/enum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {

	t.Run("new Transaction", func(t *testing.T) {
		newTransaction, _ := NewTransaction(1, enum.Deposit, 1, 100)
		expected, _ := NewTransaction(1, enum.Deposit, 1, 100)
		assert.Equal(t, expected, newTransaction)
		assert.NotNil(t, newTransaction)
	})

	t.Run("new Transactions", func(t *testing.T) {
		newTransaction, _ := NewTransactions(1, enum.Deposit, 1, 100)
		expected, _ := NewTransactions(1, enum.Deposit, 1, 100)
		assert.Equal(t, expected, newTransaction)
		assert.NotNil(t, newTransaction)
	})

	t.Run("rebuild Transactions", func(t *testing.T) {
		newTransaction := RebuildTransaction(1, 1, enum.Deposit, 1, 100)
		expected := RebuildTransaction(1, 1, enum.Deposit, 1, 100)
		assert.Equal(t, expected, newTransaction)
		assert.NotNil(t, newTransaction)
	})
}
