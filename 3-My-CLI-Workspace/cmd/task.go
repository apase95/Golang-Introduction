package cmd

import (
	"3-My-CLI-Workspace/internal/services"
	"strconv"

	"github.com/spf13/cobra"
)


var taskCmd = &cobra.Command{
	Use: 	"task",
	Short: 	"Manage your to-do tasks",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var taskAddCmd = &cobra.Command{
	Use: 	"add [title]",
	Short: 	"Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		services.AddTask(args[0])
	},
}

var taskListCmd = &cobra.Command{
	Use: 	"list",
	Short: 	"List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		services.ListTasks()
	},
}

var taskDoneCmd = &cobra.Command{
	Use: 	"done [id]",
	Short: 	"Toggle task status (Done/Undone)",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil { cmd.PrintErrln("Error: ID must be a number"); return }
		services.ToggleTaskStatus(id)
	},
}

var taskDelCmd = &cobra.Command{
	Use: 	"del [id]",
	Short: 	"Delete a task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil { cmd.PrintErrln("Error: ID must be a number"); return }
		services.DeleteTask(id)
	},
}

func init() {
	taskCmd.AddCommand(taskAddCmd, taskListCmd, taskDoneCmd, taskDelCmd)
	RootCmd.AddCommand(taskCmd)
}