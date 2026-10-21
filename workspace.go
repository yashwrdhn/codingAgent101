package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Workspace struct {
	Root string
}

func NewWorkspace(rootPath string) (*Workspace, error) {
	if rootPath == "" {
		rootPath = "."
	}

	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	info, err := os.Stat(absPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("workspace directory does not exist: %s", absPath)
		}
		return nil, fmt.Errorf("failed to stat workspace path: %w", err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("workspace path is not a directory: %s", absPath)
	}

	return &Workspace{Root: absPath}, nil

}
