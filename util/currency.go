package util

// Constants for all supported currencies
const (
	TL   = "TL"
	USD  = "USD"
	EURO = "EURO"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case TL, USD, EURO:
		return true
	}
	return false
}
