package core

import (
	"github.com/mmcdole/gofeed"
	"github.com/darthpedroo/detoxtube/types"
)

type VideosLoader interface{
	LoadFeed(feedUrl string) (*gofeed.Feed, error)
	LoadVideos(feed *gofeed.Feed, maxVideos int) ([]types.Video, error)
	LoadTitle(feed *gofeed.Feed) (string, error)
}