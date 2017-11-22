package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProMatchService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/proMatches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":3565732649,"duration":2388,"start_time":1510866000,"radiant_team_id":36,"radiant_name":"Natus Vincere","dire_team_id":2586976,"dire_name":"OG","leagueid":5627,"league_name":"DreamLeague season 8","series_id":178590,"series_type":1,"radiant_score":25,"dire_score":12,"radiant_win":true}]`)
	})

	expected := []ProMatch{
		ProMatch{
			MatchID:       3565732649,
			Duration:      2388,
			StartTime:     1510866000,
			RadiantTeamID: 36,
			RadiantName:   "Natus Vincere",
			DireTeamID:    2586976,
			DireName:      "OG",
			LeagueID:      5627,
			LeagueName:    "DreamLeague season 8",
			SeriesID:      178590,
			SeriesType:    1,
			RadiantScore:  25,
			DireScore:     12,
			RadiantWin:    true,
		},
	}

	client := NewClient(httpClient)
	promatches, _, err := client.ProMatchService.Matches()
	assert.Nil(t, err)
	assert.Equal(t, expected, promatches)
}
