package restore

import (
	"os"
	"testing"
	"time"
)

func TestRestoreFile(t *testing.T) {
	// Setup: Create a temporary backup file
	backupFile := "test_backup.txt"
	originalFile := "test_file.txt"
	content := []byte("This is a test file.")

	err := os.WriteFile(originalFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create original file: %v", err)
	}
	defer os.Remove(originalFile)

	err = os.WriteFile(backupFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create backup file: %v", err)
	}
	defer os.Remove(backupFile)

	// Test: Restore the file from backup
	err = RestoreFile(originalFile, backupFile)
	if err != nil {
		t.Fatalf("Failed to restore file: %v", err)
	}

	// Verify: Check if the original file content matches the backup
	restoredContent, err := os.ReadFile(originalFile)
	if err != nil {
		t.Fatalf("Failed to read restored file: %v", err)
	}

	if string(restoredContent) != string(content) {
		t.Errorf("Restored file content does not match backup. Got: %s, Want: %s", restoredContent, content)
	}

	// Cleanup: Remove the restored file
	os.Remove(originalFile)
}

func TestRestoreFileDelay(t *testing.T) {
	// Setup: Create a temporary backup file
	backupFile := "test_backup.txt"
	originalFile := "test_file.txt"
	content := []byte("This is a test file.")

	err := os.WriteFile(originalFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create original file: %v", err)
	}
	defer os.Remove(originalFile)

	err = os.WriteFile(backupFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create backup file: %v", err)
	}
	defer os.Remove(backupFile)

	// Simulate deletion of the original file
	os.Remove(originalFile)

	// Test: Restore the file from backup after a delay
	time.Sleep(1 * time.Second) // Simulate delay
	err = RestoreFile(originalFile, backupFile)
	if err != nil {
		t.Fatalf("Failed to restore file: %v", err)
	}

	// Verify: Check if the original file content matches the backup
	restoredContent, err := os.ReadFile(originalFile)
	if err != nil {
		t.Fatalf("Failed to read restored file: %v", err)
	}

	if string(restoredContent) != string(content) {
		t.Errorf("Restored file content does not match backup. Got: %s, Want: %s", restoredContent, content)
	}

	// Cleanup: Remove the restored file
	os.Remove(originalFile)
}