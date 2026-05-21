package cmd

import (
	"3-My-CLI-Workspace/internal/services"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	expenseCmd.AddCommand(expenseAddCmd, expenseListCmd)
	RootCmd.AddCommand(expenseCmd)
}

var expenseCmd = &cobra.Command{
	Use: "expense",
	Short: "Manage your daily expenses",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var expenseAddCmd = &cobra.Command{
	Use: "add [amount] [note]",
	Short: "Add a new expense",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string)  {
		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil { cmd.PrintErrln("Error: Amount must be a valid number"); return }
		note := args[1]
		services.AddExpense(amount, note)
	},
}

var expenseListCmd = &cobra.Command{
	Use: "list",
	Short: "List all expenses and show total",
	Run: func (cmd *cobra.Command, args []string) {
		services.ListExpense()
	},
}