package main

import (
	"tdd/stocks"
	"testing"
)

func TestMutiplication(t *testing.T) {
	tenEuros := stocks.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := stocks.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := stocks.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := stocks.NewMoney(1000.5, "KRW")

	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

func TestAddition(t *testing.T) {
	var portfolio stocks.Portfolio
	var portfolioInDollars stocks.Money

	fiveDollars := stocks.NewMoney(5, "USD")
	tenDollars := stocks.NewMoney(10, "USD")
	fifteenDollars := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioInDollars)
}

func assertEqual(t *testing.T, expected stocks.Money, actual stocks.Money) {
	if expected != actual {
		t.Errorf("Excepted %+v Got %+v", expected, actual)
	}
}
