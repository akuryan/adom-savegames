package monitor

import (
    "context"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/fsnotify/fsnotify"
    "adom-savegames/internal/backup"
    "adom-savegames/internal/restore"
)

type Monitor struct {
    watchDir string
    watcher  *fsnotify.Watcher
    ctx      context.Context
    cancel   context.CancelFunc
}

func NewMonitor(watchDir string) (*Monitor, error) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, fmt.Errorf("failed to create watcher: %w", err)
    }

    ctx, cancel := context.WithCancel(context.Background())

    return &Monitor{
        watchDir: watchDir,
        watcher:  watcher,
        ctx:      ctx,
        cancel:   cancel,
    }, nil
}

func (m *Monitor) Start() error {
    err := m.watcher.Add(m.watchDir)
    if err != nil {
        return fmt.Errorf("failed to add watch directory: %w", err)
    }

    go m.watch()

    return nil
}

func (m *Monitor) watch() {
    defer m.watcher.Close()

    for {
        select {
        case <-m.ctx.Done():
            return
        case event, ok := <-m.watcher.Events:
            if !ok {
                return
            }
            m.handleEvent(event)
        case err, ok := <-m.watcher.Errors:
            if !ok {
                return
            }
            log.Printf("error: %v", err)
        }
    }
}

func (m *Monitor) handleEvent(event fsnotify.Event) {
    if event.Op&fsnotify.Create == fsnotify.Create {
        log.Printf("New file created: %s", event.Name)
        if err := backup.BackupFile(event.Name); err != nil {
            log.Printf("Failed to back up file: %v", err)
        }
    } else if event.Op&fsnotify.Remove == fsnotify.Remove {
        log.Printf("File deleted: %s", event.Name)
        time.Sleep(1 * time.Second)
        if err := restore.RestoreFile(event.Name); err != nil {
            log.Printf("Failed to restore file: %v", err)
        }
    }
}

func (m *Monitor) Stop() {
    m.cancel()
}