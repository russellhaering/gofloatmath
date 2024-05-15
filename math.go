package main

import "fmt"

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

func main() {
	fmt.Println("First Way: ", FirstWay())
	fmt.Println("Second Way: ", SecondWay())
}
