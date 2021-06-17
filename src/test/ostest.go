package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	// get home dir
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	// executable dir
	if exePath, err := os.Executable(); err == nil {
		fmt.Println(exePath)
		fmt.Println(filepath.Dir(exePath))
	}

	// temp dir
	fmt.Println(os.TempDir())

	// use third party library to find home dir
	path, err := homedir.Dir()
	if (err != nil) {
		log.Fatal(err)
	}
	fmt.Println(path)

	// app data dir
	if runtime.GOOS == "windows" {
		if appDataDir := os.Getenv("APPDATA"); appDataDir != "" {
			fmt.Println(appDataDir)
		}
	}

	// app data dir, based on rclone
	var dir string
	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("LocalAppData")
	case "darwin":
		dir = os.Getenv("HOME")
		if dir != "" {
			dir += "/Library/Caches"
		}
	case "plan9":
		dir = os.Getenv("home")
		if dir != "" {
			dir += "/lib/cache"
		}
	default: // Unix
		dir = os.Getenv("XDG_CACHE_HOME")
		if dir == "" {
			dir = os.Getenv("HOME")
			if dir != "" {
				dir += "/.cache"
			}
		}
	}
	if dir == "" {
		dir = os.TempDir()
	}
	fmt.Println(dir)
}

func findFile(dir string, name string) string {
	path := filepath.Join(dir, name)
	if _, err := os.Stat(path); err != nil {
		return ""
	}
	return path
}