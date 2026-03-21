package styles

import "charm.land/lipgloss/v2"

type TerminalStyle struct {

	TerminalBackground lipgloss.Style
}

func NewTerminalStyle() *TerminalStyle{

	Background := lipgloss.NewStyle().
        Background(lipgloss.Color("#000000")). // Your Black
        Foreground(lipgloss.Color("#FAFAFA"))
	
	return &TerminalStyle{
		TerminalBackground: Background,
	}
} 