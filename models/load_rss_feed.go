// This is a boilerplate model to copy and paste
package models

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	core "github.com/darthpedroo/detoxtube/core"
	"github.com/darthpedroo/detoxtube/types"
	"github.com/darthpedroo/detoxtube/utils"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type LoadRssFeedModel struct {
	configManager core.ConfigManager
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	err error
	quitting   bool

}

func InitialLoadRssFeedModel(configManager core.ConfigManager) LoadRssFeedModel{
	m := LoadRssFeedModel{
		inputs: make([]textinput.Model, 2),
		configManager: configManager,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CharLimit = 32
		t.SetWidth(30)
		s := t.Styles()
		s.Cursor.Color = lipgloss.Color("205")
		s.Focused.Prompt = focusedStyle
		s.Focused.Text = focusedStyle
		s.Blurred.Prompt = blurredStyle
		s.Focused.Text = focusedStyle
		t.SetStyles(s)

		switch i {
		case 0:
			t.Placeholder = "Enter Channel Name"
			t.Focus()
		case 1:
			t.Placeholder = "Enter ChannelID"
			t.CharLimit = 64
		}

		m.inputs[i] = t
	}

	return m
}

func (m LoadRssFeedModel) Init() tea.Cmd{
	return textinput.Blink
}

func (m LoadRssFeedModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				// Save info!!!

				channelName := m.inputs[0]
				channelId := m.inputs[1]

				fixedChannelId := utils.CreateRssFeedFromChannelId(channelId.Value())

				newChannel := types.Channel{
					ChannelName: channelName.Value(),
					FeedUrl: fixedChannelId,
				}

				err := m.configManager.ConfigLoader.AddChannel(m.configManager.ConfigPath, newChannel)

				if err != nil {
					m.err = err
					return m, nil
				}

				return InitialMainMenuModel(m.configManager), tea.ClearScreen
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *LoadRssFeedModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m LoadRssFeedModel) View() tea.View {
	var b strings.Builder
	var c *tea.Cursor

	for i, in := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
		if m.cursorMode != cursor.CursorHide && in.Focused() {
			c = in.Cursor()
			if c != nil {
				c.Y += i
			}
		}
	}

	if m.err != nil {
        errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
        fmt.Fprintf(&b, "\n\n%s", errorStyle.Render("Error: "+m.err.Error()))
    }

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	if m.quitting {
		b.WriteRune('\n')
	}

	view := tea.NewView(b.String())
	view.Cursor = c
	view.AltScreen = true
	return view
}



func (m LoadRssFeedModel) headerView() string { return "Enter the channelid to load a RSS FEED?\n" }
func (m LoadRssFeedModel) footerView() string { return "\n(esc to quit)" }