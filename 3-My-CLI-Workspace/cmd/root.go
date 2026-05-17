package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var RootCmd = &cobra.Command{
	Use: 	"workspace",
	Short: 	"My-CLI-Workspace is a minimalist task and life manager",
	Long: 	`A fast, offline, and minimalist Command Line Interface (CLI) application designed for developers to manage tasks, expenses, and local music directly from the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to My-CLI-Workspace!")
		fmt.Println("Type 'workspace --help' to see available commands.")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}