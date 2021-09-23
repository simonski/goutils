package goutils

import (
	"testing"
)

func TestLPad(t *testing.T) {
	line1 := "hello"
	actual := LPad(line1, " ", 5)
	expect := "     hello"
	if actual != expect {
		t.Errorf("LPad did not pad correctly - expected '%v' actual '%v'\n", expect, actual)
	}

}

func TestRPad(t *testing.T) {
	line1 := "hello"
	actual := RPad(line1, " ", 5)
	expect := "hello     "
	if actual != expect {
		t.Errorf("LPad did not pad correctly - expected '%v' actual '%v'\n", expect, actual)
	}

}

func TestRPadToFixedLength(t *testing.T) {
	line1 := "hello"
	actual := RPadToFixedLength(line1, " ", 6)
	expect := "hello "
	if actual != expect {
		t.Errorf("LPad did not pad correctly - expected '%v' actual '%v'\n", expect, actual)
	}

}

func TestLPadToFixedLength(t *testing.T) {
	line1 := "hello"
	actual := LPadToFixedLength(line1, " ", 6)
	expect := " hello"
	if actual != expect {
		t.Errorf("LPad did not pad correctly - expected '%v' actual '%v'\n", expect, actual)
	}

}
func TestLPadToFixedLengthAtMaxAlready(t *testing.T) {
	line1 := "hello"
	actual := LPadToFixedLength(line1, " ", 5)
	expect := "hello"
	if actual != expect {
		t.Errorf("LPadToFixedLength did not pad correctly - expected '%v' actual '%v'\n", expect, actual)
	}

}

func TestRPadToFixedLengthAtMaxAlready(t *testing.T) {
	line1 := "hello"
	actual := RPadToFixedLength(line1, " ", 5)
	expect := "hello"
	if actual != expect {
		t.Errorf("RPadToFixedLength did not pad correctly - expected '%v' actual '%v'\n", expect, actual)
	}

}

func TestTokenSwitchEnvironmentVariables(t *testing.T) {
	expected := "/Users/simongauld"
	actual := TokenswitchEnvironmentVariables("~")
	if actual != expected {
		t.Fatalf("%v != %v\n", actual, expected)
	}
}
