package styles

import "charm.land/lipgloss/v2"

type ListItemStyle struct {
	CardStyle     lipgloss.Style
	SelectedStyle lipgloss.Style
	Width         int
}

func NewListItemStyle() *ListItemStyle {

	cardStyle := lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Background(ColorBg).
		MarginLeft(2)

	selectedStyle := cardStyle.
		Background(ColorSelected)

	return &ListItemStyle{
		CardStyle:     cardStyle,
		SelectedStyle: selectedStyle,
	}
}
