package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchService_Search(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"account_id":74335542,"avatarfull":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/94/94c6b24d3cae823063c4e06b73016256a89fd13a_full.jpg","personaname":"Dr. LoveWizard","last_match_time":"2017-11-10T22:04:19.000Z","similarity":0.6875}]`)
	})

	expected := []Search{
		{
			AccountID:     74335542,
			AvatarFull:    "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/94/94c6b24d3cae823063c4e06b73016256a89fd13a_full.jpg",
			Personaname:   "Dr. LoveWizard",
			LastMatchTime: "2017-11-10T22:04:19.000Z",
			Similarity:    0.6875,
		},
	}

	params := &SearchParams{
		Similarity: 0.51,
	}

	client := NewClient(httpClient)
	search, _, err := client.SearchService.Search("drlovewizard", params)
	assert.Nil(t, err)
	assert.Equal(t, expected, search)
}
