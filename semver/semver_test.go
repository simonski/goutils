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
		t.Errorf("Minor should be 1")
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

func Test_SemVerEq(t *testing.T) {
	a, _ := New("1.0.0")
	b, _ := New("1.0.0")
	if !a.Eq(b) {
		t.Errorf("a should eq b\n")
	}
}

func Test_SemVerGt(t *testing.T) {
	a, _ := New("1.0.1")
	b, _ := New("1.0.0")
	if !a.Gt(b) {
		t.Errorf("a should gt b\n")
	}
}

func Test_SemVerLt(t *testing.T) {
	a, _ := New("1.0.0")
	b, _ := New("1.0.1")
	if !a.Lt(b) {
		t.Errorf("a should lt b\n")
	}
}

func Test_SemVerGt2(t *testing.T) {
	a, _ := New("1.0.4")
	b, _ := New("1.0.10")
	if a.Gt(b) {
		t.Errorf("a should not be gt b\n")
	}
}
