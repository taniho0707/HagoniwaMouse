package repository

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	configDir = "HagoniwaMouse"
)

type Filesystem struct {
}

func NewFilesystem() (*Filesystem, error) {
	fs := &Filesystem{}
	return fs, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExist(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func GetConfigDir() (string, error) {
	dir, err := userConfigDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, configDir)
	return path, makeDir(path)
}

func makeDir(dir string) error {
	dirPath := filepath.FromSlash(dir)
	fnMakeDir := func() error { return os.MkdirAll(dirPath, os.ModePerm) }
	info, err := os.Stat(dirPath)
	switch {
	case err == nil:
		if info.IsDir() {
			return nil // diirectory already exists
		} else {
			return fmt.Errorf("path exists but is not a directory: %s", dir)
		}
	case os.IsNotExist(err):
		return fnMakeDir()
	default:
		return err
	}
}

func userConfigDir() (string, error) {
	var dir string

	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("AppData")
		if dir == "" {
			return "", errors.New("%AppData% is not set")
		}
	case "darwin":
	case "plan9":
		dir = os.Getenv("home")
		if dir == "" {
			return "", errors.New("$home is not set")
		}
		dir = filepath.Join(dir, "lib")
	default: // Unix
		dir = os.Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			dir = os.Getenv("HOME")
			if dir == "" {
				return "", errors.New("neither $XDG_CONFIG_HOME nor $HOME are set")
			}
			dir = filepath.Join(dir, ".config")
		}
	}

	return dir, nil
}
