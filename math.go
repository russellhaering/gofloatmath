package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	foo = 100
)

func FirstWay() float64 {
	return float64(foo) * (float64(6) / float64(127))
}

func SecondWay() float64 {
	var foo = 10

	return float64(foo) * (float64(6) / float64(127))
}

func PrintCPUInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	spew.Dump(cpuInfo)
}

func main() {
	PrintCPUInfo()

	fmt.Println("-- Doing Math --")
	fmt.Println("First Way: ", FirstWay())
	fmt.Println("Second Way: ", SecondWay())
}
