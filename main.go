package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	core "github.com/darthpedroo/detoxtube/core"
	config_loader "github.com/darthpedroo/detoxtube/core/config_loader"
	videoLoader "github.com/darthpedroo/detoxtube/core/video_loader"
	"github.com/darthpedroo/detoxtube/styles"

	//"github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/models"
)


func main(){

	configManager := core.ConfigManager{
		VideoLoader: &videoLoader.GoFeedVideosLoader{},
		ConfigLoader: &config_loader.JsonConfigLoader{},
		ConfigPath: "config.json",
		Styles: styles.NewEntryPoint(),
	}

	p := tea.NewProgram(models.InitialMainMenuModel(configManager))
	if _, err := p.Run(); err != nil{
		fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
	}
}