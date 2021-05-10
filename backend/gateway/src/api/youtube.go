package api

import (
	"log"
	"net/http"
)

const youtubeKey = ""

//Search youtube vedio based on the input
func YoutebeSearch(input []string) map[string]string{

	//Create search term
	seacrhTerm := ""
	for index, term := range input {
		if index == len(input) - 1 {
			seacrhTerm = seacrhTerm + term
		} else {
			seacrhTerm =  seacrhTerm + term + ", "
		}
	}

	//Create youtube client
	client := &http.Client{
		Transport: &transport.APIKey{Key: youtubeKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	call := service.Search.List("id,snippet").
		Q(flag.String("query", "Google", seacrhTerm)).
		MaxResults(10)
	response, err := call.Do()
	handleError(err, "")

	//Retrieve the information from the response
	videos := make(map[string]string)
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
	}

	return videos
}