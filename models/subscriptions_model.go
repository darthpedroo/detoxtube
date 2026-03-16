// This is the subscriptions model, where you see your subscribed channels.
package models

import (
	"fmt"
	tea "charm.land/bubbletea/v2"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/types"
)

type SubscriptionsModel struct {
	configManager core.ConfigManager
	title string
	subscriptions []types.Channel
	cursor int
	selected map[int]struct{}
}

func InitialSubscriptionsModel(configManager core.ConfigManager) SubscriptionsModel{
	
	config , err := configManager.ConfigLoader.LoadConfig(configManager.ConfigPath)

	if err != nil {
		return SubscriptionsModel{
			title: fmt.Sprintf("Error loading Config %v", err),
			configManager: configManager,
			subscriptions: nil,
		}
	}

	return SubscriptionsModel{
		title: "My Subscriptions",
		configManager: configManager,
		subscriptions: config.Channels,
	}
}

func (m SubscriptionsModel) Init() tea.Cmd{
	return nil
}

func (m SubscriptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			
			case "up", "k":
				if m.cursor > 0 {
                m.cursor--
            }

			case "down", "j":
				if m.cursor < len(m.subscriptions)-1 {
					m.cursor++
				}
			case "enter", "space":
				currentChannel := m.subscriptions[m.cursor]
				return InitialFeedModel(m.configManager, currentChannel.FeedUrl), nil
			
			case "left":
				return InitialMainMenuModel(m.configManager), nil
			}

	}
	return m, nil
}

func (m SubscriptionsModel) View() tea.View{
	title := m.title + "\n"

	for i, channel := range m.subscriptions{
		cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }
        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
        }
        // Render the row
        title += fmt.Sprintf("%s [%s] %s\n", cursor, checked, channel.ChannelName)
	}

	view := tea.NewView(title)
	view.AltScreen = true
	return view
}