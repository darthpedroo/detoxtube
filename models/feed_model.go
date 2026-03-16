package models

import (
	"fmt"
	"charm.land/bubbletea/v2"
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

func InitialFeedModel(configManager core.ConfigManager) FeedModel{
	
	feed , err := configManager.VideoLoader.LoadFeed("https://www.youtube.com/feeds/videos.xml?channel_id=UCHkYOD-3fZbuGhwsADBd9ZQ")

	if err != nil {
		return FeedModel{
            configManager: configManager,
			title: "Couldn't load feed",
			videos: []types.Video{},
			selected: make(map[int]struct{}),
		}
	}

	var title string
	title , err = configManager.VideoLoader.LoadTitle(feed)

	if err != nil {
		title = "default title"
	}

	initialVideos , err := configManager.VideoLoader.LoadVideos(feed, 10)
	
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
		title: title,
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
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m FeedModel) View() tea.View {
    // The header
    s := m.title + "\n"

    // Iterate over our choices
    for i, choice := range m.videos {

        // Is the cursor pointing at this choice?
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
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Title)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    view := tea.NewView(s)
	view.AltScreen = true
	return view
}