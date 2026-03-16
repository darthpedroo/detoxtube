package main

import (
	"fmt"
	"os"
	"charm.land/bubbletea/v2"
	"github.com/darthpedroo/detoxtube/core/video_loader"
	"github.com/darthpedroo/detoxtube/models"
)


func main(){

	videoLoader := core.GoFeedVideosLoader{}

	p := tea.NewProgram(models.InitialFeedModel(&videoLoader))
	if _, err := p.Run(); err != nil{
		fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
	}
}