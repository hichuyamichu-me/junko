package resolver

import (
	"context"

	"github.com/HichuYamichu/fetcher-api/fetcher"
)

// ChannelResolver : rosolves channels
type ChannelResolver struct {
	rpc     fetcher.GuildFetcherClient
	channel *fetcher.Channel
}

// ID : rosolves channel ID
func (c *ChannelResolver) ID(ctx context.Context) *string {
	return &c.channel.Id
}

// Name : resolves channel Name
func (c *ChannelResolver) Name(ctx context.Context) *string {
	return &c.channel.Name
}

// CreatedAt : resolves channel creation date
func (c *ChannelResolver) CreatedAt(ctx context.Context) *string {
	return &c.channel.CreatedAt
}

// Type : resolves channel Type
func (c *ChannelResolver) Type(ctx context.Context) *string {
	return &c.channel.Type
}

// Position : resolves channel Position
func (c *ChannelResolver) Position(ctx context.Context) *int32 {
	return &c.channel.Position
}

// Topic : resolves channel Topic
func (c *ChannelResolver) Topic(ctx context.Context) *string {
	return &c.channel.Topic
}

// Nsfw : resolves channel Nsfw value
func (c *ChannelResolver) Nsfw(ctx context.Context) *bool {
	return &c.channel.Nsfw
}

// Bitrate : resolves channel Bitrate
func (c *ChannelResolver) Bitrate(ctx context.Context) *int32 {
	return &c.channel.Bitrate
}

// UserLimit : resolves channel UserLimit
func (c *ChannelResolver) UserLimit(ctx context.Context) *int32 {
	return &c.channel.UserLimit
}

// ParentID : resolves channel ParentID
func (c *ChannelResolver) ParentID(ctx context.Context) *string {
	return &c.channel.ParentID
}

// RateLimitPerUser : resolves channel RateLimitPerUser
func (c *ChannelResolver) RateLimitPerUser(ctx context.Context) *int32 {
	return &c.channel.RateLimitPerUser
}