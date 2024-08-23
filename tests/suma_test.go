package test

import (
	"testing"

	"github.com/radikaledward1/go-unit-test/utils" // import the package that contains the function to test
)

func TestSuma(t *testing.T) {
	result := utils.Suma(1, 2)
	esperado := 3

	if result != esperado {
		t.Errorf("Suma(1,2) returns %d desired %d", result, esperado)
	}
}

func TestSumasMultiples(t *testing.T) {
	casos := []struct {
		a, b, esperado int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, -1, -2},
		{-2, 3, 1},
	}

	for _, caso := range casos {
		resultado := utils.Suma(caso.a, caso.b)
		if resultado != caso.esperado {
			t.Errorf("Suma(%d + %d) returns %d, desired %d", caso.a, caso.b, resultado, caso.esperado)
		}
	}
}
