package gofloatmath

import (
	"strconv"
	"testing"
)

var (
	foo = struct {
		Bar int
	}{
		Bar: 10,
	}
)

func TestDoingMath(t *testing.T) {
	expected, err := strconv.ParseFloat("25.196850393700785", 64)
	if err != nil {
		t.Error(err)
	}

	baz := float64(6) / float64(127)
	x := float64(10)*2 + float64(130-foo.Bar*2)*baz

	if x != expected {
		t.Error("Expected", expected, "got", x)
	}
}
