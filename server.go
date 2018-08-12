package main

import (
	"encoding/json"

	"google.golang.org/appengine/log"

	"net/http"
	"strings"

	youtube "google.golang.org/api/youtube/v3"
	"google.golang.org/appengine"
)

func youtubeHandler(w http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	if cachedlist == nil {

		client := getServerClient(ctx)
		service, err := youtube.New(client)
		handleError(ctx, err, "Error creating YouTube client")
		result := getYoutubeData(service, "snippet,contentDetails,statistics", "UCJS9pqu9BzkAMNTmzNMNhvg", ctx)
		cachedlist = prepdata(result)
		// cachedlist = &Youtubeitemlist{
		// 	Name:     "google Next",
		// 	Children: result,
		// }
	}
	response, err := json.Marshal(*cachedlist)
	if err != nil {
		log.Errorf(ctx, "Error!!!", err)
	}

	w.Header().Set("Contenty-Type", "applicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func prepdata(list []Youtubeitem) *Youtubedispitem {

	istio := Youtubedispitem{
		Name:     "istio",
		Children: []Youtubedispitem{},
	}
	gke := Youtubedispitem{
		Name:     "gke",
		Children: []Youtubedispitem{},
	}
	security := Youtubedispitem{
		Name:     "security",
		Children: []Youtubedispitem{}}
	migration := Youtubedispitem{
		Name:     "migration",
		Children: []Youtubedispitem{},
	}
	serverless := Youtubedispitem{
		Name:     "serverless",
		Children: []Youtubedispitem{},
	}

	oil_gas := Youtubedispitem{
		Name:     "oil&gas",
		Children: []Youtubedispitem{},
	}

	for _, v := range list {
		if strings.Contains(strings.ToLower(v.Name), "istio") {
			istio.Children = append(istio.Children, NewYoutubedispitem(v.Name, v.Contentdetails.VideoId))
		}
		if strings.Contains(strings.ToLower(v.Name), "gke") {
			gke.Children = append(gke.Children, NewYoutubedispitem(v.Name, v.Contentdetails.VideoId))
		}
		if strings.Contains(strings.ToLower(v.Name), "security") {
			security.Children = append(security.Children, NewYoutubedispitem(v.Name, v.Contentdetails.VideoId))
		}
		if strings.Contains(strings.ToLower(v.Name), "migration") {
			migration.Children = append(migration.Children, NewYoutubedispitem(v.Name, v.Contentdetails.VideoId))
		}
		if strings.Contains(strings.ToLower(v.Name), "serverless") {
			serverless.Children = append(serverless.Children, NewYoutubedispitem(v.Name, v.Contentdetails.VideoId))
		}
		if strings.Contains(strings.ToLower(v.Name), "oil") || strings.Contains(strings.ToLower(v.Name), "gas") || strings.Contains(strings.ToLower(v.Name), "schlumberger") {
			oil_gas.Children = append(oil_gas.Children, NewYoutubedispitem(v.Name, v.Contentdetails.VideoId))
		}

	}

	result := Youtubedispitem{
		Name: "Google Next '18",
		Children: []Youtubedispitem{
			istio,
			gke,
			security,
			migration,
			serverless,
			oil_gas,
		},
	}
	return &result

}
