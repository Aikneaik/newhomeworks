package main

import "fmt"

const (
	usdEur = 0.85
	usdRub = 85
)

func main() {
	eurRub := usdRub / usdEur
	fmt.Println(eurRub)
}
