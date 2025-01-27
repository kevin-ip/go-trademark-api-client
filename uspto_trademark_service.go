package go_markerapi_client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TrademarkAvailableResponse struct {
	Keyword   string
	Available string
}

type usptoTradeMarkService struct {
	urlFormat    string
	rapidAPIKey  string
	rapidAPIHost string
}

func NewUSPTOTradeMarkService(apiKey string) TradeMarkService {
	return &usptoTradeMarkService{
		urlFormat:    "https://uspto-trademark.p.rapidapi.com/v1/trademarkAvailable/%v",
		rapidAPIKey:  apiKey,
		rapidAPIHost: "uspto-trademark.p.rapidapi.com",
	}
}

func (t *usptoTradeMarkService) IsAvailable(searchTerm string) (bool, error) {
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
	err = json.Unmarshal(bodyBytes, &responses)
	if err != nil {
		return false, err
	}

	for _, res := range responses {
		log.Printf("res: %+v", res)
		if res.Keyword == searchTerm && res.Available == "no" {
			return false, nil
		}
	}
	return true, nil
}

func (t *usptoTradeMarkService) createRequest(searchTerm string) (*http.Request, error) {
	searchURL := fmt.Sprintf(t.urlFormat, searchTerm)
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", t.rapidAPIKey)
	req.Header.Add("x-rapidapi-host", t.rapidAPIHost)
	return req, err
}
