package currencyconversation

import (
	"testing"

	"github.com/andersfylling/concurrencyparser"
)

func TestCurrenciesToJsonStr(t *testing.T) {
	result := currenciesToJsonStr("EUR", "NOK")
	expects := `{"baseCurrency": "EUR", "targetCurrency":"NOK"}`
	if result != expects {
		t.Errorf("Incorrect json creation, got: %s, want: %s.", result, expects)
	}

	result = currenciesToJsonStr("USD", "NOK")
	expects = `{"baseCurrency": "USD", "targetCurrency":"NOK"}`
	if result != expects {
		t.Errorf("Incorrect json creation, got: %s, want: %s.", result, expects)
	}
}

func TestGetCurrencyRate(t *testing.T) {
	_, err := getCurrencyRate(currenciesToJsonStr("EUR", "NOK"))
	if err != nil {
		t.Error("Failed retrieving data from server. (assumption server is always available)")
	}
}

func TestFindCurrencyRate(t *testing.T) {
	exc := &concurrencyparser.ExchangeRate{
		Base:   "NOK",
		Target: "USD",
	}
	_, err := findCurrencyRate(exc)
	if err != nil {
		t.Error("Failed retrieving data from server. (assumption server is always available)")
	}

	exc.Base = "EUR"
	_, err = findCurrencyRate(exc)
	if err != nil {
		t.Error("Failed retrieving data from server. (assumption server is always available)")
	}

	exc.Target = "EUR"
	_, err = findCurrencyRate(exc)
	if err != nil {
		t.Error("Failed retrieving data from server. (assumption server is always available)")
	}
}
