# ADOM Savegames Utility

## Overview
The ADOM Savegames Utility is a Go-based application designed to monitor a specific file system location for changes. It provides functionality to back up newly created files and restore deleted files from backup after a short delay.

## Features
- **Backup New Files**: Automatically backs up any new files created in the monitored directory.
- **Restore Deleted Files**: Restores any deleted files from the backup after a 1-second delay.

## Project Structure
```
adom-savegames
├── cmd
│   └── main.go               # Entry point of the application
├── internal
│   ├── backup
│   │   ├── backup.go         # Backup functionality implementation
│   │   └── backup_test.go    # Unit tests for backup functionality
│   ├── monitor
│   │   ├── monitor.go        # File monitoring logic
│   │   └── monitor_test.go   # Unit tests for monitoring functionality
│   └── restore
│       ├── restore.go        # Restoration logic for deleted files
│       └── restore_test.go   # Unit tests for restoration functionality
├── pkg
│   └── utils
│       ├── fileutils.go      # Utility functions for file operations
│       └── fileutils_test.go # Unit tests for utility functions
├── go.mod                     # Go module configuration
├── go.sum                     # Checksums for module dependencies
└── README.md                  # Project documentation
```

## Getting Started

### Prerequisites
- Go 1.16 or later installed on your machine.

### Installation
1. Clone the repository:
   ```
   git clone <repository-url>
   cd adom-savegames
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```

### Usage
To run the application, execute the following command:
```
go run cmd/main.go
```

### Running Tests
To run the tests for the project, use:
```
go test ./...
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.