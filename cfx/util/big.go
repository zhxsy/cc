package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

const (
	COIN_BASE float64 = 1000000000000000000
)

func EmptyStr(s string) string {
	if s == "" {
		return "0"
	}
	return s
}
func Add(a string, b string) (string, bool) {

	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return "", ok
	}

	rb = new(big.Int).Add(ra, rb)

	return rb.String(), true
}

func Sub(a string, b string) (string, bool) {

	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return "", ok
	}

	rb = new(big.Int).Sub(ra, rb)

	return rb.String(), true
}

func Mul(a string, b string) (string, bool) {

	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return "", ok
	}

	rb = new(big.Int).Mul(ra, rb)
	return rb.String(), true
}

func Divide(a string, b string, prec int) (string, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}

	rb, ok := new(big.Float).SetString(EmptyStr(b))
	if !ok {
		return "", ok
	}

	rb = rb.SetPrec(64).Quo(ra, rb)

	return rb.Text('f', prec), true
}

// IsLessThenZero if a < 0 then true, else false
func IsLessThenZero(a string) bool {
	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return ok
	}
	return ra.Cmp(big.NewInt(0)) == -1
}

// MulBase 剩以基数
func MulBase(a string) (string, bool) {
	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	return new(big.Int).Mul(ra, big.NewInt(int64(COIN_BASE))).String(), true
}

// DivideBase 除以基数
func DivideBase(a string, prec int) (string, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}

	rb := big.NewFloat(COIN_BASE)

	rb = rb.SetPrec(64).Quo(ra, rb)

	return rb.Text('f', prec), true
}

func ToFloat(s string) (float64, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(s))
	if !ok {
		return 0, ok
	}

	rb := big.NewFloat(COIN_BASE)

	rb = rb.SetPrec(64).Quo(ra, rb)
	f, _ := rb.Float64()
	return f, true
}

func ToInt(s string) (int64, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(s))
	if !ok {
		return 0, ok
	}

	rb := big.NewFloat(COIN_BASE)

	rb = rb.SetPrec(64).Quo(ra, rb)
	f, _ := rb.Int64()
	return f, true
}

// 比较大小
func Cmp(a, b string) (int, bool) {
	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return -2, ok
	}

	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return -2, ok
	}

	res := ra.Cmp(rb)
	return res, true
}

//16近制的字符串转化为10近制的字符串，并且除18位
func Hex2String(hexStr string) (string, bool) {
	dates, _ := hexutil.Decode(hexStr)
	a := new(big.Int).SetBytes(dates)
	return DivideBase(a.String(), 6)
}

// 16进制字符串转int64
func Hex2Dec(val string) int64 {
	if val[0:2] == "0x" {
		val = val[2:]
	}
	n, err := strconv.ParseInt(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

// 10进制转16进制
func Dec2Hex(n int64, x bool) string {
	if x {
		return fmt.Sprintf("0x%x", n)
	}
	return fmt.Sprintf("%x", n)
}
