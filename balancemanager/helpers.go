package balancemanager

import (
	"errors"
	"strconv"
	"strings"
)

func monetaryValueFromString(amount string) (inMonetaryValue int64, err error) {
	if !strings.Contains(amount, ".") {
		return 0, errors.New("the amount should have the form like '340.26', where digits before dot are major units of your currency and digits after dot are minor units")
	}
	s := strings.Replace(amount, ".", "", -1)
	inMonetaryValue, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return inMonetaryValue, nil
}
