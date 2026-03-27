package core

import (
	"fmt"
	"github.com/darthpedroo/detoxtube/types"
	"github.com/mmcdole/gofeed"
)

type GoFeedVideosLoader struct {
}

func (g *GoFeedVideosLoader) LoadFeed(feedUrl string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedUrl)
	if err != nil {
		return nil, fmt.Errorf("Error loading videos with gofeed %v" , err)
	}
	return feed, nil
}

func (g *GoFeedVideosLoader) LoadTitle(feed *gofeed.Feed) (string, error) {
	return feed.Title, nil
}

func (g *GoFeedVideosLoader) LoadVideos(feed *gofeed.Feed, maxVideos int) ([]types.Video, error) {
	var fetchedVideos []types.Video
	var loadedViedos int
	for _, item := range feed.Items {
		newVideo := types.Video{
			item.Title,
			item.Link,
			item.Published,
			item.Author.Name,
		}
		fetchedVideos = append(fetchedVideos, newVideo)
		loadedViedos += 1
		if loadedViedos == maxVideos {
			break
		}
	}
	return fetchedVideos, nil
}
