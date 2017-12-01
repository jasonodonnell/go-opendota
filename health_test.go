package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthService_Health(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"postgresUsage":{"metric":142794834092,"threshold":200000000000,"timestamp":1511193048},"redisUsage":{"metric":"1605460248","threshold":2500000000,"timestamp":1511193044},"parseDelay":{"metric":31678,"threshold":2000,"timestamp":1511193044},"cassandraUsage":{"metric":6808540713048,"threshold":8000000000000,"timestamp":1511193049},"seqNumDelay":{"metric":99,"threshold":10000,"timestamp":1511193049},"steamApi":{"metric":0,"threshold":1,"timestamp":1511193051}}`)
	})

	expected := Health{
		PostgresUsage: Usage{
			Metric:    142794834092,
			Threshold: 200000000000,
			Timestamp: 1511193048,
		},
		RedisUsage: RedisUsage{
			Metric:    "1605460248",
			Threshold: 2500000000,
			Timestamp: 1511193044,
		},
		ParseDelay: Usage{
			Metric:    31678,
			Threshold: 2000,
			Timestamp: 1511193044,
		},
		CassandraUsage: Usage{
			Metric:    6808540713048,
			Threshold: 8000000000000,
			Timestamp: 1511193049,
		},
		SeqNumDelay: Usage{
			Metric:    99,
			Threshold: 10000,
			Timestamp: 1511193049,
		},
		SteamAPI: Usage{
			Metric:    0,
			Threshold: 1,
			Timestamp: 1511193051,
		},
	}

	client := NewClient(httpClient)
	health, _, err := client.HealthService.Health()
	assert.Nil(t, err)
	assert.Equal(t, expected, health)
}
