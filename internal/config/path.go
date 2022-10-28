package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getHomePathFromEnvOrPoint(key string) string {
	homeDir := os.Getenv(key)
	if len(homeDir) == 0 {
		return "."
	}

	return homeDir
}

func getConfigDirPath() string {
	pathToConfigFolder := "./.gyverlampCLI"

	switch runtime.GOOS {
	case "windows":
		homeDrive := getHomePathFromEnvOrPoint("HOMEDRIVE")
		homePath := getHomePathFromEnvOrPoint("HOMEPATH")
		pathToConfigFolder = fmt.Sprintf("%s%s/.gyverlampCLI", homeDrive, homePath)
	case "linux", "darwin":
		homeDir := getHomePathFromEnvOrPoint("HOME")
		pathToConfigFolder = fmt.Sprintf("%s/.config/gyverlampCLI", homeDir)
	}

	return pathToConfigFolder
}

// GetConfigPath returns the path to the configuration file, depending on the OS
//
// For Windows path is %HOMEDRIVE%%HOMEPATH%\\.gyverlampCLI\\config.yml
// For Linux and macOS path is $HOME/.config/gyverlampCLI/config.yml
// If OS is undefined, path will be located at the tool launch location ./.gyverlampCLI/config.yml
func GetConfigPath() string {
	configDirPath := getConfigDirPath()
	return filepath.FromSlash(fmt.Sprintf("%s/config.yml", configDirPath))
}
