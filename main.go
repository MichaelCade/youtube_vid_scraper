package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	playlistID := os.Getenv("YOUTUBE_PLAYLIST_ID")

	service, err := youtubeService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	call := service.PlaylistItems.List([]string{"snippet"}).PlaylistId(playlistID).MaxResults(25)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	videos := make([]map[string]string, 0)
	for _, item := range response.Items {
		video := map[string]string{
			"title": item.Snippet.Title,
			"url":   fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Snippet.ResourceId.VideoId),
		}
		videos = append(videos, video)
	}

	file, err := os.Create("playlist_videos.json")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(videos)
	if err != nil {
		log.Fatalf("Error encoding videos to JSON: %v", err)
	}

	fmt.Println("Exported playlist video data to playlist_videos.json")
}

func youtubeService(ctx context.Context, opts ...option.ClientOption) (*youtube.Service, error) {
	return youtube.NewService(ctx, opts...)
}
