package utils

import (
    "io"
    "os"
    "path/filepath"
)

// CopyFile copies a file from src to dst. If dst does not exist, it will be created.
// If dst exists, it will be overwritten.
func CopyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    _, err = io.Copy(destinationFile, sourceFile)
    if err != nil {
        return err
    }

    return destinationFile.Sync()
}

// DeleteFile deletes the specified file.
func DeleteFile(filePath string) error {
    return os.Remove(filePath)
}

// FileExists checks if a file exists at the given path.
func FileExists(filePath string) bool {
    _, err := os.Stat(filePath)
    return !os.IsNotExist(err)
}

// GetBackupFilePath generates a backup file path based on the original file path.
func GetBackupFilePath(originalPath string) string {
    return filepath.Join(filepath.Dir(originalPath), "."+filepath.Base(originalPath))
}