// This is a boilerplate model to copy and paste
package models

import (
	"fmt"
	"io"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/styles"
	"github.com/darthpedroo/detoxtube/types"
)



type MainMenuModel struct {
	options []types.Button
	list list.Model
	footer FooterModel
	width int
	height int
	configManger core.ConfigManager
}

type itemMenu struct {
	button types.Button
}

func (i itemMenu) FilterValue() string {return i.button.Title}

type itemMenuDelegate struct {
	styles styles.EntryPoint
}

func (d itemMenuDelegate) Height() int {return 1}
func (d itemMenuDelegate) Spacing() int {return 0}
func (d itemMenuDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {return nil}
func (d itemMenuDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item){
	i, ok := listItem.(itemMenu)
	if !ok {
		return
	}
	str := fmt.Sprintf("%d. %s", index+1, i.button.Title)
	fn := d.styles.ListItemStyle.CardStyle.Render
	if index == m.Index(){
		fn = func(s ...string) string {
			selectedStyle := d.styles.ListItemStyle.SelectedStyle.Width(len(i.button.Title)+5)
			return selectedStyle.Render(s...)
		}
	}
	fmt.Fprint(w,fn(str))

}

func InitialMainMenuModel(configManager core.ConfigManager) MainMenuModel{
	
	options:= []types.Button{
		{
			Title: "Subscriptions",
			Redirect: InitialSubscriptionsModel(configManager),
		},
		{
			Title: "Recent Videos",
			Redirect: InitialRecentVideosModel(configManager),
		},
		{
			Title: "Load RSS Feed",
			Redirect: InitialLoadRssFeedModel(configManager),
		},
		
		
	}

	items := make([]list.Item, len(options))

	for i, buttonOption := range options {
		newItemButton := itemMenu{
			button: buttonOption,
		}
		items[i] = newItemButton
	} 

	
	delegate := itemMenuDelegate{styles: configManager.Styles }

	l := list.New(items, delegate, 100, 10) // set it as (0,0) here and in the Update we dinamically change it
	l.Title = "Welcome to Detoxtube"
	l.Styles.Title = configManager.Styles.TitleStyle.TitleStyle.Margin(0)
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)


	return MainMenuModel{
		options: options,
		list: l,
		footer: InitialFooterModel(configManager),
		configManger: configManager,

	}
}

func (m MainMenuModel) Init() tea.Cmd{
	return nil
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(msg.Width, len(m.list.Items())+5)
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String(){
			case "q", "esc":
				return m, tea.Quit
			case "enter":
				if currentItem, ok := m.list.SelectedItem().(itemMenu); ok {
					return currentItem.button.Redirect, nil
				}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m,cmd
}

func (m MainMenuModel) View() tea.View {
    content := lipgloss.JoinVertical(
        lipgloss.Left,
        m.list.View(),
        m.footer.View(),
    )
	fullView := m.configManger.Styles.Terminal.TerminalBackground.Width(m.width).Height(m.height).Render(content)
    
    view := tea.NewView(fullView)
    view.AltScreen = true
    return view
}