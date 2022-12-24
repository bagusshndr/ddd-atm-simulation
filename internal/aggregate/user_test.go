package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	t.Run("new User", func(t *testing.T) {
		newUser, _ := NewUser("Bagus", 0)
		expected, _ := NewUser("Bagus", 0)
		assert.Equal(t, expected, newUser)
		assert.NotNil(t, newUser)
	})

	t.Run("new Users", func(t *testing.T) {
		newUser, _ := NewUsers("Bagus", 0)
		expected, _ := NewUsers("Bagus", 0)
		assert.Equal(t, expected, newUser)
		assert.NotNil(t, newUser)
	})

	t.Run("decrease amount", func(t *testing.T) {
		newUser, _ := NewUsers("Bagus", 150)
		decreaseAmount := newUser.DecreaseAmount(100)

		assert.NoError(t, decreaseAmount)
	})

	t.Run("failed decrease amount", func(t *testing.T) {
		newUser, _ := NewUsers("Bagus", 150)
		decreaseAmount := newUser.DecreaseAmount(0)

		assert.Error(t, decreaseAmount)
	})

	t.Run("increase amount", func(t *testing.T) {
		newUser, _ := NewUsers("Bagus", 150)
		increaseAmount := newUser.IncreaseAmount(100)

		assert.NoError(t, increaseAmount)
	})

	t.Run("failed increase amount", func(t *testing.T) {
		newUser, _ := NewUsers("Bagus", 150)
		increaseAmount := newUser.IncreaseAmount(0)

		assert.Error(t, increaseAmount)
	})

	t.Run("rebuild Users", func(t *testing.T) {
		newUser := RebuildUser(1, "Bagus", 0)
		expected := RebuildUser(1, "Bagus", 0)
		assert.Equal(t, expected, newUser)
		assert.NotNil(t, newUser)
	})

}
