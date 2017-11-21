package opendota

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

func newTeamService(sling *sling.Sling) *TeamService {
	return &TeamService{
		sling: sling.Path("teams/"),
	}
}

// TeamService provides methods for accessing teams
// endpoints.
type TeamService struct {
	sling *sling.Sling
}

// TeamParam is the parameter for specifying a team.
type TeamParam struct {
	TeamID int64 `url:"team_id,omitempty"`
}

// Heroes is a collection of heroes played
// by a team.
type Heroes struct {
	HeroID        int    `json:"hero_id"`
	LocalizedName string `json:"localized_name"`
	GamesPlayed   int    `json:"games_played"`
	Wins          int    `json:"wins"`
}

// TeamMatch is a collection of matches played
// by a team.
type TeamMatch struct {
	MatchID    int64  `json:"match_id"`
	RadiantWin bool   `json:"radiant_win"`
	Radiant    bool   `json:"radiant"`
	Duration   int    `json:"duration"`
	StartTime  int    `json:"start_time"`
	LeagueID   int    `json:"leagueid"`
	LeagueName string `json:"league_name"`
	Cluster    int    `json:"cluster"`
}

// TeamPlayers is a collection of people who have played on a team.
type TeamPlayers struct {
	AccountID           int    `json:"account_id"`
	Name                string `json:"name"`
	GamesPlayed         int    `json:"games_played"`
	Wins                int    `json:"wins"`
	IsCurrentTeamMember bool   `json:"is_current_team_member"`
}

// Team is a collection of stats about a team.
type Team struct {
	TeamID        int     `json:"team_id"`
	Rating        float64 `json:"rating"`
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	LastMatchTime int     `json:"last_match_time"`
	Name          string  `json:"name"`
	Tag           string  `json:"tag"`
	LogoURL       string  `json:"logo_url"`
}

// Teams returns a collection of teams.
func (s *TeamService) Teams() ([]Team, *http.Response, error) {
	teams := new([]Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(teams, apiError)
	return *teams, resp, relevantError(err, *apiError)
}

// Team returns a specific team.
func (s *TeamService) Team(params *TeamParam) (Team, *http.Response, error) {
	team := new(Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(int(params.TeamID))).Receive(team, apiError)
	return *team, resp, relevantError(err, *apiError)
}

// Matches returns a collection of matches for a specific
// team.
func (s *TeamService) Matches(params *TeamParam) ([]TeamMatch, *http.Response, error) {
	matches := new([]TeamMatch)
	apiError := new(APIError)
	path := strconv.Itoa(int(params.TeamID)) + "/matches"
	resp, err := s.sling.New().Get(path).Receive(matches, apiError)
	return *matches, resp, relevantError(err, *apiError)
}

// Players returns a collection of people that played
// on a specific team.
func (s *TeamService) Players(params *TeamParam) ([]TeamPlayers, *http.Response, error) {
	players := new([]TeamPlayers)
	apiError := new(APIError)
	path := strconv.Itoa(int(params.TeamID)) + "/players"
	resp, err := s.sling.New().Get(path).Receive(players, apiError)
	return *players, resp, relevantError(err, *apiError)
}

// Heroes returns a collection of stats about the heroes
// played by a specific team.
func (s *TeamService) Heroes(params *TeamParam) ([]Heroes, *http.Response, error) {
	heroes := new([]Heroes)
	apiError := new(APIError)
	path := strconv.Itoa(int(params.TeamID)) + "/heroes"
	resp, err := s.sling.New().Get(path).Receive(heroes, apiError)
	return *heroes, resp, relevantError(err, *apiError)
}
