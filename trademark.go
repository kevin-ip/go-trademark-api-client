package go_markerapi_client

import (
	"context"
)

type TrademarkService interface {
	// IsAvailable checks whether the search term is available and free to use
	// True if available, false if trademarked
	IsAvailable(ctx context.Context, searchTerm string) (bool, error)
}

func NewTrademarkService(vendor Vendor, apiKey string) TrademarkService {
	return NewUSPTOTradeMarkService(apiKey)
}

type Vendor int64

const (
	PENTIUM10_USPTO_TRADEBMARK Vendor = iota
)
