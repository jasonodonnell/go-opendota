package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

// HealthService provides methods for getting health stats
// about the OpenDota API.
type HealthService struct {
	sling *sling.Sling
}

func newHealthService(sling *sling.Sling) *HealthService {
	return &HealthService{
		sling: sling.Path("health"),
	}
}

// Health is a collection of health stats about the OpenDota API
// stack.
type Health health

type health struct {
	PostgresUsage  usage      `json:"postgresUsage"`
	RedisUsage     redisUsage `json:"redisUsage"`
	ParseDelay     usage      `json:"parseDelay"`
	CassandraUsage usage      `json:"cassandraUsage"`
	SeqNumDelay    usage      `json:"seqNumDelay"`
	SteamAPI       usage      `json:"steamApi"`
}

// Requires special struct due to a bug in OpenDota.
type redisUsage struct {
	Metric    string `json:"metric"`
	Threshold int64  `json:"threshold"`
	Timestamp int    `json:"timestamp"`
}

type usage struct {
	Metric    int64 `json:"metric"`
	Threshold int64 `json:"threshold"`
	Timestamp int   `json:"timestamp"`
}

// Health returns information about the current health of
// the OpenDota API stack.
func (s *HealthService) Health() (Health, *http.Response, error) {
	health := new(Health)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(health, apiError)
	return *health, resp, relevantError(err, *apiError)
}
