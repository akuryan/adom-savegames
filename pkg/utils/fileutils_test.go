package utils

import (
    "os"
    "path/filepath"
    "testing"
)

// TestCopyFile tests the CopyFile function.
func TestCopyFile(t *testing.T) {
    src := "test_source.txt"
    dst := "test_destination.txt"

    // Create a test source file
    err := os.WriteFile(src, []byte("test content"), 0644)
    if err != nil {
        t.Fatalf("Failed to create source file: %v", err)
    }
    defer os.Remove(src) // Clean up

    // Call the CopyFile function
    err = CopyFile(src, dst)
    if err != nil {
        t.Fatalf("CopyFile() failed: %v", err)
    }
    defer os.Remove(dst) // Clean up

    // Check if the destination file exists
    if _, err := os.Stat(dst); os.IsNotExist(err) {
        t.Fatalf("Destination file does not exist: %v", err)
    }

    // Check if the content is the same
    srcContent, _ := os.ReadFile(src)
    dstContent, _ := os.ReadFile(dst)
    if string(srcContent) != string(dstContent) {
        t.Fatalf("Content mismatch: expected %s, got %s", string(srcContent), string(dstContent))
    }
}

// TestDeleteFile tests the DeleteFile function.
func TestDeleteFile(t *testing.T) {
    file := "test_file.txt"

    // Create a test file
    err := os.WriteFile(file, []byte("test content"), 0644)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }

    // Call the DeleteFile function
    err = DeleteFile(file)
    if err != nil {
        t.Fatalf("DeleteFile() failed: %v", err)
    }

    // Check if the file has been deleted
    if _, err := os.Stat(file); !os.IsNotExist(err) {
        t.Fatalf("File was not deleted: %v", err)
    }
}

// TestGetFilesInDirectory tests the GetFilesInDirectory function.
func TestGetFilesInDirectory(t *testing.T) {
    dir := "test_directory"
    os.Mkdir(dir, 0755)
    defer os.RemoveAll(dir) // Clean up

    // Create test files
    for i := 0; i < 3; i++ {
        os.WriteFile(filepath.Join(dir, "file"+string(i)+".txt"), []byte("test content"), 0644)
    }

    files, err := GetFilesInDirectory(dir)
    if err != nil {
        t.Fatalf("GetFilesInDirectory() failed: %v", err)
    }

    if len(files) != 3 {
        t.Fatalf("Expected 3 files, got %d", len(files))
    }
}