package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicMatchService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/publicMatches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":3567797504,"match_seq_num":3103150349,"radiant_win":false,"start_time":1510962873,"duration":890,"avg_mmr":4015,"num_mmr":7,"lobby_type":0,"game_mode":4,"radiant_team":"97,14,42,96,92","dire_team":"39,9,20,36,78"}]`)
	})

	expected := []PublicMatch{
		{
			MatchID:     3567797504,
			MatchSeqNum: 3103150349,
			RadiantWin:  false,
			StartTime:   1510962873,
			Duration:    890,
			AvgMmr:      4015,
			NumMmr:      7,
			LobbyType:   0,
			GameMode:    4,
			RadiantTeam: "97,14,42,96,92",
			DireTeam:    "39,9,20,36,78",
		},
	}

	publcMatchParam := &PublicMatchParam{}

	client := NewClient(httpClient)
	publicmatches, _, err := client.PublicMatchService.Matches(publcMatchParam)
	assert.Nil(t, err)
	assert.Equal(t, expected, publicmatches)
}
