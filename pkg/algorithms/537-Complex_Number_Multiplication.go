package algorithms

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseComplex(number string) (real int, imag int, err error) {
	num := strings.Split(number, "+")
	if len(num) != 2 {
		return real, imag, errors.New("parse error")
	}
	real, err = strconv.Atoi(num[0])
	if err != nil {
		return
	}
	imag, err = strconv.Atoi(
		strings.TrimRight(num[1], "i"),
	)
	return
}

func complexNumberMultiply(a string, b string) string {
	real1, imag1, _ := parseComplex(a)
	real2, imag2, _ := parseComplex(b)
	real := real1*real2 - imag1*imag2 //i^2 = -1
	imag := real1*imag2 + real2*imag1
	return fmt.Sprintf("%d+%di", real, imag)
}
