package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const AppDirName = "My-CLI-Workspace-Data/Storage"

func GetDataPath(filename string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil { return "", fmt.Errorf("could not find home directory: %v", err) }

	appDirPath := filepath.Join(homeDir, "Downloads", AppDirName)
	if err := os.MkdirAll(appDirPath, 0755); err != nil {
		return "", fmt.Errorf("could not create data directory: %v", err)
	}

	return filepath.Join(appDirPath, filename), nil
}

func WriteJSON(filename string, data any) error {
	filePath, err := GetDataPath(filename)
	if err != nil { return err }

	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil { return fmt.Errorf("failed to encode JSON: %v", err) }
	
	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func ReadJSON(filename string, dest any) error {
	filePath, err := GetDataPath(filename)
	if err != nil { return nil }
	if _, err := os.Stat(filePath); os.IsNotExist(err) { return nil }

	fileData, err := os.ReadFile(filePath)
	if err != nil { return fmt.Errorf("failed to read file: %v", err) }
	
	if err := json.Unmarshal(fileData, dest); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err) 
	}

	return nil
}
