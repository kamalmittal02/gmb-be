package config

import (
	"encoding/json"
	"fmt"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/dtos"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func InitConfig(filePath string) (*dtos.Config, error) {
	ext := filepath.Ext(filePath)
	if ext == ".yaml" || ext == ".yml" || ext == ".json" {
		configFile, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("config file read failure: %s", err.Error())
		}
		var config dtos.Config
		if ext == ".json" {
			json.Unmarshal(configFile, &config)
		} else {
			yamlErr := yaml.Unmarshal(configFile, &config)
			if yamlErr != nil {
				fmt.Print(fmt.Errorf("error in  parsing the yaml file: %s", yamlErr.Error()))
			}
		}
		return &config, nil
	}
	return nil, fmt.Errorf("invalid config file format, support yaml/json") // todo: move to error (gokit)
}
