package core

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darthpedroo/detoxtube/types"
)

type JsonConfigLoader struct {
}

func (c *JsonConfigLoader) LoadConfig(configPath string) (*types.Config, error) {

	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Error when loading config from json with path:  %v. \n Error: %v ", configPath, err)
	}

	var payload types.Config

	err = json.Unmarshal(content, &payload)

	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling json file")
	}

	return &payload, nil
}

func (c *JsonConfigLoader) AddChannel(configPath string, channel types.Channel) error {

	data, err := c.LoadConfig(configPath)

	if err != nil {
		return err
	}

	data.Channels = append(data.Channels, channel)

	dataBytes, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("Error Adding Channel to json file")
	}

	err = os.WriteFile(configPath, dataBytes, 0644)

	if err != nil {
		return fmt.Errorf("Error writing to file %v", configPath)
	}
	return nil
}
