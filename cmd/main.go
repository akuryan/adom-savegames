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
    // Define the directory to monitor
    watchDir := "./path/to/monitor" // Update this path as needed

    // Create a backup directory if it doesn't exist
    backupDir := filepath.Join(watchDir, "backup")
    if err := os.MkdirAll(backupDir, os.ModePerm); err != nil {
        log.Fatalf("Failed to create backup directory: %v", err)
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