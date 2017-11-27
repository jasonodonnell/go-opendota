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

// TeamHeroes is a collection of heroes played by a team.
type TeamHeroes struct {
	HeroID        int    `json:"hero_id"`
	LocalizedName string `json:"localized_name"`
	GamesPlayed   int    `json:"games_played"`
	Wins          int    `json:"wins"`
}

// TeamMatch is a collection of matches played by a team.
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

// Heroes returns a collection of stats about the heroes played by a specific team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D~1heroes%2Fget
func (s *TeamService) Heroes(teamID int64) ([]TeamHeroes, *http.Response, error) {
	heroes := new([]TeamHeroes)
	apiError := new(APIError)
	path := strconv.Itoa(int(teamID)) + "/heroes"
	resp, err := s.sling.New().Get(path).Receive(heroes, apiError)
	return *heroes, resp, relevantError(err, *apiError)
}

// Matches returns a collection of matches for a specific team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D~1matches%2Fget
func (s *TeamService) Matches(teamID int64) ([]TeamMatch, *http.Response, error) {
	matches := new([]TeamMatch)
	apiError := new(APIError)
	path := strconv.Itoa(int(teamID)) + "/matches"
	resp, err := s.sling.New().Get(path).Receive(matches, apiError)
	return *matches, resp, relevantError(err, *apiError)
}

// Players returns a collection of people that played on a specific team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D~1players%2Fget
func (s *TeamService) Players(teamID int64) ([]TeamPlayers, *http.Response, error) {
	players := new([]TeamPlayers)
	apiError := new(APIError)
	path := strconv.Itoa(int(teamID)) + "/players"
	resp, err := s.sling.New().Get(path).Receive(players, apiError)
	return *players, resp, relevantError(err, *apiError)
}

// Team returns a collection for a specific team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D%2Fget
func (s *TeamService) Team(teamID int64) (Team, *http.Response, error) {
	team := new(Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(int(teamID))).Receive(team, apiError)
	return *team, resp, relevantError(err, *apiError)
}

// Teams returns a collection of teams.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams%2Fget
func (s *TeamService) Teams() ([]Team, *http.Response, error) {
	teams := new([]Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(teams, apiError)
	return *teams, resp, relevantError(err, *apiError)
}
