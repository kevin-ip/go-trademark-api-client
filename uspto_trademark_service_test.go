package go_markerapi_client

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestIsAvailable(t *testing.T) {
	err := godotenv.Load(".env")
	require.NoError(t, err)

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Skip("Skip, no API_KEY environment variable found")
	}

	for _, testCase := range []struct {
		name         string
		searchTerm   string
		expected     bool
		errorMessage string
	}{
		// {
		// 	name:       "a known trademark should not be available",
		// 	searchTerm: "google",
		// 	expected:   false,
		// },
		// {
		// 	name:       "an unknown trademark should be available",
		// 	searchTerm: "_asdf_",
		// 	expected:   true,
		// },
		// {
		// 	name:         "an empty searchTerm should not be available",
		// 	searchTerm:   "",
		// 	expected:     false,
		// 	errorMessage: "search term is empty",
		// },
		// {
		// 	name:         "happy campers, happy trails",
		// 	searchTerm:   "happy campers, happy trails",
		// },
		{
			name:       "HAPPY CAMPERS",
			searchTerm: "HAPPY CAMPERS",
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			service := NewUSPTOTradeMarkService(apiKey)
			actual, err := service.IsAvailable(context.Background(), testCase.searchTerm)
			if testCase.errorMessage != "" {
				require.EqualError(t, err, testCase.errorMessage)
			} else {
				require.NoErrorf(t, err, "Error: %v", err)
			}
			require.Equal(t, testCase.expected, actual)
		})
	}
}
