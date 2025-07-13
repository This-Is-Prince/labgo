package encoderdecoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

const FILE_PATH = "./configs.json"

type Config struct {
	Name          string
	Version       string
	EnableLogging bool
}

func defaultConfigs() []Config {
	return []Config{
		{
			Name:          "DefaultEncoder",
			Version:       "1.0",
			EnableLogging: true,
		},
		{
			Name:          "DefaultDecoder",
			Version:       "1.1",
			EnableLogging: false,
		},
	}
}

func populateConfigs() error {
	file, err := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_APPEND, 0644)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(FILE_PATH)
	}

	if err != nil {
		return err
	}

	defer file.Close()

	configs := defaultConfigs()
	totalErrors := 0

	encoder := json.NewEncoder(file)

	for _, config := range configs {
		if _err := encoder.Encode(config); _err != nil {
			totalErrors++
			err = _err
		}
	}

	if totalErrors == len(configs) {
		return err
	}

	return nil
}

func getConfigs() ([]Config, error) {
	file, err := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var configs []Config
	var tempConfig *Config
	totalErrors := 0

	decoder := json.NewDecoder(file)

	for {
		if _err := decoder.Decode(&tempConfig); _err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			totalErrors++
			err = _err
		} else {
			configs = append(configs, *tempConfig)
			tempConfig = nil
		}
	}

	if totalErrors == len(configs) {
		return nil, err
	}

	return configs, nil
}

func Learn(run bool) {
	if !run {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:-", r)
		}
	}()

	if err := populateConfigs(); err != nil {
		panic(err)
	}

	fmt.Println("Configs populated successfully.")

	configs, err := getConfigs()
	if err != nil {
		panic(err)
	}

	fmt.Println("Configs retrieved successfully:")
	for _, config := range configs {
		fmt.Println(config)
	}
}
