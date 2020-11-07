package main

import (
	"testing"
)

func TestAll(t *testing.T) {

	c := GetConfig()

	maTest(c.Ma)

	mpTest(c.Mp)

	payTest(c.Pay)

}
