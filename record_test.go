package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordService_Records(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/records/kills", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":"3567354666","start_time":"1510938727","hero_id":"34","score":"103"}]`)
	})

	expected := []Record{
		Record{
			MatchID:   "3567354666",
			StartTime: "1510938727",
			HeroID:    "34",
			Score:     "103",
		},
	}

	client := NewClient(httpClient)
	records, _, err := client.RecordService.Records("kills")
	assert.Nil(t, err)
	assert.Equal(t, expected, records)
}
