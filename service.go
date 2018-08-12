package main

import (
	"fmt"
	"strings"

	"google.golang.org/appengine/log"

	"golang.org/x/net/context"
	youtube "google.golang.org/api/youtube/v3"
)

func getYoutubeData(service *youtube.Service, part string, forUsername string, ctx context.Context) []Youtubeitem {
	call := service.Channels.List(part)
	call = call.Id(forUsername)
	response, err := call.Do()
	handleError(ctx, err, "")

	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+"with upload id '%s'"+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].ContentDetails.RelatedPlaylists.Uploads,
		response.Items[0].Statistics.ViewCount))

	return getPlaylistvidiews(service, "snippet,contentDetails", response.Items[0].ContentDetails.RelatedPlaylists.Uploads, ctx)

}

func getPlaylistvidiews(service *youtube.Service, part string, playlistid string, ctx context.Context) []Youtubeitem {
	call := service.PlaylistItems.List(part)
	call = call.PlaylistId(playlistid)
	call = call.MaxResults(49)
	response, err := call.Do()
	handleError(ctx, err, "")
	result := make([]Youtubeitem, 0)
	for nexttoekn := response.NextPageToken; response.NextPageToken != ""; {
		response, err = call.PageToken(nexttoekn).Do()
		nexttoekn = response.NextPageToken
		handleError(ctx, err, "")
		log.Infof(ctx, "firstlop", len(response.Items), response.NextPageToken)
		for _, item := range response.Items {
			if strings.Contains(item.Snippet.Title, "(Cloud Next '18)") {
				//fmt.Println(item.Id, ": ", item.Snippet.Title,item.Snippet.ResourceId)

				result = append(result, Youtubeitem{
					Name:             item.Snippet.Title,
					ID:               item.Id,
					Description:      item.Snippet.Description,
					Contentdetails:   *item.ContentDetails,
					ThumbnailDetails: *item.Snippet.Thumbnails,
				})
			}

		}
	}
	return result
}
