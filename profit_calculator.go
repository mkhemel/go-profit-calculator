package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//var revenue float64
	//var expenses float64
	//var taxRate float64

	revenue, err := getUserInput("Revenue")
	if err != nil {
		println(err)
		return
	}
	expenses, err := getUserInput("Expenses")
	if err != nil {
		println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate")
	if err != nil {
		println(err)
		return
	}

	ebt, profit, ratio := calculateFunctions(revenue, expenses, taxRate)

	fmt.Printf("EBT : %.1f \n", ebt)

	fmt.Printf("Profit : %.1f\n", profit)

	fmt.Printf("Ratio : %.1f\n", ratio)
	storeResult(ebt, profit, ratio)

}

func calculateFunctions(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func storeResult(ebt, profit, ratio float64) {
	resultText := fmt.Sprintf("Ebt: %.2f \n Profit: %.2f \n Ratio: %.2f \n", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(resultText), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText, " : ")
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("value must be positive numbers")
	}
	return userInput, nil
}
