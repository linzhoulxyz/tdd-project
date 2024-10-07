package stocks

import "errors"

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(bank Bank, currency string) (*Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, m := range p {
		if convertedAmount, err := bank.Convert(m, currency); err == nil {
			total += convertedAmount.amount
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}

	if len(failedConversions) == 0 {
		totalMoney := NewMoney(total, currency)
		return &totalMoney, nil
	}

	failures := "["
	for _, f := range failedConversions {
		failures += f + ","
	}
	failures += "]"
	return nil, errors.New("Missing exchange rate(s): " + failures)
}
