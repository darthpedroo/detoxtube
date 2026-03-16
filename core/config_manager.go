package core

import (
	video_loader "github.com/darthpedroo/detoxtube/core/video_loader"
	config_loader "github.com/darthpedroo/detoxtube/core/config_loader"
	)

type ConfigManager struct {
	VideoLoader video_loader.VideosLoader
	ConfigLoader config_loader.ConfigLoader
	ConfigPath string
}