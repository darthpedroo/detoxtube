// This is a boilerplate model to copy and paste
package models

import tea "charm.land/bubbletea/v2"

type BoilerplateModel struct {

}

func InitialBoilerplateModel() BoilerplateModel{
	return BoilerplateModel{}
}

func (m BoilerplateModel) Init() tea.Cmd{
	return nil
}

func (m BoilerplateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			}
	}
	return m, nil
}

func (m BoilerplateModel) View() tea.View{
	title := "Boilerplate Model"
	return tea.NewView(title)
}