package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	foo = 10
)

func FirstWay() float64 {
	a := float64(6) / float64(127)
	b := float64(130 - foo*2)

	return 20 + b*a
}

func SecondWay() float64 {
	var foo = 10

	a := float64(6) / float64(127)
	b := float64(130 - foo*2)

	return 20 + b*a
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
