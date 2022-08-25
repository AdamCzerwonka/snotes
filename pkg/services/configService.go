package services

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigProvider struct {
	config map[string]interface{}
}

func LoadConfig(file string) ConfigProvider {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read config file: %s", file)
	}

	jsonString := string(data)

	config := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		log.Fatalln("Unable to parse config")
	}

	log.Println(config)

	log.Printf("Loaded config from file: %s", file)

	return ConfigProvider{config: config}
}

func (c *ConfigProvider) Get(key string) string {
	value, isPresent := c.config[key]
	if isPresent == false {
		log.Panic("Could not find config key in the config provider")
	}

	return value.(string)

}
