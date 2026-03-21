// This is the subscriptions model, where you see your subscribed channels.
package models

import (
	"fmt"
	"io"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/styles"
	"github.com/darthpedroo/detoxtube/types"
	"github.com/darthpedroo/detoxtube/utils"
)

type itemChannel struct{
	channel types.Channel
}

func (c itemChannel) FilterValue() string { return c.channel.ChannelName}

type itemDelegate struct {
	styles *styles.EntryPoint
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(itemChannel)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.channel.ChannelName)

	fn := d.styles.ListItemStyle.CardStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			selectedStyle := d.styles.ListItemStyle.SelectedStyle.Width(len(i.channel.ChannelName)+5)
			return selectedStyle.Render(s...)
		}
	}

	fmt.Fprint(w, fn(str))
}


type SubscriptionsModel struct {
	configManager core.ConfigManager
	title string
	list list.Model
	listStyle list.Styles
	videoSort types.VideoSort
	order 	types.Order
}

func InitialSubscriptionsModel(configManager core.ConfigManager) SubscriptionsModel{
	
	config , err := configManager.ConfigLoader.LoadConfig(configManager.ConfigPath)

	if err != nil {
		return SubscriptionsModel{
			title: fmt.Sprintf("Error loading Config %v", err),
			configManager: configManager,
		}
	}

	items := make([]list.Item, len(config.Channels))

	sortedSubscriptions := utils.SortSubscriptions(config.Channels, types.Alphabetically, types.Ascendant)
	
	for i, channel := range sortedSubscriptions {
		newItemChannel := itemChannel{
			channel: channel,
		}
		items[i] = newItemChannel
	}

	styles := styles.NewEntryPoint()
	delegate := itemDelegate{styles: &styles}

	l := list.New(items, delegate,300,500)
	l.Title = "My Subscriptions"
	l.Styles.Title = styles.TitleStyle.TitleStyle
	l.Styles.PaginationStyle = l.Styles.PaginationStyle.Padding(0)
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)

	return SubscriptionsModel{
		title: "My Subscriptions",
		configManager: configManager,
		list: l,
	}
}

func (m SubscriptionsModel) Init() tea.Cmd{
	return nil
}

func (m SubscriptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.list.SetSize(msg.Width, msg.Height)
        return m, nil

    case tea.KeyPressMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "enter":
            // Get the selected item from the list model
            if c, ok := m.list.SelectedItem().(itemChannel); ok {
                return InitialFeedModel(m.configManager, c.channel.FeedUrl), nil
            }
        case "left":
            return InitialMainMenuModel(m.configManager), nil
        }
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg) 
    return m, cmd
}

func (m SubscriptionsModel) View() tea.View {
    // No more manual loops!
    view := tea.NewView("\n" + m.list.View())
	view.AltScreen = true
	return view
}
