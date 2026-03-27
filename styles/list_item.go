package styles

import "charm.land/lipgloss/v2"

type ListItemStyle struct {
	IdStyle lipgloss.Style
	AuthorStyle lipgloss.Style
	TitleStyle lipgloss.Style
	DateStyle lipgloss.Style
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

	idStyle        := lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Italic(true)
	authorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Bold(true)
	titleStyle  := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	dateStyle   := lipgloss.NewStyle().Foreground(lipgloss.Color("238"))

	return &ListItemStyle{
		CardStyle:     cardStyle,
		SelectedStyle: selectedStyle,
		IdStyle: idStyle,
		AuthorStyle: authorStyle,
		TitleStyle: titleStyle,
		DateStyle: dateStyle,
	}
}
