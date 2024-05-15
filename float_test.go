package main

import (
	"testing"
)

var (
	// expected is the value returned most often by the function(s) in testing, I've made
	// no attempt to confirm that it's actually "correct".
	expected = float64(25.196850393700785)
)

func TestDoingMath(t *testing.T) {
	PrintCPUInfo()

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
