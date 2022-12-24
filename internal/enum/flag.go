package enum

type Flag int

func (f Flag) String() string {
	switch f {
	case Deposit:
		return "Deposit"
	case Withdram:
		return "Withdraw"
	case Transfer:
		return "Transfer"
	}
	return ""
}

const (
	UndefinedFlag Flag = iota
	Deposit
	Withdram
	Transfer
)
