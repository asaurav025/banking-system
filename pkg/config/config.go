package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Init
func Init() {
	body, err := fetchConfiguration()
	if err != nil {
		log.Error("Failed to initalize configuration")
	}
	parseConfig(body)
}

func fetchConfiguration() ([]byte, error) {
	var bodyByte []byte
	var err error
	bodyByte, err = ioutil.ReadFile("pkg/config/config.json")
	if err != nil {
		log.Panic("Failed to fetch configuration")
		return nil, err
	}

	return bodyByte, nil
}

func parseConfig(body []byte) {
	var config configBody
	err := json.Unmarshal(body, &config)
	if err != nil {
		log.Error("Could not parse the configuration")
	}
	for key, val := range config.Property {
		viper.Set(key, val)
	}
}

type configBody struct {
	Name     string                 `json:"name"`
	Property map[string]interface{} `json:"property"`
}
