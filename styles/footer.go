package styles

import "charm.land/lipgloss/v2"

// inside your styles package
type FooterStyle struct {
    Background lipgloss.Style
    Key        lipgloss.Style
    Desc       lipgloss.Style
}

func NewFooterStyle() *FooterStyle {
    return &FooterStyle{
        Background: lipgloss.NewStyle().
            Background(ColorBg),
           // Height(1),
        Key: lipgloss.NewStyle().
            Foreground(ColorPrimary).
            Bold(true),
            //Padding(0, 1),
        Desc: lipgloss.NewStyle().
            Foreground(ColorSelected),
    }
}