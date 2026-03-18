package utils

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"charm.land/bubbletea/v2"
	"github.com/darthpedroo/detoxtube/types"
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

func CreateRssFeedFromChannelId(channelId string)(rssFeed string){
	//https://stackoverflow.com/questions/19795987/youtube-channel-and-playlist-id-prefixes/77816885#77816885

	newChannelId := strings.TrimPrefix(channelId, "UC")
	return fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?playlist_id=UULF%v",newChannelId) 
}

func SortSubscriptions(subscriptions []types.Channel, videoSort types.VideoSort, order types.Order) (sortedSubscriptions []types.Channel) {
    if videoSort == types.Alphabetically {
        sort.Slice(subscriptions, func(i, j int) bool {
            titleI := strings.ToLower(subscriptions[i].ChannelName)
            titleJ := strings.ToLower(subscriptions[j].ChannelName)
            
            if order == types.Descending {
                return titleI > titleJ
            }
            return titleI < titleJ
        })

		return subscriptions
    } else {
		return subscriptions
	}
}