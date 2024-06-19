package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	apiKey     = os.Getenv("YOUTUBE_API_KEY")
	playlistID = os.Getenv("YOUTUBE_PLAYLIST_ID")
)

func main() {
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	var allVideos []*youtube.PlaylistItem

	call := service.PlaylistItems.List([]string{"snippet"}).PlaylistId(playlistID).MaxResults(50)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	allVideos = append(allVideos, response.Items...)

	for response.NextPageToken != "" {
		call = service.PlaylistItems.List([]string{"snippet"}).PlaylistId(playlistID).MaxResults(50).PageToken(response.NextPageToken)
		response, err = call.Do()
		if err != nil {
			log.Fatalf("Error fetching next page of videos: %v", err)
		}
		allVideos = append(allVideos, response.Items...)
	}

	// Simplify the video data for output
	type SimpleVideo struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	var simpleVideos []SimpleVideo

	for _, item := range allVideos {
		video := SimpleVideo{
			Title: item.Snippet.Title,
			URL:   "https://www.youtube.com/watch?v=" + item.Snippet.ResourceId.VideoId,
		}
		simpleVideos = append(simpleVideos, video)
	}

	// Convert the simplified video data to JSON
	jsonData, err := json.MarshalIndent(simpleVideos, "", "  ")
	if err != nil {
		log.Fatalf("Error converting data to JSON: %v", err)
	}

	// Write the JSON data to a file named playlist_videos.json
	err = ioutil.WriteFile("playlist_videos.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON data to file: %v", err)
	}

	fmt.Println("Playlist data exported to playlist_videos.json")
}
