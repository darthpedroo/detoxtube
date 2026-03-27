package core

import (
	"github.com/darthpedroo/detoxtube/types"
	"github.com/mmcdole/gofeed"
)

type VideosLoader interface {
	LoadFeed(feedUrl string) (*gofeed.Feed, error)
	LoadTitle(feed *gofeed.Feed) (string, error)
	LoadVideos(feed *gofeed.Feed, maxVideos int) ([]types.Video, error)
}
