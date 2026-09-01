package main

import "fmt"

const (
	usdEur         = 0.85
	usdRub         = 85
	eurRub float64 = usdRub / usdEur
)

func main() {
	fmt.Println(eurRub)
}
