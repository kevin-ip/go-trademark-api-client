package go_markerapi_client

import "context"

type TradeMarkService interface {
	// IsAvailable Check whether the search term is available
	// True if available, false otherwise
	IsAvailable(ctx context.Context, searchTerm string) (bool, error)
}
