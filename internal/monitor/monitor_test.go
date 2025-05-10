package monitor

import (
	"os"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
)

func TestMonitorFileChanges(t *testing.T) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		t.Fatalf("Failed to create watcher: %v", err)
	}
	defer watcher.Close()

	testDir := "./testdir"
	err = os.MkdirAll(testDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	err = watcher.Add(testDir)
	if err != nil {
		t.Fatalf("Failed to add directory to watcher: %v", err)
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		_, err := os.Create(testDir + "/newfile.txt")
		if err != nil {
			t.Errorf("Failed to create new file: %v", err)
		}
	}()

	select {
	case event, ok := <-watcher.Events:
		if !ok {
			t.Fatal("Watcher events channel closed")
		}
		assert.Equal(t, event.Op&fsnotify.Create, fsnotify.Create, "Expected create event")
	case err := <-watcher.Errors:
		t.Errorf("Watcher error: %v", err)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for file create event")
	}
}