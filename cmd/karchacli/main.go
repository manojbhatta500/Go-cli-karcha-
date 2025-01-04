package main

import (
	"bufio"
	"fmt"
	"karcha/expense"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	} else {
		if os.Args[1] != "open" {
			fmt.Println("wrong password, try again ;")
			fmt.Println(os.Args)
		}
	}
	startingText()
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("🌟 Welcome to the Expense Tracker 🌟")
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("Choose an option:")
		fmt.Println("1️⃣  See Expenses")
		fmt.Println("2️⃣  Add Expenses")
		fmt.Println("3️⃣   Exit Program")
		var input int
		fmt.Scanf("%d", &input)
		if input == 1 {
			showExpenseMenu()
		} else if input == 2 {
			addExpense()
		} else if input == 3 {
			break
		} else {
			fmt.Println("plese select the correct option")
		}
	}
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("🙏  Good Bye! See You Again Soon!  👋")
	fmt.Println(strings.Repeat("-", 40))

}

func showExpenseMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("💰 Expense Tracker Menu")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("1️⃣  View all expenses")
		fmt.Println("2️⃣  View an individual expense")
		fmt.Println("0️⃣  Exit")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Print("Enter your choice: ")

		// Take user input
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("❌ Invalid input. Please enter a valid number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("📋 Displaying all expenses...")
			expense.GetAllExpense()
		case 2:
			fmt.Print("🔍 Enter the expense ID to view: ")
			var id int
			_, err := fmt.Scanf("%d", &id)
			if err != nil {
				fmt.Println("❌ Invalid input. Please enter a valid ID.")
				continue
			}
			fmt.Printf("🧾 Showing details for Expense ID: %d\n", id)
			expense.GetIndividualExpense(id)
		case 0:
			fmt.Println("👋 Goodbye! Thanks for using the Expense Tracker.")
			return
		default:
			fmt.Println("❌ Sorry, that's not a valid choice. Try again.")
		}
	}
}

func startingText() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("💰 Welcome to the Expense Tracker CLI Tool 💻")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("📋  This is an Expense App: A CLI Tool and Server")
	fmt.Println("🛠️   Created by Manoj Bhatta for Personal Use")
	fmt.Println(strings.Repeat("=", 50))
}

func addExpense() {
	var desc string
	var price string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("========================================")
		fmt.Println("🌟 Please enter your expense description:")
		fmt.Println("========================================")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading description:", err)
			continue
		}
		desc = string(input)
		desc = strings.TrimSpace(desc)

		fmt.Println("========================================")
		fmt.Println("🌟 Please enter the expense amount:")
		fmt.Println("========================================")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading price:", err)
			continue
		}
		amount, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
		if err != nil {
			log.Fatal("Error converting amount:", err)
			continue
		}
		price = strconv.FormatFloat(amount, 'f', 2, 64)
		break
	}

	timenow := time.Now().Format("2006-01-02")
	expense.AddExpense(timenow, "others", desc, price)
}
