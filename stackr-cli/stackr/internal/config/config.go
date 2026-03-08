package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const (
	Version          = "1.2.0"
	ReleaseBaseURL   = "https://github.com/stackr-lat/cli/releases/download"
	ReleaseLatestURL = "https://github.com/stackr-lat/cli/releases/latest"
)

type Config struct {
	Token string `json:"token"`
}

type StackrConfig struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Lang    string `json:"language"`
	Memory  string `json:"memory"`
	Command string `json:"command"`
}

func configPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".stackr", "config.json")
}

func Load() *Config {
	cfg := &Config{}
	data, err := os.ReadFile(configPath())
	if err != nil {
		return cfg
	}
	_ = json.Unmarshal(data, cfg)
	return cfg
}

func Save(cfg *Config) error {
	p := configPath()
	if err := os.MkdirAll(filepath.Dir(p), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0600)
}

func GetToken() string {
	if t := os.Getenv("STACKR_API_TOKEN"); t != "" {
		return t
	}
	return Load().Token
}

func FindLocalConfig() (*StackrConfig, string) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, ""
	}
	for {
		candidate := filepath.Join(dir, "stackr.config")
		data, err := os.ReadFile(candidate)
		if err == nil {
			cfg := &StackrConfig{}
			if json.Unmarshal(data, cfg) == nil && cfg.ID != "" {
				return cfg, candidate
			}
			for _, line := range strings.Split(string(data), "\n") {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}
				parts := strings.SplitN(line, "=", 2)
				if len(parts) != 2 {
					continue
				}
				key, val := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
				switch key {
				case "id":
					cfg.ID = val
				case "name":
					cfg.Name = val
				case "language":
					cfg.Lang = val
				case "memory":
					cfg.Memory = val
				case "command":
					cfg.Command = val
				}
			}
			if cfg.ID != "" {
				return cfg, candidate
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return nil, ""
}

func FindLocalAppID() (string, string) {
	cfg, path := FindLocalConfig()
	if cfg != nil {
		return cfg.ID, path
	}
	return "", ""
}
