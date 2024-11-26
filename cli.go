package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	Timestamp string  `json:"timestamp"`
}

var expenses []Expense
var fileName = "expenses.json"

func main() {
	loadExpenses()
	fmt.Println("Welcome to Expense Tracker!")
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. View Total Expenses")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			addExpense()
		case 2:
			viewExpenses()
		case 3:
			viewTotalExpenses()
		case 4:
			saveExpenses()
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

func addExpense() {
	var amount float64
	var category string

	fmt.Print("Enter expense amount: ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Invalid amount.")
		return
	}

	fmt.Print("Enter expense category: ")
	fmt.Scan(&category)

	newExpense := Expense{
		Amount:    amount,
		Category:  category,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
	expenses = append(expenses, newExpense)
	fmt.Println("Expense added successfully!")
}

func viewExpenses() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}
	fmt.Println("\nExpenses:")
	for i, exp := range expenses {
		fmt.Printf("%d. %s - %.2f (%s)\n", i+1, exp.Category, exp.Amount, exp.Timestamp)
	}
}

func viewTotalExpenses() {
	var total float64
	for _, exp := range expenses {
		total += exp.Amount
	}
	fmt.Printf("\nTotal Expenses: %.2f\n", total)
}

func saveExpenses() {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(expenses)
	if err != nil {
		fmt.Println("Error encoding expenses:", err)
	}
}

func loadExpenses() {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error loading expenses:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&expenses)
	if err != nil {
		fmt.Println("Error decoding expenses:", err)
	}
}
