package core

import "github.com/darthpedroo/detoxtube/types"



type ConfigLoader interface{
	LoadConfig(configPath string) (*types.Config, error)
	AddChannel(configPath string, channel types.Channel) error
}