package backup

import (
	"io"
	"os"
	"path/filepath"
)

// BackupFile copies a new file to the backup location.
func BackupFile(src string, dest string) error {
	input, err := os.Open(src)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	if err != nil {
		return err
	}

	return output.Sync()
}

// GetBackupPath generates the backup file path based on the source file path.
func GetBackupPath(src string, backupDir string) string {
	return filepath.Join(backupDir, filepath.Base(src))
}