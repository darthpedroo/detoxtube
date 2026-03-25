// This is a boilerplate model to copy and paste
package models

import (
	"charm.land/lipgloss/v2"
	core "github.com/darthpedroo/detoxtube/core"
)

type FooterModel struct {
	configManager core.ConfigManager
	Width int
}

func InitialFooterModel(configManager core.ConfigManager) FooterModel{
	return FooterModel{
		configManager: configManager,
		Width: 200,
	}
}

func (f FooterModel) View() string {
	row1 := lipgloss.JoinHorizontal(lipgloss.Top, f.configManager.Styles.Footer.Key.Render("ENTER"), f.configManager.Styles.Footer.Desc.Render("select"))
    row2 := lipgloss.JoinHorizontal(lipgloss.Top, f.configManager.Styles.Footer.Key.Render("Q"), f.configManager.Styles.Footer.Desc.Render("quit"))
    row3 := lipgloss.JoinHorizontal(lipgloss.Top, f.configManager.Styles.Footer.Key.Render("←"), f.configManager.Styles.Footer.Desc.Render("back"))
    content := lipgloss.JoinHorizontal(
        lipgloss.Left,
        row1,
        row2,
        row3,
    )


    return f.configManager.Styles.Footer.Background.
        Width(f.Width).
		MarginLeft(2).
        Render(content)
}