package go_markerapi_client

import (
	"context"

	"github.com/kevin-ip/go-trademark-api-client/uspto_trademark"
)

type TrademarkService interface {
	// IsAvailable checks whether the search term is available and free to use
	// True if available, false if trademarked
	IsAvailable(ctx context.Context, searchTerm string) (bool, error)
}

func NewTrademarkService(vendor Vendor, apiKey string) TrademarkService {
	return uspto_trademark.NewUSPTOTradeMarkService(apiKey)
}

type Vendor int64

const (
	PENTIUM10_USPTO_TRADEBMARK Vendor = iota
)
