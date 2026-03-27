// This is a boilerplate model to copy and paste
package models

import (
	tea "charm.land/bubbletea/v2"
	core "github.com/darthpedroo/detoxtube/core"
)

type BoilerplateModel struct {
	configManager core.ConfigManager
}

func InitialBoilerplateModel(configManager core.ConfigManager) BoilerplateModel {
	return BoilerplateModel{
		configManager: configManager,
	}
}

func (m BoilerplateModel) Init() tea.Cmd {
	return nil
}

func (m BoilerplateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m BoilerplateModel) View() tea.View {
	title := "Boilerplate Model"
	view := tea.NewView(title)
	view.AltScreen = true
	return view
}
