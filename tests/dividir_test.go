package test

import (
	"testing"

	"github.com/radikaledward1/go-unit-test/utils"
)

func TestDividir(t *testing.T) {
	casos := []struct {
		a, b        float64
		esperado    float64
		seEsperaErr bool
	}{
		{10, 2, 5, false},
		{10, 0, 0, true}, // Debería devolver un error
	}

	for _, caso := range casos {
		resultado, err := utils.Dividir(caso.a, caso.b)

		if caso.seEsperaErr && err == nil {
			t.Errorf("Se esperaba un error al dividir %f por %f, pero no se obtuvo error", caso.a, caso.b)
		}

		if !caso.seEsperaErr && resultado != caso.esperado {
			t.Errorf("Para %f / %f; se esperaba %f, pero se obtuvo %f", caso.a, caso.b, caso.esperado, resultado)
		}
	}
}
