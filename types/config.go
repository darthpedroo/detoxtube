package types

type Channel struct {
	ChannelName string
	FeedUrl string
}

type Config struct {
	VideoPlayer string
	Channels []Channel
}