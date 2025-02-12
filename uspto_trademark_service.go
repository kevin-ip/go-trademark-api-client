package go_trademark_api_client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type TrademarkAvailableResponse struct {
	Keyword   string
	Available string
}

// See https://rapidapi.com/pentium10/api/uspto-trademark
type usptoTrademarkService struct {
	urlFormat    string
	rapidAPIKey  string
	rapidAPIHost string
}

func NewUSPTOTradeMarkService(apiKey string) TrademarkService {
	return &usptoTrademarkService{
		urlFormat:    "https://uspto-trademark.p.rapidapi.com/v1/trademarkAvailable/%v",
		rapidAPIKey:  apiKey,
		rapidAPIHost: "uspto-trademark.p.rapidapi.com",
	}
}

func (t *usptoTrademarkService) IsAvailable(ctx context.Context, searchTerm string) (bool, error) {
	if searchTerm == "" {
		return false, fmt.Errorf("search term is empty")
	}

	req, err := t.createRequest(searchTerm)
	if err != nil {
		return false, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer func() {
		if err = res.Body.Close(); err != nil {
			log.Printf("Closing error: %v", err)
		}
	}()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	var responses []TrademarkAvailableResponse
	if err = json.Unmarshal(bodyBytes, &responses); err != nil {
		return false, fmt.Errorf("error unmarshalling response body (%v); %v", string(bodyBytes), err)
	}

	for _, res := range responses {
		if strings.ToLower(res.Keyword) == strings.ToLower(searchTerm) && res.Available == "no" {
			return false, nil
		}
	}
	return true, nil
}

func (t *usptoTrademarkService) createRequest(searchTerm string) (*http.Request, error) {
	searchURL := fmt.Sprintf(t.urlFormat, searchTerm)
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", t.rapidAPIKey)
	req.Header.Add("x-rapidapi-host", t.rapidAPIHost)
	return req, err
}
