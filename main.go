package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Authentication struct {
	userName string
	password string
}

func main() {

	authData := Authentication{userName: "manojbhatta", password: "manoj"}

	for {
		fmt.Println("Welcome to Expense App.")
		fmt.Print("Username: ")
		userNameData := userInput()
		fmt.Print("Passkey: ")
		userPasskeyData := userInput()

		userData := Authentication{userName: userNameData, password: userPasskeyData}

		result := checkPassword(&authData, &userData)

		if result {
			fmt.Println("Successfully Logged in.")
			break
		} else {
			fmt.Println("sorry Authentication didn't matched")
		}

	}

	fmt.Println("options:")
	fmt.Println("1. Add Expenses")
	fmt.Println("2. Read Expenses")

	fmt.Print(":")
	userChoiceInput := userInput()

	userChoiceInputNumber, err := strconv.ParseInt(userChoiceInput, 10, 64)

	handleError(err)

	switch userChoiceInputNumber {
	case 1:
		addExpenseToDataFile()
	case 2:
		readExpenseFromDataFile("./data.txt")

	default:
		fmt.Println("in future Updates.")

	}

}

func addExpenseToDataFile() {
	file, err := os.OpenFile("./data.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	handleError(err)
	defer file.Close()

	fmt.Print("Expense : ")

	rawExpenseData := userInput()
	currentTime := time.Now()
	currentTimeFormatted := currentTime.Format("2006-01-02 15:04:05")

	fmt.Print("Amount : ")

	amount := userInput()
	cleanExpenseData := "Date Time:" + currentTimeFormatted + ", Expense: " + rawExpenseData + " ,amount:" + amount + "\n"

	_, err = io.WriteString(file, cleanExpenseData)
	handleError(err)
}

func readExpenseFromDataFile(filename string) {
	dataBytes, err := os.ReadFile(filename)
	handleError(err)
	fmt.Println(string(dataBytes))
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func checkPassword(a, b *Authentication) bool {
	return a.userName == b.userName && a.password == b.password
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
