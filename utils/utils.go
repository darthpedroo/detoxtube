package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

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

func GetHome()(string, error){

	dirname, err := os.UserHomeDir()
    if err != nil {
        return "", fmt.Errorf("Error getting Home directory %s",err)
    }

	return dirname, nil
}

func SeedConfig(configPath string) error {
	data , _ := os.ReadFile("config_template.json")
	configFile := configPath + "/config.json"
    err := os.WriteFile(configFile, data, 0644)
	
	if err != nil {
		return err
	}

	return nil

}

func CreateConfigDir() error{

	home , err := GetHome()
	if err != nil {
		return err
	}

	configDirectory := home + "/.config/detoxtube"
	err = os.Mkdir(configDirectory,0777)

	if err != nil {
		return err
	}
	
	err = SeedConfig(configDirectory)
	
	if err != nil {
		return err
	}

	return nil
}

func FormatRelativeTime(input string) string {
    past, err := time.Parse(time.RFC3339, input)
    if err != nil {
        return "invalid date"
    }

    diff := time.Since(past)
    days := int(diff.Hours() / 24)

    switch {
    case days < 1:
        hours := int(diff.Hours())
        if hours == 1 {
            return "1 hour ago"
        }
        return fmt.Sprintf("%d hours ago", hours)

    case days < 7:
        if days == 1 {
            return "1 day ago"
        }
        return fmt.Sprintf("%d days ago", days)

    case days < 30:
        weeks := days / 7
        if weeks == 1 {
            return "1 week ago"
        }
        return fmt.Sprintf("%d weeks ago", weeks)

    case days < 365:
        months := days / 30 
        if months == 1 {
            return "1 month ago"
        }
        return fmt.Sprintf("%d months ago", months)

    default:
        years := days / 365
        if years == 1 {
            return "1 year ago"
        }
        return fmt.Sprintf("%d years ago", years)
    }
}

func WriteLog(message string){
	d1 := []byte(fmt.Sprintf("%s\n", message))
	path1 := filepath.Join("logs", "dat1.txt")
	data , _ := os.ReadFile(path1)

	data = append(data, d1...)
	
    _ = os.WriteFile(path1, data, 0644)
}