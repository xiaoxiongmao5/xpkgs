package xtime

import (
	"fmt"
	"testing"
)

func TestParseTimeString(t *testing.T) {
	t1, err := ParseTimeString("2021-01-01")
	if err != nil {
		t.Error(err)
	}
	t2, err := ParseTimeString("2021-01-01 00:00:00")
	if err != nil {
		t.Error(err)
	}
	t3, err := ParseTimeString("2021-01-01T00:00:00Z")
	if err != nil {
		t.Error(err)
	}
	t4, err := ParseTimeString("2021-01-01T00:00:00+08:00")
	if err != nil {
		t.Error(err)
	}
	t5, err := ParseTimeString("2021-1-1")
	if err != nil {
		t.Error(err)
	}
	t6, err := ParseTimeString("2021-1-1 00:00:00")
	if err != nil {
		t.Error(err)
	}
	t7, err := ParseTimeString("2021-1-1T00:00:00Z")
	if err != nil {
		t.Error(err)
	}
	t8, err := ParseTimeString("2021-1-1T00:00:00+08:00")
	if err != nil {
		t.Error(err)
	}
	if !Equal(t, t1, t2) {
		fmt.Println("t1 != t2")
		return
	}

	if !Equal(t, t1, t3) {
		fmt.Println("t1 != t3")
		return
	}

	if !Equal(t, t1, t5) {
		fmt.Println("t1 != t5")
		return
	}

	if !Equal(t, t1, t6) {
		fmt.Println("t1 != t6")
		return
	}

	if !Equal(t, t1, t7) {
		fmt.Println("t1 != t7")
		return
	}

	if !Equal(t, t4, t8) {
		fmt.Println("t4 != t8")
		return
	}
}

func Equal(t *testing.T, a, b interface{}) bool {
	if a != b {
		t.Errorf("not equal, %v != %v", a, b)
		return false
	}
	return true
}
