package util

import (
	"testing"
)

func TestToMap(t *testing.T) {
	a := struct {
		A string
		B uint64
		C float64
		D struct {
			D string
		}
	}{
		A: "1",
		B: 1,
		C: 1.1,
		D: struct{ D string }{D: "D"},
	}
	t.Log(ToMap(a))
}
