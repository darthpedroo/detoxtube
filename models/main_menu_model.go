// This is a boilerplate model to copy and paste
package models

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/types"
)



type MainMenuModel struct {
	title string
	options []types.Button
	cursor int
	selected map[int]struct{}
}

func InitialMainMenuModel(configManager core.ConfigManager) MainMenuModel{
	
	options:= []types.Button{
		{
			Title: "Subscriptions",
			Redirect: InitialSubscriptionsModel(configManager),
		},
		{
			Title: "Load RSS Feed",
			Redirect: InitialBoilerplateModel(configManager),
		},
	}

	return MainMenuModel{
		title: "Welcome to DetoxTube",
		options: options,
	}
}

func (m MainMenuModel) Init() tea.Cmd{
	return nil
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {

	case tea.Model:
		return msg, msg.Init()

		case tea.KeyPressMsg:
			switch msg.String() {
			
			case "ctrl+c", "q":
				return m, tea.Quit
			
			case "up", "k":
				if m.cursor > 0 {
                m.cursor--
            }

			case "down", "j":
            	if m.cursor < len(m.options)-1 {
                	m.cursor++
            }

			case "enter", "space":
			currentVideo := m.options[m.cursor]

			return currentVideo.Redirect, nil

			}

		
	}
	return m, nil
}

func (m MainMenuModel) View() tea.View{
	title := m.title + "\n"

	for i, choice := range m.options{
		cursor := " "
		
		if m.cursor == i{
			cursor = ">"
		}

		checked := " "

		if _, ok := m.selected[i]; ok{
			checked = "x"
		}
		title += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Title)
	}
	
	view := tea.NewView(title)
	view.AltScreen = true
	return view
}