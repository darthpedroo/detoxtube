// This is the model for Recent Videos. It fetches all videos from all channels and shows them ordered by date
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
	"github.com/darthpedroo/detoxtube/utils"
)

// si codeo esto bien, me debería servir
// para reemplazar la parte de "Select a Video" de cuando te metes a un canal

type itemVideo struct {
	video types.Video
}

func (v itemVideo) FilterValue() string {
	return v.video.Title + v.video.Author + utils.FormatRelativeTime(v.video.PublishedDate)
}

type itemVideoDelegate struct {
	styles styles.EntryPoint
}

func (d itemVideoDelegate) Height() int                             { return 1 }
func (d itemVideoDelegate) Spacing() int                            { return 0 }
func (d itemVideoDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemVideoDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {

	i, ok := listItem.(itemVideo)

	if !ok {

		return

	}
	defaultStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("242"))

	idx := defaultStyle.Render(fmt.Sprintf("%d.", index+1))

	author := defaultStyle.Render(i.video.Author)

	title := defaultStyle.Render(i.video.Title)

	date := defaultStyle.Render(utils.FormatRelativeTime(i.video.PublishedDate))

	str := lipgloss.JoinHorizontal(lipgloss.Bottom, idx, author, " ", title, " ", date)

	block := d.styles.ListItemStyle.CardStyle.Render(str)

	if index == m.Index() {

		block = d.styles.ListItemStyle.SelectedStyle.
			Width(m.Width()).
			Render("> " + str) // Adds a cursor for the selected item

	}

	fmt.Fprint(w, block)

}

type RecentVideosModel struct {
	configManager core.ConfigManager
	title         string
	list          list.Model
	videoSort     types.VideoSort
	order         types.Order
	footer        FooterModel
	width         int
}

func InitialRecentVideosModel(configManager core.ConfigManager) RecentVideosModel {

	config, err := configManager.ConfigLoader.LoadConfig(configManager.ConfigPath)

	if err != nil {
		utils.WriteLog(err.Error())
		return RecentVideosModel{
			configManager: configManager,
			title:         "Error loading config",
			footer:        InitialFooterModel(configManager),
		}
	}

	// loop every channel and load its config
	allVideos := make([]types.Video, 0)

	for _, channel := range config.Channels {

		currentChannelFeed, err := configManager.VideoLoader.LoadFeed(channel.FeedUrl)

		if err != nil {
			utils.WriteLog(err.Error())
			return RecentVideosModel{
				configManager: configManager,
				title:         fmt.Sprintf("Error loading feed from channel %s", channel.ChannelName),
				footer:        InitialFooterModel(configManager),
			}
		}

		currentChannelVideos, err := configManager.VideoLoader.LoadVideos(currentChannelFeed, 15)

		if err != nil {
			utils.WriteLog(err.Error())
			return RecentVideosModel{
				configManager: configManager,
				title:         fmt.Sprintf("Error loading feed from channel %s", channel.ChannelName),
				footer:        InitialFooterModel(configManager),
			}
		}

		for _, video := range currentChannelVideos {
			allVideos = append(allVideos, video)
		}

	}

	items := make([]list.Item, len(allVideos))

	sortedVideos := utils.SortVideos(allVideos, types.Date, types.Descending)

	for i, video := range sortedVideos {
		newItemVideo := itemVideo{
			video: video,
		}
		items[i] = newItemVideo
		utils.WriteLog(fmt.Sprintf("Adding: %s video", video.Title))
	}

	delegate := itemVideoDelegate{styles: configManager.Styles}

	l := list.New(items, delegate, 400, 20)
	l.Title = "Recent Videos"
	l.Styles.Title = configManager.Styles.TitleStyle.TitleStyle.Margin(0)

	l.Styles.PaginationStyle = l.Styles.PaginationStyle.Padding(0)
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)

	return RecentVideosModel{
		configManager: configManager,
		title:         "Recent Videos",
		list:          l,
		footer:        InitialFooterModel(configManager),
	}
}

func (m RecentVideosModel) Init() tea.Cmd {
	return nil
}

func (m RecentVideosModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if c, ok := m.list.SelectedItem().(itemVideo); ok {
				return m, tea.Batch(
					utils.OpenInNewTerminal(InitialWatchingVideoModel(m.configManager), "mpv", c.video.Link),
				)
			}
		case "shift+left":
			return InitialMainMenuModel(m.configManager), nil
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m RecentVideosModel) View() tea.View {
	view := tea.NewView(m.list.View() + "\n" + m.footer.View())
	view.AltScreen = true
	return view
}
