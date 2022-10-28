package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const (
	perm0664 os.FileMode = 0664 // Permission 0664: Owner and Group - Read, Write, Other - Only read
)

// Config contains the entire instrument configuration
type Config struct {
	GyverLamp gyverLampYAML `yaml:"gyverLamp"`
}

// NewConfig returns new default Config
func NewConfig() *Config {
	gyverLampYaml := newGyverLampYaml()

	config := &Config{
		GyverLamp: gyverLampYaml,
	}

	return config
}

func (c *Config) readYamlFile(path string) (bytesRead []byte, err error) {
	bytesRead, err = os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read yaml config file, %w", err)
	}

	return bytesRead, nil
}

// Read reads configuration from .yml file
func (c *Config) Read(path string) error {
	yamlBytesRead, err := c.readYamlFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlBytesRead, c)
	if err != nil {
		return fmt.Errorf("unable to parse yaml config file, %w", err)
	}

	return nil
}

func (c *Config) makeConfigFolder(path string) error {
	err := os.MkdirAll(path, perm0664)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) createYamlFile(path string, data []byte) error {
	err := os.WriteFile(path, data, perm0664)
	if err != nil {
		return err
	}

	return nil
}

// Create creates config file at path with current configuration
func (c *Config) Create(path string) error {
	err := c.makeConfigFolder(filepath.Dir(path))
	if err != nil {
		return err
	}

	marshaledYaml, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	err = c.createYamlFile(path, marshaledYaml)
	if err != nil {
		return err
	}

	return nil
}
