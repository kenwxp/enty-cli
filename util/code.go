package util

import (
	"math/big"
)

const HSF_PRO = "100"                 //*2
const USTD_PRO = "100000000000000000" //*18

func CodeCalculation(num string, codeTypes string) string {
	switch codeTypes {
	case "USDT":
		b1, _ := new(big.Float).SetString(num)
		b2, _ := new(big.Float).SetString("0.000000000000000001")
		mul := new(big.Float).Mul(b1, b2)
		return mul.String()
	case "HSF":
		b1, _ := new(big.Float).SetString(num)
		b2, _ := new(big.Float).SetString("0.01")
		mul := new(big.Float).Mul(b1, b2)
		return mul.String()
	}
	return "0"
}

func CodeCalculationToInteger(num string, codeTypes string) string {
	switch codeTypes {
	case "USDT":
		b1, _ := new(big.Float).SetString(num)
		b2, _ := new(big.Float).SetString("1000000000000000000")
		mul := new(big.Float).Mul(b1, b2)
		return mul.Text('f', -1)
	case "HSF":
		b1, _ := new(big.Float).SetString(num)
		b2, _ := new(big.Float).SetString("100")
		mul := new(big.Float).Mul(b1, b2)
		return mul.Text('f', -1)
	}
	return "0"
}
