package expense

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Expense struct {
	Id          string
	Date        string
	Category    string
	Description string
	Amount      string
}

func GetAllExpense() {
	file, err := os.Open("expense.csv")
	checkError(err)
	defer file.Close()
	res := csv.NewReader(file)
	data, err := res.ReadAll()
	checkError(err)
	var list []Expense
	for _, r := range data {
		list = append(list, Expense{
			Id:          r[0],
			Date:        r[1],
			Category:    r[2],
			Description: r[3],
			Amount:      r[4],
		})
	}
	list = list[1:]
	for _, val := range list {
		fmt.Println(strings.Repeat("-", 80))
		fmt.Printf("🧾 Expense Detail:\n")
		fmt.Printf("%-15s %-10s %-20s %-10s\n", "📅 Date", "🆔 ID", "📂 Category", "💵 Amount")
		fmt.Printf("%-15s %-10s %-20s %-10s\n", val.Date, val.Id, val.Category, val.Amount)
		fmt.Printf("\n✏️  Description: %s\n", val.Description)
		fmt.Println(strings.Repeat("-", 80))
	}
}

func GetIndividualExpense(id int) {
	file, err := os.Open("expense.csv")
	checkError(err)
	defer file.Close()
	res := csv.NewReader(file)
	data, err := res.ReadAll()
	checkError(err)
	var list []Expense
	for _, r := range data {
		list = append(list, Expense{
			Id:          r[0],
			Date:        r[1],
			Category:    r[2],
			Description: r[3],
			Amount:      r[4],
		})
	}
	list = list[1:]
	for _, val := range list {
		if val.Id == strconv.Itoa(id) {
			fmt.Println(strings.Repeat("-", 40))
			fmt.Println("🧾 Expense Details:")
			fmt.Printf("📅 Date        : %s\n", val.Date)
			fmt.Printf("🆔 Id          : %s\n", val.Id)
			fmt.Printf("📂 Category    : %s\n", val.Category)
			fmt.Printf("✏️  Description : %s\n", val.Description)
			fmt.Printf("💵 Amount      : %s\n", val.Amount)
			fmt.Println(strings.Repeat("-", 40))
			return
		}
	}
	fmt.Println("❌ Sorry, we couldn't find the expense with the ID you provided. Please double-check the ID and try again.")
}

func AddExpense(d string, c string, desp string, amt string) {
	file, err := os.Open("expense.csv")
	checkError(err)
	defer file.Close()
	res := csv.NewReader(file)
	data, err := res.ReadAll()
	checkError(err)
	var list []Expense
	for _, r := range data {
		list = append(list, Expense{
			Id:          r[0],
			Date:        r[1],
			Category:    r[2],
			Description: r[3],
			Amount:      r[4],
		})
	}
	list = list[1:]
	lastId := len(list)

	// writing section

	file, err = os.OpenFile("expense.csv", os.O_APPEND|os.O_WRONLY, 0644)
	checkError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	newExpense := []string{
		strconv.Itoa(lastId + 1),
		d,
		c,
		desp,
		amt,
	}

	err = writer.Write(newExpense)
	if err != nil {
		log.Fatalf("Error writing to CSV: %v", err)
	}
	fmt.Println("Expense added successfully!")
}

// edit expenses
// delete expenses

func checkError(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func ToCSVRow(e Expense) []string {
	return []string{e.Id, e.Date, e.Category, e.Description, e.Amount}
}

// Usage
