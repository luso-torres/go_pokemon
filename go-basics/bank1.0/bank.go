package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	fmt.Println("Welcome to Go Bank!") // the structure for finite loops is: for i:=0; i<2; i++{ }
	// besides that execution, there is for someCondition, which holds the loop until the bool variable assumes False.
	for {
		var accountBalance, err = getBalanceFromFile()

		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("--------------")
			panic("Can't continue. Sorry")
		}

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		//fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Deposit value:")
			var depositValue float64
			fmt.Scan(&depositValue)
			if depositValue <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return //stops the execution of the function thereafter.
				continue
			}
			accountBalance += depositValue
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount ", accountBalance)
		case 3:
			fmt.Print("Withdrawal amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("You dont have enough balance to do this operation.")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Print("Goodbye!!")
			fmt.Print("Thanks for choosing our bank!")
			return
		}
	} //fmt.Println("Your choice: ", choice)
}
