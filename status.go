package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newStatusService(sling *sling.Sling) *StatusService {
	return &StatusService{
		sling: sling.Path("status"),
	}
}

// StatusService provides a method for accessing OpenDota API stats.
type StatusService struct {
	sling *sling.Sling
}

// Status represents the stats for the OpenDota API.
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
	Retriever               []HostnameCounts `json:"retriever"`
	APIPaths                []HostnameCounts `json:"api_paths"`
	LastAdded               []MatchStatus    `json:"last_added"`
	LastParsed              []MatchStatus    `json:"last_parsed"`
	LoadTimes               map[string]int   `json:"load_times"`
	Health                  Health           `json:"health"`
}

type HostnameCounts struct {
	Hostname string `json:"hostname"`
	Count    string `json:"count"`
}

type MatchStatus struct {
	MatchID   int64 `json:"match_id"`
	Duration  int   `json:"duration"`
	StartTime int   `json:"start_time"`
}

// Status returns the current status of the OpenDota API.
// https://docs.opendota.com/#tag/status%2Fpaths%2F~1status%2Fget
func (s *StatusService) Status() (Status, *http.Response, error) {
	status := new(Status)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(status, apiError)
	return *status, resp, relevantError(err, *apiError)
}
