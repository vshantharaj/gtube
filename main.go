package main

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	_ "github.com/lib/pq"
	youtube "google.golang.org/api/youtube/v3"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

var cachedlist *Youtubedispitem

func getServerClient(ctx context.Context) *http.Client {

	transport := &oauth2.Transport{
		Source: google.AppEngineTokenSource(ctx, youtube.YoutubeReadonlyScope),
		Base:   &urlfetch.Transport{Context: ctx},
	}
	client := &http.Client{Transport: transport}

	return client
}

func handleError(ctx context.Context, err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Errorf(ctx, message+": %v", err)
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./public/")))
	http.HandleFunc("/api/youtube/", youtubeHandler)
	appengine.Main()

}
