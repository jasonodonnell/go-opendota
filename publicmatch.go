package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newPublicMatchService(sling *sling.Sling) *PublicMatchService {
	return &PublicMatchService{
		sling: sling.Path("publicMatches"),
	}
}

// PublicMatchService provides methods for accessing public matches.
type PublicMatchService struct {
	sling *sling.Sling
}

// PublicMatchParam are the parameters for customizing queries.
type PublicMatchParam struct {
	MmrAscending    int   `url:"mmr_ascending,omitempty"`
	MmrDescending   int   `url:"mmr_descending,omitempty"`
	LessThanMatchID int64 `url:"less_than_match_id,omitempty"`
}

// PublicMatch is a collection about public matches.
type PublicMatch struct {
	MatchID     int64  `json:"match_id"`
	MatchSeqNum int64  `json:"match_seq_num"`
	RadiantWin  bool   `json:"radiant_win"`
	StartTime   int    `json:"start_time"`
	Duration    int    `json:"duration"`
	AvgMmr      int    `json:"avg_mmr"`
	NumMmr      int    `json:"num_mmr"`
	LobbyType   int    `json:"lobby_type"`
	GameMode    int    `json:"game_mode"`
	RadiantTeam string `json:"radiant_team"`
	DireTeam    string `json:"dire_team"`
}

// Matches returns information about public matches.
func (s *PublicMatchService) Matches(params *PublicMatchParam) ([]PublicMatch, *http.Response, error) {
	publicmatches := new([]PublicMatch)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(publicmatches, apiError)
	return *publicmatches, resp, relevantError(err, *apiError)
}
