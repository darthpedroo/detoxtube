package styles

import "charm.land/lipgloss/v2"

type TitleStyle struct {
	TitleStyle lipgloss.Style
}

func NewTitleStyle() *TitleStyle {

	titleStyle := lipgloss.NewStyle().
		Blink(true).
		Bold(true).
		MarginLeft(2).
		Padding(0)

	return &TitleStyle{
		TitleStyle: titleStyle,
	}
}
