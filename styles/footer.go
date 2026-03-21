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
            Background(ColorBg).
            PaddingRight(2).
            Foreground(ColorPrimary).
            Bold(true),
            //Padding(0, 1),
        Desc: lipgloss.NewStyle().
            Background(ColorBg).
            PaddingRight(2).
            Foreground(ColorSelected),
    }
}