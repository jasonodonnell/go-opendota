package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

// StatusService provides methods for getting stats
// about the OpenDota API.
type StatusService struct {
	sling *sling.Sling
}

func newStatusService(sling *sling.Sling) *StatusService {
	return &StatusService{
		sling: sling.Path("status"),
	}
}

// Status is a collection of stats about the OpenDota API.
type Status struct {
	UserPlayers             int              `json:"user_players"`
	TrackedPlayers          int              `json:"tracked_players"`
	MatchesLastDay          int              `json:"matches_last_day"`
	MatchesLastHour         int              `json:"matches_last_hour"`
	RetrieverMatchesLastDay int              `json:"retriever_matches_last_day"`
	ParsedMatchesLastDay    int              `json:"parsed_matches_last_day"`
	RequestsLastDay         int              `json:"requests_last_day"`
	APIHitsLastDay          int              `json:"api_hits_last_day"`
	APIHitsUILastDay        int              `json:"api_hits_ui_last_day"`
	FhQueue                 int              `json:"fhQueue"`
	GcQueue                 int              `json:"gcQueue"`
	MmrQueue                int              `json:"mmrQueue"`
	Retriever               []hostnameCounts `json:"retriever"`
	APIPaths                []hostnameCounts `json:"api_paths"`
	LastAdded               []matchStatus    `json:"last_added"`
	LastParsed              []matchStatus    `json:"last_parsed"`
	LoadTimes               map[string]int   `json:"load_times"`
	Health                  health           `json:"health"`
}

type hostnameCounts struct {
	Hostname string `json:"hostname"`
	Count    string `json:"count"`
}

type matchStatus struct {
	MatchID   int64 `json:"match_id"`
	Duration  int   `json:"duration"`
	StartTime int   `json:"start_time"`
}

// Status returns information about the current status of
// the OpenDota API.
func (s *StatusService) Status() (Status, *http.Response, error) {
	status := new(Status)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(status, apiError)
	return *status, resp, relevantError(err, *apiError)
}
