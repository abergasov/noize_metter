package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	AppPort                 int      `yaml:"app_port"`
	SlackHookURLs           []string `yaml:"slack_hook_urls"`
	RemoteHost              string   `yaml:"remote_host"`
	RemotePass              string   `yaml:"remote_pass"`
	StorageNoiseFolder      string   `yaml:"storage_noise_folder"`
	StorageAudioFolder      string   `yaml:"storage_audio_folder"`
	StorageSubstationFolder string   `yaml:"storage_substation_folder"`
	DataHost                string   `yaml:"data_host"`
	APIKey                  string   `yaml:"api_key"`
	AppFolder               string   `yaml:"app_folder"`
	BoxName                 string   `yaml:"box_name"`
	BoxIP                   string   `yaml:"box_ip"`
	CFModbusHost            string   `yaml:"cf_modbus_host"`
}

func InitConf(confFile string) (*AppConfig, error) {
	file, err := os.Open(filepath.Clean(confFile))
	if err != nil {
		return nil, fmt.Errorf("error open config file: %w", err)
	}
	defer func() {
		if e := file.Close(); e != nil {
			log.Fatal("Error close config file", e)
		}
	}()

	var cfg AppConfig
	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("error decode config file: %w", err)
	}

	return &cfg, nil
}
