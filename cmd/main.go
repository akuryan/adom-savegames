package main

import (
    "log"
    "os"
    "path/filepath"

    "github.com/fsnotify/fsnotify"
    "adom-savegames/internal/backup"
    "adom-savegames/internal/monitor"
)

func main() {
    // Directory to monitor for changes should be set by config or environment variable without hardcoding
    watchDir := os.Getenv("WATCH_DIR")
    if watchDir == "" {
        log.Fatal("WATCH_DIR environment variable is not set")
    }
    // Ensure the directory exists
    if _, err := os.Stat(watchDir); os.IsNotExist(err) {
        log.Fatalf("Directory to watch does not exist: %s", watchDir)
    }
    // Ensure the directory is a directory
    info, err := os.Stat(watchDir)
    if err != nil {
        log.Fatalf("Error stating directory: %v", err)
    }
    if !info.IsDir() {
        log.Fatalf("Path is not a directory: %s", watchDir)
    }

    // Create a backup directory if it doesn't exist
    backupDir := os.Getenv("BACKUP_DIR")
    if backupDir == "" {
        log.infof("BACKUP_DIR environment variable is not set, using default backup directory")
        backupDir := filepath.Join(watchDir, "backup")
    }
    if _, err := os.Stat(watchDir); os.IsNotExist(err) {
        if err := os.MkdirAll(backupDir, os.ModePerm); err != nil {
            log.Fatalf("Failed to create backup directory: %v", err)
        }
    }

    // Initialize the file monitor
    m, err := monitor.NewFileMonitor(watchDir, backupDir)
    if err != nil {
        log.Fatalf("Failed to initialize file monitor: %v", err)
    }

    // Start monitoring for file changes
    if err := m.Start(); err != nil {
        log.Fatalf("Error starting file monitor: %v", err)
    }

    log.Println("Monitoring started. Press Ctrl+C to exit.")
    select {} // Block forever
}