package main

import (
	"google.golang.org/api/youtube/v3"
)

type Youtubeitem struct {
	Name             string                             `json:"name"`
	Description      string                             `json:"description"`
	ID               string                             `json:"id"`
	ThumbnailDetails youtube.ThumbnailDetails           `json:"thumbnaildetails"`
	Contentdetails   youtube.PlaylistItemContentDetails `json:"children"`
}

type Youtubeitemlist struct {
	Name     string        `json:"name"`
	Children []Youtubeitem `json:"children"`
}

type Youtubedispitem struct {
	Name     string            `json:"name"`
	ID       string            `json:"id"`
	Children []Youtubedispitem `json:"children"`
}

func NewYoutubedispitem(name string, id string) Youtubedispitem {
	return Youtubedispitem{
		Name: name,
		ID:   id,
	}
}
