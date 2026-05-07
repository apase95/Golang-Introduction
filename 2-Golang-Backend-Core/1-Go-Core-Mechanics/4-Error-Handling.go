package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrTimeout = errors.New("job execution timeout")
	ErrNoDiskSpace = errors.New("no disk space left on device")
)

type JobError struct {
	JobID 		string
	ExitCode 	int
	Message 	string
}

func (e *JobError) Error() string {
	return fmt.Sprintf("[Job %s failed] Code: %d, Msg: %s", e.JobID, e.ExitCode, e.Message)
}

func executeScript(scriptName string) error {
	if scriptName == "sleep.sh" {
		return ErrTimeout
	}
	if scriptName == "missing.sh" {
		return os.ErrNotExist
	}
	return nil
}

func WorkerRun(jobID string, script string) error {
	err := executeScript(script)
	if err != nil {
		return fmt.Errorf("WorkerRun failed to execute job %s: %w", jobID, err)
	}
	if script == "crash.sh" {
		return &JobError{
			JobID: 		jobID,
			ExitCode: 	137,
			Message: 	"Out of memory crash",
		}
	}
	return nil
}

func main() {
	fmt.Println("--- SCENARIO 1: ERROR WRAPPING & errors.Is ---")
	err1 := WorkerRun("JOB-01", "sleep.sh")
	if err1 != nil {
		fmt.Println("Print Error normally:\n ->", err1)
		
		if errors.Is(err1, ErrTimeout) {
			fmt.Println("Error Analysis: System detected this is a Timeout error! Need to push to Retry Queue.")
		}
	}

	fmt.Println("\n--- SCENARIO 2: SYSTEM ERROR (os.ErrNotExist) ---")
	err2 := WorkerRun("JOB-02", "missing.sh")
	if err2 != nil {
		fmt.Println("Print Error normally:\n ->", err2)
		
		if errors.Is(err2, os.ErrNotExist) {
			fmt.Println("Error Analysis: Script file does not exist on the system!")
		}
	}

	fmt.Println("\n--- SCENARIO 3: CUSTOM ERROR & errors.As ---")
	err3 := WorkerRun("JOB-03", "crash.sh")
	if err3 != nil {
		fmt.Println("Print Error normally:\n ->", err3)

		var jobErr *JobError
		if errors.As(err3, &jobErr) {
			fmt.Printf("Error Analysis [Custom Error Extracted]:\n - Job ID: %s\n - Exit Code: %d (Alert DevOps immediately!)\n - Details: %s\n", 
				jobErr.JobID, jobErr.ExitCode, jobErr.Message)
		}
	}
}