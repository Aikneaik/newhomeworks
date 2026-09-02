package main

import (
	"errors"
	"fmt"
)

const (
	usdEur = 0.94
	usdRub = 94.0
)

func main() {
	var (
		firstCurr  string
		secondCurr string
		value      int
		err        error
	)

	for {
		firstCurr, err = checkCurrency("Введи исходную валюту в формате: usd, rub, eur", "")
		if err != nil {
			fmt.Println("Warnin:", err)
			continue
		}
		break
	}

	for {
		value, err = checkValue("Введите сумму для конвертации. Целое в формате 1, 100, 2000.")
		if err != nil {
			fmt.Println("Warning:", err)
			continue
		}
		break
	}

	for {
		secondCurr, err = checkCurrency("Введи целевую валюту в формате: usd, rub, eur", firstCurr)
		if err != nil {
			fmt.Println("Warning:", err)
			continue
		}
		break
	}

	fmt.Printf("При конвертации %d %s в %s получится %.2f ", value, firstCurr, secondCurr, changeMoney(firstCurr, value, secondCurr))

}

func changeMoney(firstCurr string, value int, secondCurr string) float64 {
	var equalUsd float64

	switch firstCurr {
	case "usd":
		equalUsd = float64(value) * 1.0
	case "eur":
		equalUsd = float64(value) / usdEur
	case "rub":
		equalUsd = float64(value) * usdRub
	}

	switch secondCurr {
	case "usd":
		return equalUsd * 1
	case "eur":
		return equalUsd * usdEur
	case "rub":
		return equalUsd * usdRub
	default:
		return 0
	}

}

func checkValue(prompt string) (int, error) {
	var value int

	fmt.Println(prompt)
	_, err := fmt.Scanln(&value)

	if err != nil {
		fmt.Println("Wrong entry", err)
		return 0, err
	}
	return value, nil
}

func checkCurrency(prompt, exclude string) (string, error) {

	var currency string

	fmt.Println(prompt)

	_, err := fmt.Scanln(&currency)
	if err != nil {
		fmt.Println("Wrong entry", err)
		return "", err
	} else if currency == exclude {
		return "", errors.New("can't choose the same currency")
	}

	switch currency {
	case "usd", "rub", "eur":
		return currency, nil
	default:
		return "", errors.New("Incorrect entry, please choose from the example")
	}

}

// Финализируем приложение калькулятор. Для этого:

// Сделать меню с шагами
// Ввод исходной валюты (подсказываем варианты) - если ошибка, заново вводим
// Ввод числа - если ошибка, заново вводим
// Ввод целевой валюты (подсказываем варианты) - если ошибка, заново вводим
// Выделить функцию ввода / проверки валюты и числа

// После получения всех данных с помощью if / switch вычислить итог и вывести результат.
