package youtube

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	ytApi "google.golang.org/api/youtube/v3"
)

// Ref: https://developers.google.com/youtube/v3/docs/search/list#go

func (client *YoutubeClient) SearchByQuery(query string, maxResults int64, after time.Time) ([]*ytApi.SearchResult, error) {
	request := client.service.Search.List([]string{"id", "snippet"}).
		Q(query).MaxResults(maxResults).Type("video").Order("date").
		PublishedAfter(after.Format(time.RFC3339))
	client.changeApiKey() // After every request updating API Key

	response, err := request.Do()
	if err != nil {
		// TODO: Could add a check here on which type of error it is and then change the api key accordingly
		log.Error("Error occured while searching by query: ", err)
		return nil, err
	}

	return response.Items, err
}
