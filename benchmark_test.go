package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBenchmarkService_Benchmarks(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/benchmarks", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"hero_id":36,"result":{"gold_per_min":[{"percentile":0.1,"value":268}],"xp_per_min":[{"percentile":0.1,"value":334}],"kills_per_min":[{"percentile":0.1,"value":0.07518796992481203}],"last_hits_per_min":[{"percentile":0.1,"value":1.4967177242888403}],"hero_damage_per_min":[{"percentile":0.1,"value":260.321608040201}],"hero_healing_per_min":[{"percentile":0.1,"value":25.52536231884058}],"tower_damage":[{"percentile":0.1,"value":12}]}}`)
	})

	expected := Benchmark{
		HeroID: 36,
		Result: BenchmarkResult{
			GoldPerMin: []BenchmarkPercentile{
				{
					Percentile: 0.1,
					Value:      268,
				},
			},
			XpPerMin: []BenchmarkPercentile{
				{
					Percentile: 0.1,
					Value:      334,
				},
			},
			KillsPerMin: []BenchmarkPercentile{
				{
					Percentile: 0.1,
					Value:      0.07518796992481203,
				},
			},
			LastHitsPerMin: []BenchmarkPercentile{
				{
					Percentile: 0.1,
					Value:      1.4967177242888403,
				},
			},
			HeroDamagePerMin: []BenchmarkPercentile{
				{
					Percentile: 0.1,
					Value:      260.321608040201,
				},
			},
			HeroHealingPerMin: []BenchmarkPercentile{
				{
					Percentile: 0.1,
					Value:      25.52536231884058,
				},
			},
			TowerDamage: []BenchmarkPercentile{
				{
					Percentile: 0.1,
					Value:      12,
				},
			},
		},
	}

	client := NewClient(httpClient)
	benchmark, _, err := client.BenchmarkService.Benchmarks(36)
	assert.Nil(t, err)
	assert.Equal(t, expected, benchmark)
}
