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


	idxStr := fmt.Sprintf("%d. ", index+1)
    authorStr := i.video.Author
    titleStr := i.video.Title
    dateStr := utils.FormatRelativeTime(i.video.PublishedDate)
	
	var str string
    itemStyle := d.styles.ListItemStyle.CardStyle
	
	if index == m.Index() {
        //  If SELECTED: Join raw strings first, THEN style the whole block
        content := fmt.Sprintf("> %s%s %s %s", idxStr, authorStr, titleStr, dateStr)
        str = d.styles.ListItemStyle.SelectedStyle.Width(m.Width()).Render(content)
    } else {
        //  If NOT SELECTED: Style components individually		
		
		sIdx := d.styles.ListItemStyle.IdStyle.Render(idxStr)
		sAuthor := d.styles.ListItemStyle.AuthorStyle.Render(authorStr)
        sTitle := d.styles.ListItemStyle.TitleStyle.Render(titleStr)
        sDate := d.styles.ListItemStyle.DateStyle.Render(dateStr)
        
        content := lipgloss.JoinHorizontal(lipgloss.Bottom, "  ", sIdx, sAuthor, " ", sTitle, " ", sDate)
        str = itemStyle.Render(content)
    }
	fmt.Fprint(w, str)

}

type RecentVideosModel struct {
	configManager core.ConfigManager
	title         string
	list          list.Model
	videoSort     types.VideoSort
	order         types.Order
	footer        FooterModel
	width         int
	height int
	showError bool
	errorFooter ErrorModel
}

func InitialRecentVideosModel(configManager core.ConfigManager) RecentVideosModel {
	var showError bool
	showError = false

	errorFetchVideo := types.ErrorFetchVideo{

	}

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
			utils.WriteLog(fmt.Sprintf("Error loading feed from channel %s . Error: %s",channel.ChannelName, err.Error()))
			showError = true
			errorFetchVideo.UnavailableChannels = append(errorFetchVideo.UnavailableChannels, channel.ChannelName)
			continue
			//return RecentVideosModel{
			//	configManager: configManager,
			//	title:         fmt.Sprintf("Error loading feed from channel %s", channel.ChannelName),
			//	footer:        InitialFooterModel(configManager),
			//}/
		}

		currentChannelVideos, err := configManager.VideoLoader.LoadVideos(currentChannelFeed, 15)

		if err != nil {
			utils.WriteLog(fmt.Sprintf("Error loading video from channel %s . Error: %s",currentChannelFeed, err.Error()))
			return RecentVideosModel{
				configManager: configManager,
				title:         fmt.Sprintf("Error loading feed from channel %s", channel.ChannelName),
				footer:        InitialFooterModel(configManager),
				showError: showError,
				errorFooter: InitialErrorModel(configManager, errorFetchVideo),
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
		showError: showError,
		errorFooter: InitialErrorModel(configManager, errorFetchVideo),
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
		case "ctrl+x":
			m.showError = !m.showError
            return m, nil
        
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
    // 1. If Error is active, ignore the list and render ONLY the error in the center
    if m.showError {
        errorView := m.errorFooter.View()
        
        // This centers the errorBox both horizontally and vertically
        centeredView := lipgloss.Place(
            m.width / 3, 
            m.height / 3, 
            lipgloss.Center, 
            lipgloss.Center, 
            errorView,
        )
        
        view := tea.NewView(centeredView)
        view.AltScreen = true
        return view
    }

    // 2. Normal View
    content := lipgloss.JoinVertical(lipgloss.Left, m.list.View(), m.footer.View())
    view := tea.NewView(content)
    view.AltScreen = true
    return view
}