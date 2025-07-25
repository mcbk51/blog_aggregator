package config

import (
  "encoding/json"
	"os"
	"path/filepath"
)


const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}	

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath() 
	if err != nil {
	  return Config{}, err 
	}
	
	data, err := os.Open(configPath)
	if err != nil { 
			return Config{}, err 
	}
	defer data.Close()
  
	decoder := json.NewDecoder(data)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err 
	}

	return cfg, nil
}


func getConfigFilePath() (string, error){
  homeDir, err := os.UserHomeDir()
  if err != nil {
   return "", err 
  }
  return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error{
	configPath, err := getConfigFilePath() 
	if err != nil {
	  return err 
	}

	data, err := os.Create(configPath)
  if err != nil {
		return err 
  }
	defer data.Close()

  encoder :=json.NewEncoder(data)	
	err = encoder.Encode(cfg)
  if err != nil {
		return err 
  }

	return nil
}
