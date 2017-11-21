package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newReplayService(sling *sling.Sling) *ReplayService {
	return &ReplayService{
		sling: sling.Path("replays"),
	}
}

// ReplayService provides methods for accesing information about
// match replays.
type ReplayService struct {
	sling *sling.Sling
}

// ReplayParam allows replays to be queried by MatchIDs.
type ReplayParam struct {
	MatchID []int `url:"match_id"`
}

// Replay is a collection of information about a specific replay.
type Replay struct {
	MatchID    int64 `json:"match_id"`
	Cluster    int   `json:"cluster"`
	ReplaySalt int   `json:"replay_salt"`
	SeriesID   int   `json:"series_id"`
	SeriesType int   `json:"series_type"`
}

// Replays returns a collection of match replays.
func (s *ReplayService) Replays(params *ReplayParam) ([]Replay, *http.Response, error) {
	replays := new([]Replay)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(replays, apiError)
	return *replays, resp, relevantError(err, *apiError)
}
