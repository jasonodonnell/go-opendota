package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamService_Top(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/teams", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"team_id":2163,"rating":1628.74,"wins":582,"losses":368,"last_match_time":1509224338,"name":"Team Liquid","tag":"Liquid","logo_url":"http://cloud-3.steamusercontent.com/ugc/858347654776522964/E70F0E063879154A1982B3C907D6A5DFDA183BF9/"}]`)
	})

	expected := []Team{
		Team{
			TeamID:        2163,
			Rating:        1628.74,
			Wins:          582,
			Losses:        368,
			LastMatchTime: 1509224338,
			Name:          "Team Liquid",
			Tag:           "Liquid",
			LogoURL:       "http://cloud-3.steamusercontent.com/ugc/858347654776522964/E70F0E063879154A1982B3C907D6A5DFDA183BF9/",
		},
	}

	client := NewClient(httpClient)
	teams, _, err := client.TeamService.Teams()
	assert.Nil(t, err)
	assert.Equal(t, expected, teams)
}
