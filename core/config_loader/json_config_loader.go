package core

import (
	"github.com/darthpedroo/detoxtube/types"
	//"encoding/json"
	)

type JsonConfigLoader struct {

}

	
func (c *JsonConfigLoader) LoadConfig() types.Config{

	loadedConfig := types.Config{
		VideoPlayer: "mpv",
		Channels: []types.Channel{
			{
				ChannelName: "Lawren Systems",
				FeedUrl: "https://www.youtube.com/feeds/videos.xml?channel_id=UCHkYOD-3fZbuGhwsADBd9ZQ",
			},
			{
				ChannelName: "Ken",
				FeedUrl: "https://www.youtube.com/feeds/videos.xml?channel_id=UCiFOL6V9KbvxfXvzdFSsqCw",
			},
		},
	}

	return loadedConfig
}