package core

import (
	config_loader "github.com/darthpedroo/detoxtube/core/config_loader"
	video_loader "github.com/darthpedroo/detoxtube/core/video_loader"
	"github.com/darthpedroo/detoxtube/styles"
)

type ConfigManager struct {
	VideoLoader  video_loader.VideosLoader
	ConfigLoader config_loader.ConfigLoader
	ConfigPath   string
	Styles       styles.EntryPoint
}
