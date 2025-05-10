package backup

import (
	"os"
	"testing"
)

func TestBackupFile(t *testing.T) {
	// Setup: Create a temporary directory for testing
	tempDir := os.TempDir()
	testFile := tempDir + "/testfile.txt"
	backupFile := tempDir + "/backup/testfile.txt"

	// Create a test file
	err := os.WriteFile(testFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile) // Clean up

	// Call the function to back up the file
	err = BackupFile(testFile, backupFile)
	if err != nil {
		t.Fatalf("BackupFile() error = %v", err)
	}

	// Check if the backup file exists
	if _, err := os.Stat(backupFile); os.IsNotExist(err) {
		t.Fatalf("Backup file does not exist: %v", err)
	}

	// Clean up the backup file
	defer os.Remove(backupFile)
}

func TestRestoreFile(t *testing.T) {
	// Setup: Create a temporary directory for testing
	tempDir := os.TempDir()
	backupFile := tempDir + "/backup/testfile.txt"
	restoredFile := tempDir + "/restored/testfile.txt"

	// Create a backup file
	err := os.MkdirAll(tempDir+"/backup", 0755)
	if err != nil {
		t.Fatalf("Failed to create backup directory: %v", err)
	}
	err = os.WriteFile(backupFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create backup file: %v", err)
	}
	defer os.Remove(backupFile) // Clean up

	// Call the function to restore the file
	err = RestoreFile(backupFile, restoredFile)
	if err != nil {
		t.Fatalf("RestoreFile() error = %v", err)
	}

	// Check if the restored file exists
	if _, err := os.Stat(restoredFile); os.IsNotExist(err) {
		t.Fatalf("Restored file does not exist: %v", err)
	}

	// Clean up the restored file
	defer os.Remove(restoredFile)
}