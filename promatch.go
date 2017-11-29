package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newProMatchService(sling *sling.Sling) *ProMatchService {
	return &ProMatchService{
		sling: sling.Path("proMatches"),
	}
}

// ProMatchService provides methods for accessing pro matches.
type ProMatchService struct {
	sling *sling.Sling
}

// ProMatch is a collection of stats about a pro match.
type ProMatch struct {
	MatchID       int64  `json:"match_id"`
	Duration      int    `json:"duration"`
	StartTime     int    `json:"start_time"`
	RadiantTeamID int    `json:"radiant_team_id"`
	RadiantName   string `json:"radiant_name"`
	DireTeamID    int    `json:"dire_team_id"`
	DireName      string `json:"dire_name"`
	LeagueID      int    `json:"leagueid"`
	LeagueName    string `json:"league_name"`
	SeriesID      int    `json:"series_id"`
	SeriesType    int    `json:"series_type"`
	RadiantScore  int    `json:"radiant_score"`
	DireScore     int    `json:"dire_score"`
	RadiantWin    bool   `json:"radiant_win"`
}

// Matches returns a collection about pro matches.
// https://docs.opendota.com/#tag/pro-matches%2Fpaths%2F~1proMatches%2Fget
func (s *ProMatchService) Matches() ([]ProMatch, *http.Response, error) {
	promatches := new([]ProMatch)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(promatches, apiError)
	return *promatches, resp, relevantError(err, *apiError)
}
