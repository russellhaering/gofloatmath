package main

import (
	"testing"
)

var (
	expected = float64(25.196850393700785)
)

func TestDoingMath(t *testing.T) {
	fns := map[string]func() float64{
		"FirstWay":  FirstWay,
		"SecondWay": SecondWay,
	}

	for name, fn := range fns {
		t.Run(name, func(t *testing.T) {
			result := fn()
			if result != expected {
				t.Error("expected", expected, "got", result)
			}
		})
	}
}
