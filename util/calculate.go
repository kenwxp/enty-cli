package util

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
)

func CalculateInt64(x int64, y int64, operator string) (i int64) {
	switch operator {
	case "add":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Add(a, b)
		i := z.Int64()
		return i
	case "sub":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Sub(a, b)
		i := z.Int64()
		return i
	case "mul":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Mul(a, b)
		i := z.Int64()
		return i
	case "div":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Div(a, b)
		i := z.Int64()
		return i
	}
	return i
}
func CalculateString(x string, y string, operator string) (i string) {
	if x == "" {
		x = "0"
	}
	if y == "" {
		y = "0"
	}
	a := new(big.Float)
	b := new(big.Float)
	xbf, xok := a.SetString(x)
	if !xok {
		fmt.Println("err:", xok)
		panic("error in trans string into big float")
	}
	ybf, yok := b.SetString(y)
	if !yok {
		fmt.Println("err:", yok)
		panic("error in trans string into big float")
	}
	switch operator {
	case "add":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			fmt.Println("err:", xErr)
			fmt.Println("err:", yErr)
			panic("error in trans string into int")
		}
		i := strconv.Itoa(ai + bi)
		return i
	case "sub":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			fmt.Println("err:", xErr)
			fmt.Println("err:", yErr)
			panic("error in trans string into int")
		}
		if ai < bi {
			panic("wrong order")
		}
		i := strconv.Itoa(ai - bi)
		return i
	case "mul":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			fmt.Println("err:", xErr)
			fmt.Println("err:", yErr)
			panic("error in trans string into int")
		}
		i := strconv.Itoa(ai * bi)
		return i
	case "div":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			fmt.Println("err:", xErr)
			fmt.Println("err:", yErr)
			panic("error in trans string into int")
		}
		if bi == 0 {
			panic("the denominator is zero")
		}
		i := strconv.Itoa(ai / bi)
		return i
	case "cmp":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			fmt.Println("err:", xErr)
			fmt.Println("err:", yErr)
			panic("error in trans string into int")
		}
		if ai > bi {
			i = "1"
		} else if ai < bi {
			i = "-1"
		} else {
			i = "0"
		}
		return i
	case "addBigFU":
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 18)
		return i
	case "subBigFU":
		if xbf.Cmp(ybf) < 0 {
			panic("wrong order")
		}
		ibf := xbf.Sub(xbf, ybf)
		i = ibf.Text('f', 18)
		return i
	case "addBigFH":
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 2)
		return i
	case "cmpBigF":
		i = strconv.Itoa(xbf.Cmp(ybf))
		return i
	case "divBigF":
		ibf := xbf.Quo(xbf, ybf)
		i = ibf.Text('f', 18) //保留6位小数
		return i
	case "mulBigF":
		ibf := xbf.Mul(xbf, ybf)
		i = ibf.Text('f', 18) //保留6位小数
		return i
	case "subBigFH":
		if xbf.Cmp(ybf) < 0 {
			panic("wrong order")
		}
		ibf := xbf.Sub(xbf, ybf)
		i = ibf.Text('f', 2)
		return i
	}

	return i
}
func Digit(x string, operator string) (i string, definedErr *MessageError) {
	unit := new(big.Float)
	a := new(big.Float)
	bf, ok := a.SetString(x)
	if !ok {
		definedErr := NewMsgError(4, "error in trans string into bigint")
		return "", definedErr
	}
	switch operator {
	case "div18":
		unit.SetString("1000000000000000000")
		bf.Quo(bf, unit)
		i = bf.Text('f', 18)
		return i, definedErr
	case "div2":
		unit.SetString("100")
		bf.Quo(bf, unit)
		i = bf.Text('f', 2)
		return i, definedErr
	case "div9":
		unit.SetString("1000000000")
		bf.Quo(bf, unit)
		i = bf.Text('f', 18)
		return i, definedErr
	case "mul2":
		unit.SetString("100")
		bf.Mul(bf, unit)
		i = bf.Text('f', 0)
		return i, definedErr
	case "mul9":
		unit.SetString("1000000000")
		bf.Mul(bf, unit)
		i = bf.Text('f', 0)
		return i, definedErr
	case "mul18":
		unit.SetString("1000000000000000000")
		bf.Mul(bf, unit)
		i = bf.Text('f', 0)
		return i, definedErr
	}

	return i, definedErr
}

/*
	addition
	support type
	string,int,int64,float64
*/
func Addition(args ...interface{}) string {
	var num float64
	for _, v := range args {
		types := reflect.TypeOf(v).String()
		switch types {
		case "string":
			v1, _ := strconv.ParseFloat(v.(string), 64)
			num += v1
		case "int":
			num += float64(v.(int))
		case "int64":
			num += float64(v.(int64))
		case "float64":
			num += v.(float64)
		}
	}
	return strconv.FormatFloat(num, 'f', -1, 64)
}

/*
	Subtraction
	support type
	string,int,int64,float64
	return  [1] - [2] - [3]
*/
func Subtraction(args ...interface{}) string {
	var num float64
	for i, v := range args {
		types := reflect.TypeOf(v).String()
		switch types {
		case "string":
			v1, _ := strconv.ParseFloat(v.(string), 64)
			num -= v1
		case "int":
			num -= float64(v.(int))
		case "int64":
			num -= float64(v.(int64))
		case "float64":
			num -= v.(float64)
		}
		if i == 0 {
			num = num - num - num
		}
	}
	return strconv.FormatFloat(num, 'f', -1, 64)
}

/*
	Multiplication
	support type
	string,int,int64,float64
	return  [1] * [2] * [3]
*/
func Multiplication(args ...interface{}) string {
	var num float64 = 1
	for _, v := range args {
		types := reflect.TypeOf(v).String()
		switch types {
		case "string":
			v1, _ := strconv.ParseFloat(v.(string), 64)
			num *= v1
		case "int":
			num *= float64(v.(int))
		case "int64":
			num *= float64(v.(int64))
		case "float64":
			num *= v.(float64)
		}
	}
	return strconv.FormatFloat(num, 'f', -1, 64)
}

/*
	Multiplication
	support type
	string,int,int64,float64
	return  [1] * [2] * [3]
*/
func MultiplicationPrec0(args ...interface{}) string {
	var num float64 = 1
	for _, v := range args {
		types := reflect.TypeOf(v).String()
		switch types {
		case "string":
			v1, _ := strconv.ParseFloat(v.(string), 64)
			num *= v1
		case "int":
			num *= float64(v.(int))
		case "int64":
			num *= float64(v.(int64))
		case "float64":
			num *= v.(float64)
		}
	}
	return strconv.FormatFloat(num, 'f', 0, 64)
}

func StrNanoFILToFilStr2(nanoFilStr string) string {
	float, err := strconv.ParseFloat(nanoFilStr, 64)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%.2f", float/1000000000.0)
}

func StrNanoFILToFilStr(nanoFilStr string, num string) string {
	float, err := strconv.ParseFloat(nanoFilStr, 64)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%."+num+"f", float/1000000000.0)
}
func StrNanoFILToFilStr8(nanoFilStr string) string {
	float, err := strconv.ParseFloat(nanoFilStr, 64)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%.8f", float/1000000000000000000.0)
}

/*float64String to int64string*/
func FSToIS(floatstr string) string {
	floatV, _ := strconv.ParseFloat(floatstr, 64)
	return strconv.FormatInt(int64(floatV), 10)
}
