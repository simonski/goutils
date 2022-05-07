package semver

import (
	"testing"
)

func Test_SemVerGood(t *testing.T) {
	value := "1.2.3"
	sv, err := New(value)
	if err != nil {
		t.Errorf("Failed to parse: %v, : %v", value, err.Error())
	}
	if sv.Major != 1 {
		t.Errorf("Major should be 1")
	}
	if sv.Minor != 2 {
		t.Errorf("Minro should be 1")
	}
	if sv.Increment != 3 {
		t.Errorf("Increment should be 3")
	}
	if sv.Value != "1.2.3" {
		t.Errorf("Value should be 1.2.3")
	}
}

func Test_SemVer(t *testing.T) {
	value := "a.b.c"
	_, err := New(value)
	if err == nil {
		t.Errorf("Shoud have failed parsing: %v", err.Error())
	}
}
