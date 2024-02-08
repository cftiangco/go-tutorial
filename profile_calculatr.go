package main

import (
	"fmt"
	"errors"
	"os"
)

func main() {
	revenue, errRevenue := getUserInput("Revenue: ")
	expenses, errExpenses := getUserInput("Expenses: ")
	taxRate, errTaxRate := getUserInput("Tax Rate: ")

	if errRevenue != nil || errExpenses != nil || errTaxRate != nil {
		fmt.Println("Error")
		fmt.Println(errRevenue,errExpenses,errTaxRate)
		panic("Can't continue")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	storeCalculatedResults(ebt,profit,ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if(userInput <= 0) {
		return 0, errors.New("Input must be greater than 0")
	}

	return userInput, nil
}

func storeCalculatedResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n",ebt,profit,ratio)
	os.WriteFile("CalculatedResultFile.txt", []byte(results),0644)
}