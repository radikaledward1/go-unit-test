package utils

import "errors"

func Dividir(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	c := a / b
	println(c)
	return c, nil
}
