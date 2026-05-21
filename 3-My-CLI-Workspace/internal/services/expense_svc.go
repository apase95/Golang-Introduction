package services

import (
	"3-My-CLI-Workspace/internal/models"
	"3-My-CLI-Workspace/internal/storage"
	"fmt"
	"time"

	"github.com/fatih/color"
)


const ExpenseFile = "expenses.json"

func formatMoney(amount float64) string {
	return fmt.Sprintf("%.0f VND", amount)
}

func AddExpense(amount float64, note string) {
	var expenses []models.Expense
	_ = storage.ReadJSON(ExpenseFile, &expenses)

	id := 1
	if len(expenses) > 0 { id = expenses[len(expenses) - 1].ID + 1 }

	newExpense := models.Expense{
		ID: 		id,
		Amount: 	amount,
		Note: 		note,
		CreatedAd: 	time.Now(),
	}
	expenses = append(expenses, newExpense)
	_ = storage.WriteJSON(ExpenseFile, expenses)
	color.Green("✔ Expense added: %s - %s\n", formatMoney(amount), note)
}

func ListExpense() {
	var expenses []models.Expense
	_ = storage.ReadJSON(ExpenseFile, &expenses)
	
	if len(expenses) == 0 { color.Yellow("No expenses recorded yet! \n"); return }
	
	var total float64 = 0
	fmt.Println("====== YOUR EXPENSES ======")
	for _, e := range expenses {
		total += e.Amount
		dateStr := e.CreatedAd.Format("02/06")
		fmt.Printf("[%02d] %s | %s - %s\n", e.ID, color.CyanString(dateStr), color.RedString(formatMoney(e.Amount)), e.Note)
	}
	fmt.Println("---------------------------")
	color.HiGreen("💰 TOTAL SPENT: %s\n", formatMoney(total))
}