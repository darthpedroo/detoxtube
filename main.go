package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	core "github.com/darthpedroo/detoxtube/core"
	config_loader "github.com/darthpedroo/detoxtube/core/config_loader"
	videoLoader "github.com/darthpedroo/detoxtube/core/video_loader"
	"github.com/darthpedroo/detoxtube/styles"
	"github.com/darthpedroo/detoxtube/utils"

	//"github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/models"
)

func main() {

	homePath, err := utils.GetHome()

	if err != nil {
		utils.WriteLog(err.Error())
		os.Exit(1)
	}

	err = utils.CreateConfigDir()

	if err != nil {
		utils.WriteLog(err.Error())

	}

	configPath := homePath + "/.config/detoxtube/config.json"
	configPath = "config.json"
	utils.WriteLog(fmt.Sprintf("este es el path %v ", configPath))

	configManager := core.ConfigManager{
		VideoLoader:  &videoLoader.GoFeedVideosLoader{},
		ConfigLoader: &config_loader.JsonConfigLoader{},
		ConfigPath:   configPath,
		Styles:       styles.NewEntryPoint(),
	}

	p := tea.NewProgram(models.InitialMainMenuModel(configManager))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
