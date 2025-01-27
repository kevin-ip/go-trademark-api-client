package go_markerapi_client

type TradeMarkService interface {
	// Check whether the search term is available
	// True if available, false otherwise
	IsAvailable(searchTerm string) (bool, error)
}
