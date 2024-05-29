package valueobject

import (
	"errors"
	"fmt"
)

type Pair struct {
	BaseCurrency  string
	QuoteCurrency string
}

func NewPair(baseCurrency string, quoteCurrency string) *Pair {
	return &Pair{BaseCurrency: baseCurrency, QuoteCurrency: quoteCurrency}
}

func (receiver *Pair) ToStringKraken() (string, error) {
	symbol := receiver.BaseCurrency + receiver.QuoteCurrency
	switch symbol {
	case "BTCUSD":
		return btcusd, nil
	case "BTCCHF":
		return btcchf, nil
	case "BTCEUR":
		return btceur, nil
	default:
		return "", errors.New("invalid kraken symbol standardization")
	}
}

func (receiver *Pair) ToStringDivider() string {
	return fmt.Sprintf(
		"%s/%s",
		receiver.BaseCurrency,
		receiver.QuoteCurrency,
	)
}

func (receiver *Pair) ToString() string {
	return fmt.Sprintf(
		"%s%s",
		receiver.BaseCurrency,
		receiver.QuoteCurrency,
	)
}
