package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeagueService_Leagues(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/leagues", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"leagueid":2019,"ticket":"econ/leagues/subscriptions_ggleagues2ticket","banner":"econ/leagues/subscriptions_ggleagues2ticket_ingame","tier":"professional","name":"GG League Season 2 Ticket"}]`)
	})

	expected := []League{
		{
			LeagueID: 2019,
			Ticket:   "econ/leagues/subscriptions_ggleagues2ticket",
			Banner:   "econ/leagues/subscriptions_ggleagues2ticket_ingame",
			Tier:     "professional",
			Name:     "GG League Season 2 Ticket",
		},
	}

	client := NewClient(httpClient)
	leagues, _, err := client.LeagueService.Leagues()
	assert.Nil(t, err)
	assert.Equal(t, expected, leagues)
}
