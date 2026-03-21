package styles

import "charm.land/lipgloss/v2"

type TitleStyle struct {
	TitleStyle lipgloss.Style
}

func NewTitleStyle() *TitleStyle{

	titleStyle := lipgloss.NewStyle().
        Blink(true).
        Bold(true).
		Margin(0).
		Padding(0)

	return &TitleStyle{
		TitleStyle: titleStyle,
	}
} 