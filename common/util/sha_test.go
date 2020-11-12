package util

import "testing"

func TestCheckSignature(t *testing.T) {
	t.Log(CheckSignature("1", "1", "1", "1"))
}
