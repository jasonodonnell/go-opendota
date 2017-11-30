package opendota

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

func newBenchmarkService(sling *sling.Sling) *BenchmarkService {
	return &BenchmarkService{
		sling: sling.Path("benchmarks"),
	}
}

// BenchmarkService provides a method for benchmark
// statistics in OpenDota.
type BenchmarkService struct {
	sling *sling.Sling
}

type benchmarkParam struct {
	heroID string `url:"hero_id"`
}

// Benchmark holds a collection of benchmarks for a hero.
type Benchmark struct {
	HeroID int             `json:"hero_id"`
	Result benchmarkResult `json:"result"`
}

type benchmark struct {
	Percentile float64 `json:"percentile"`
	Value      float64 `json:"value"`
}

type benchmarkResult struct {
	GoldPerMin        []benchmark `json:"gold_per_min"`
	XpPerMin          []benchmark `json:"xp_per_min"`
	KillsPerMin       []benchmark `json:"kills_per_min"`
	LastHitsPerMin    []benchmark `json:"last_hits_per_min"`
	HeroDamagePerMin  []benchmark `json:"hero_damage_per_min"`
	HeroHealingPerMin []benchmark `json:"hero_healing_per_min"`
	TowerDamage       []benchmark `json:"tower_damage"`
}

// Benchmarks takes a Hero ID and returns the corresponding benchmarks
// for that hero.
// https://docs.opendota.com/#tag/benchmarks%2Fpaths%2F~1benchmarks%2Fget
func (s *BenchmarkService) Benchmarks(heroID int) (Benchmark, *http.Response, error) {
	param := &benchmarkParam{}
	param.heroID = strconv.Itoa(heroID)
	benchmarks := new(Benchmark)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(param).Receive(benchmarks, apiError)
	return *benchmarks, resp, relevantError(err, *apiError)
}
