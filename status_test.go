package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusService_Status(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"user_players":545758,"tracked_players":70289,"matches_last_day":865758,"matches_last_hour":55340,"retriever_matches_last_day":78746,"parsed_matches_last_day":75822,"requests_last_day":148980,"api_hits_last_day":3178674,"api_hits_ui_last_day":901771,"fhQueue":0,"gcQueue":0,"mmrQueue":0,"retriever":[{"hostname":"retriever9","count":"1"}],"api_paths":[{"hostname":"benchmarks","count":"1"}],"last_added":[{"match_id":3572944537,"duration":2460,"start_time":1511187728}],"last_parsed":[{"match_id":3572868667,"duration":3892,"start_time":1511185271}],"load_times":{"50":28},"health":{"postgresUsage":{"metric":142788255916,"threshold":200000000000,"timestamp":1511190988},"redisUsage":{"metric":"1568683016","threshold":2500000000,"timestamp":1511190984},"parseDelay":{"metric":28866,"threshold":2000,"timestamp":1511190990},"cassandraUsage":{"metric":6881122290934,"threshold":8000000000000,"timestamp":1511190987},"seqNumDelay":{"metric":-4794,"threshold":10000,"timestamp":1511190989},"steamApi":{"metric":0,"threshold":1,"timestamp":1511190988}}}`)
	})

	expected := Status{
		UserPlayers:             545758,
		TrackedPlayers:          70289,
		MatchesLastDay:          865758,
		MatchesLastHour:         55340,
		RetrieverMatchesLastDay: 78746,
		ParsedMatchesLastDay:    75822,
		RequestsLastDay:         148980,
		APIHitsLastDay:          3178674,
		APIHitsUILastDay:        901771,
		FhQueue:                 0,
		GcQueue:                 0,
		MmrQueue:                0,
		Retriever: []hostnameCounts{
			hostnameCounts{
				Hostname: "retriever9",
				Count:    "1",
			},
		},
		APIPaths: []hostnameCounts{
			hostnameCounts{
				Hostname: "benchmarks",
				Count:    "1",
			},
		},
		LastAdded: []matchStatus{
			matchStatus{
				MatchID:   3572944537,
				Duration:  2460,
				StartTime: 1511187728,
			},
		},
		LastParsed: []matchStatus{
			matchStatus{
				MatchID:   3572868667,
				Duration:  3892,
				StartTime: 1511185271,
			},
		},
		LoadTimes: map[string]int{
			"50": 28,
		},
		Health: Health{
			PostgresUsage: usage{
				Metric:    142788255916,
				Threshold: 200000000000,
				Timestamp: 1511190988,
			},
			RedisUsage: redisUsage{
				Metric:    "1568683016",
				Threshold: 2500000000,
				Timestamp: 1511190984,
			},
			ParseDelay: usage{
				Metric:    28866,
				Threshold: 2000,
				Timestamp: 1511190990,
			},
			CassandraUsage: usage{
				Metric:    6881122290934,
				Threshold: 8000000000000,
				Timestamp: 1511190987,
			},
			SeqNumDelay: usage{
				Metric:    -4794,
				Threshold: 10000,
				Timestamp: 1511190989,
			},
			SteamAPI: usage{
				Metric:    0,
				Threshold: 1,
				Timestamp: 1511190988,
			},
		},
	}

	client := NewClient(httpClient)
	status, _, err := client.StatusService.Status()
	assert.Nil(t, err)
	assert.Equal(t, expected, status)
}
