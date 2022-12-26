package data

import (
	"ddd-atm-simulation/internal/aggregate"
	"ddd-atm-simulation/internal/enum"
)

var (
	Transaction = func() *aggregate.Transaction {
		return aggregate.RebuildTransaction(1, 1, enum.Transfer, 1, 1000)
	}

	User = func() *aggregate.User {
		return aggregate.RebuildUser(1, "Bagus", 1000)
	}
)
