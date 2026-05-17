package services

import (
	"3-My-CLI-Workspace/internal/models"
	"3-My-CLI-Workspace/internal/storage"
	"fmt"
	"time"

	"github.com/fatih/color"
)

const TaskFile = "tasks.json"


func AddTask(title string) {
	var tasks []models.Task
	_ = storage.ReadJSON(TaskFile, &tasks)

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks) - 1].ID + 1
	}

	newTask := models.Task{
		ID: 		id,
		Title: 		title,
		Done: 		false,
		CreatedAt: 	time.Now(),
	}
	tasks = append(tasks, newTask)
	_ = storage.WriteJSON(TaskFile, tasks)

	color.Green("Task added successfully: %s (ID: %d)\n", title, id)
}


func ListTasks() {
	var tasks []models.Task
	_ = storage.ReadJSON(TaskFile, &tasks)
	if len(tasks) == 0 {
		color.Yellow("No tasks found!")
		return
	}

	fmt.Println("\n === YOUR TASKS ===")
	for _, t := range tasks {
		if t.Done {
			color.Green("  [%d] [x] %s", t.ID, t.Title)
		} else {
			color.Red("  [%d] [ ] %s", t.ID, t.Title)
		}
	}
	fmt.Println()
}


func ToggleTaskStatus(id int) {
	var tasks []models.Task
	_ = storage.ReadJSON(TaskFile, &tasks)
	
	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = !tasks[i].Done
			found = true
			if tasks[i].Done {
				color.Green("Task %d marked as DONE!\n", id)
			} else {
				color.Yellow("Task %d marked as UNDONE!\n", id)
			}; break
		}
	}; if !found { color.Red("Task ID %d not found.\n", id); return }
_ = storage.WriteJSON(TaskFile, tasks)
}


func DeleteTask(id int) {
	var tasks []models.Task
	_ = storage.ReadJSON(TaskFile, &tasks)

	var newTasks []models.Task
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}; newTasks = append(newTasks, t)
	};if !found { color.Red("Task ID %d not found. \n", id); return }

	_ = storage.WriteJSON(TaskFile, newTasks)
	color.Green("Task %d deleted successfully!\n", id)
}