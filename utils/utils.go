package utils

import (
	"os"
	"os/exec"
	"charm.land/bubbletea/v2"
)

func OpenApp(returnModel tea.Model, app string, args ...string) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command(app, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		_ = cmd.Run() // blocks until finished
		return returnModel
	}
}

func OpenInNewTerminal(returnModel tea.Model, app string, args ...string) tea.Cmd {
    return func() tea.Msg {
        var cmd *exec.Cmd

        fullArgs := append([]string{"--detach", app}, args...)
        cmd = exec.Command("kitty", fullArgs...)

        // We don't use Stdin/Stdout here because the new terminal handles its own IO
        _ = cmd.Start() 
        
        return returnModel
    }
}