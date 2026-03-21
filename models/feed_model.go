package models

import (
	"fmt"
    "strings"
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/types"
	"github.com/darthpedroo/detoxtube/utils"
)

type FeedModel struct{
    configManager core.ConfigManager
	title string
	videos []types.Video
	cursor int
	selected map[int]struct{}
}

func InitialFeedModel(configManager core.ConfigManager, feedUrl string) FeedModel{
	
	feed , err := configManager.VideoLoader.LoadFeed(feedUrl)

	if err != nil {
		return FeedModel{
            configManager: configManager,
			title: fmt.Sprintf("Couldn't load feed %v", err),
			videos: []types.Video{},
			selected: make(map[int]struct{}),
		}
	}

	var title string
	title , err = configManager.VideoLoader.LoadTitle(feed)

	if err != nil {
		title = "default title"
	}

	initialVideos , err := configManager.VideoLoader.LoadVideos(feed, 15)
	
	if err != nil {
		return FeedModel{
            configManager: configManager,
			title: title + " warning, couldn't load videos",
			videos: []types.Video{},
			selected: make(map[int]struct{}),
		}
	}

	return FeedModel{
        configManager: configManager,
		title: "Select a video:",
		videos: initialVideos,
		selected: make(map[int]struct{}),
	}
}

func (m FeedModel) Init() tea.Cmd{
	return nil
}

func (m FeedModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    
    //if the message is a model, load that mode (example, when calling OpenApp)
    case tea.Model: 
        return msg, msg.Init()
    
    // Is it a key press?
    case tea.KeyPressMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {
        
        // These keys should exit the program.
         case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.videos)-1 {
                m.cursor++
            }

        // The "enter" key and the space bar toggle the selected state
        // for the item that the cursor is pointing at.
        case "enter", "space":
			currentVideo := m.videos[m.cursor]
			return m, tea.Batch(
				utils.OpenInNewTerminal(InitialWatchingVideoModel(m.configManager), "mpv", currentVideo.Link),
			)
        
        case "left":
            return InitialSubscriptionsModel(m.configManager), nil
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m FeedModel) View() tea.View {
    var cardStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FAFAFA")).
        Background(lipgloss.Color("#000100"))

    var selectedStyle = cardStyle.
        Background(lipgloss.Color("#454444")) // Different color for cursor

    var titleStyle = lipgloss.NewStyle().
        Blink(true).
        Bold(true)
    
    doc := strings.Builder{}
    doc.WriteString(titleStyle.Render(m.title)+"\n"+"\n")

    // Iterate over our choices
    for i, video := range m.videos {

        style := cardStyle
        style.Width(len(video.Title))
        if m.cursor == i {
            style = selectedStyle
        }
        cardContent := fmt.Sprintf("%v) %s",i+1, video.Title)
        doc.WriteString(style.Render(cardContent)+"\n")
    }

    // The footer
    doc.WriteString("\n Press q to quit.\n")
    view := tea.NewView(doc.String())
	view.AltScreen = true
	return view
}
