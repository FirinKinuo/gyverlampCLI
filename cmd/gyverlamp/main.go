package main

import (
	"github.com/FirinKinuo/gyverlampCLI/internal/config"
	"github.com/FirinKinuo/gyverlampCLI/internal/logger"
)

func main() {
	cfg := config.NewConfig()
	configPath := config.GetConfigPath()

	err := cfg.Read(configPath)
	if err != nil {
		logger.Errorf("Unable to read config from %s: %s\n", configPath, err)
		logger.Infoln("Defaults settings will be used")
	}
}
