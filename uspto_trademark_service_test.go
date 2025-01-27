package go_markerapi_client

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestIsAvailable(t *testing.T) {
	err := godotenv.Load()
	require.NoError(t, err)

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Skip("Skip, no API_KEY environment variable found")
	}

	for _, testCase := range []struct {
		name       string
		searchTerm string
		expected   bool
	}{
		{
			name:       "a known trademark should not be available",
			searchTerm: "google",
			expected:   false,
		},
		{
			name:       "an unknown trademark should be available",
			searchTerm: "_asdf_",
			expected:   true,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			service := NewUSPTOTradeMarkService(apiKey)
			actual, err := service.IsAvailable(testCase.searchTerm)
			require.NoErrorf(t, err, "Error: %v", err)
			require.Equal(t, testCase.expected, actual)
		})
	}
}
