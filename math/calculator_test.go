package main

import (
	"testing"
)

func TestSumString(t *testing.T) {

	//mis casos
	testCase := []struct {
		name   string
		a      string
		b      string
		expect int
	}{
		{
			name:   "sumar dos numeros",
			a:      "5",
			b:      "4",
			expect: 9,
		},
		{
			name:   "El parametro a es incorrecto",
			a:      "a",
			b:      "4",
			expect: 0,
		},
		{
			name:   "El parametro b es incorrecto",
			a:      "5",
			b:      "b",
			expect: 0,
		},
	}

	for _, aux := range testCase {
		t.Run(aux.name, func(t *testing.T) {
			result := sumString(aux.a, aux.b)
			if result != aux.expect {
				t.Errorf("sum es incorrecto, got: %d, want: %d.", result, aux.expect)
			}
		})
	}

}
