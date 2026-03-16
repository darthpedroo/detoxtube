// This is the model that shows when you are watching a video
// This is a boilerplate model to copy and paste
package models

import tea "charm.land/bubbletea/v2"

type WatchingVideoModel struct {

}

func InitialWatchingVideoModel() BoilerplateModel{
	return BoilerplateModel{}
}

func (m WatchingVideoModel) Init() tea.Cmd{
	return nil
}

func (m WatchingVideoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "q":
				return BoilerplateModel{}, nil
			}
	}
	return m, nil
}

func (m WatchingVideoModel) View() tea.View{
	title := "You are currently watching a Video! Press 'q' to go back"
	return tea.NewView(title)
}