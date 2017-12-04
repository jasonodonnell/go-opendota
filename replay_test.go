package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplayService_Replays(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/replays", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":3559037317,"cluster":123,"replay_salt":897752925,"series_id":177752,"series_type":1},{"match_id":3573137571,"cluster":133,"replay_salt":1678969045,"series_id":180022,"series_type":1}]`)
	})

	expected := []Replay{
		{
			MatchID:    3559037317,
			Cluster:    123,
			ReplaySalt: 897752925,
			SeriesID:   177752,
			SeriesType: 1,
		},
		{
			MatchID:    3573137571,
			Cluster:    133,
			ReplaySalt: 1678969045,
			SeriesID:   180022,
			SeriesType: 1,
		},
	}

	client := NewClient(httpClient)
	replays, _, err := client.ReplayService.Replays([]int{3559037317, 3573137571})
	assert.Nil(t, err)
	assert.Equal(t, expected, replays)
}
