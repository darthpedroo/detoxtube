// This is a boilerplate model to copy and paste
package models

import (
	"charm.land/lipgloss/v2"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/styles"
)

type FooterModel struct {
	configManager core.ConfigManager
	styles *styles.EntryPoint
	Width int
}

func InitialFooterModel(configManager core.ConfigManager) FooterModel{
	c := styles.NewEntryPoint()
	return FooterModel{
		configManager: configManager,
		styles: &c,
		Width: 200,
	}
}

func (f FooterModel) View() string {

    
	row1 := lipgloss.JoinHorizontal(lipgloss.Top, f.styles.Footer.Key.Render("ENTER"), f.styles.Footer.Desc.Render("select"))
    row2 := lipgloss.JoinHorizontal(lipgloss.Top, f.styles.Footer.Key.Render("Q"), f.styles.Footer.Desc.Render("quit"))
    row3 := lipgloss.JoinHorizontal(lipgloss.Top, f.styles.Footer.Key.Render("←"), f.styles.Footer.Desc.Render("back"))
    content := lipgloss.JoinHorizontal(
        lipgloss.Left,
        row1,
        row2,
        row3,
    )


    return f.styles.Footer.Background.
        Width(f.Width).
		MarginLeft(2).
        Render(content)
}