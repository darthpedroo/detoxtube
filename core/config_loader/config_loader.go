package core

import "github.com/darthpedroo/detoxtube/types"



type ConfigLoader interface{
	LoadConfig(configPath string) (*types.Config, error)
}