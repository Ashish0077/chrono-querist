package youtube

import (
	"context"
	"log"

	"google.golang.org/api/option"
	ytApi "google.golang.org/api/youtube/v3"
)

type YoutubeClient struct {
	ApiKey               []string
	currentIndexOfApiKey int
	service              *ytApi.Service
}

func NewYoutubeClient(apiKey []string) *YoutubeClient {
	if len(apiKey) < 1 {
		log.Fatalf("Missing API Keys")
	}

	svc, err := ytApi.NewService(context.Background(), option.WithAPIKey(apiKey[0]))
	if err != nil {
		log.Fatalf("Unable to initialize youtube client, %v", err)
	}

	return &YoutubeClient{
		ApiKey:               apiKey,
		service:              svc,
		currentIndexOfApiKey: 0,
	}
}

func (client *YoutubeClient) changeApiKey() error {
	numOfKeys := len(client.ApiKey)
	client.currentIndexOfApiKey = (client.currentIndexOfApiKey + 1) % numOfKeys

	service, err := ytApi.NewService(context.Background(), option.WithAPIKey(client.ApiKey[client.currentIndexOfApiKey]))
	if err != nil {
		log.Fatalf("Unable to initialize youtube client, %v", err)
	}
	client.service = service
	return nil
}
