package worker

import (
	"fmt"
	"golang-mini-runner-project/internal/models"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecuteJob(job *models.Job) error {
	artifactsDir := "artifacts"
	if err := os.MkdirAll(artifactsDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create artifacts directory: %w", err)
	}

	logFileName := fmt.Sprintf("job_%d_log.txt", job.ID)
	logPath := filepath.Join(artifactsDir, logFileName)
	logFile, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}; defer logFile.Close()

	cmd := exec.Command("sh", "-c", job.Command)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	job.LogsPath = logPath

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("command execution failed; %w", err)
	}

	return nil
}