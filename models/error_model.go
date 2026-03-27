// This is a boilerplate model to copy and paste
package models

import (
	"charm.land/lipgloss/v2"
	"fmt"
	"strings"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/types"
)

type ErrorModel struct {
    configManager core.ConfigManager
    Width         int
    Error         types.ErrorFetchVideo
    ScrollIndex   int 
}

func InitialErrorModel(configManager core.ConfigManager, currentError types.ErrorFetchVideo) ErrorModel {
	return ErrorModel{
		configManager: configManager,
		Width:         200,
		Error: currentError,
	}
}

func (f ErrorModel) View() string {
    if len(f.Error.UnavailableChannels) == 0 {
        return ""
    }

    // 1. Define the Minimalist Red Border Style
    errorBoxStyle := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).      // Rounded is more modern/minimal
        BorderForeground(lipgloss.Color("196")). // A clean Red
        Padding(0, 1).                         // Horizontal padding
        MarginLeft(2).
        Width(f.Width - 4)                     // Account for borders/padding

    // 2. Style the text inside
    messageStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
    channelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)

    var content string
    for _, channel := range f.Error.UnavailableChannels {
        // Build the line: "● ChannelName is offline"
        line := fmt.Sprintf(" %s %s %s\n", 
            lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("●"),
            channelStyle.Render(channel),
            messageStyle.Render("offline"),
        )
        content += line
    }

    // 3. Render the content inside the bordered box
    return errorBoxStyle.Render(strings.TrimSpace(content))
}
