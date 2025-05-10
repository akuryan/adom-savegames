package restore

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// RestoreFile restores a deleted file from the backup after a specified delay.
func RestoreFile(backupDir, fileName string, delay time.Duration) error {
	// Wait for the specified delay before restoring the file
	time.Sleep(delay)

	// Construct the path to the backup file
	backupFilePath := filepath.Join(backupDir, fileName)

	// Check if the backup file exists
	if _, err := os.Stat(backupFilePath); os.IsNotExist(err) {
		return fmt.Errorf("backup file does not exist: %s", backupFilePath)
	}

	// Read the backup file
	data, err := ioutil.ReadFile(backupFilePath)
	if err != nil {
		return fmt.Errorf("failed to read backup file: %v", err)
	}

	// Restore the file to its original location
	originalFilePath := filepath.Join(filepath.Dir(backupDir), fileName)
	err = ioutil.WriteFile(originalFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to restore file: %v", err)
	}

	return nil
}