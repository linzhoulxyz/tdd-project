package stocks

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

func (m Money) Times(mulitplier int) Money {
	return Money{
		amount:   m.amount * float64(mulitplier),
		currency: m.currency,
	}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}
