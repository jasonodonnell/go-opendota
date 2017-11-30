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

// TeamService provides methods for accessing information
// about teams in Dota 2.
type TeamService struct {
	sling *sling.Sling
}

// Team represents the stats for a team.
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

// TeamHeroes represents the hero data for a team.
type TeamHeroes struct {
	HeroID        int    `json:"hero_id"`
	LocalizedName string `json:"localized_name"`
	GamesPlayed   int    `json:"games_played"`
	Wins          int    `json:"wins"`
}

// TeamMatch represents the match data for a team.
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

// TeamPlayers represents the players who have played on a team.
type TeamPlayers struct {
	AccountID           int    `json:"account_id"`
	Name                string `json:"name"`
	GamesPlayed         int    `json:"games_played"`
	Wins                int    `json:"wins"`
	IsCurrentTeamMember bool   `json:"is_current_team_member"`
}

// Heroes takes a Team ID and returns stats for the heroes played by a team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D~1heroes%2Fget
func (s *TeamService) Heroes(teamID int64) ([]TeamHeroes, *http.Response, error) {
	heroes := new([]TeamHeroes)
	apiError := new(APIError)
	path := strconv.Itoa(int(teamID)) + "/heroes"
	resp, err := s.sling.New().Get(path).Receive(heroes, apiError)
	return *heroes, resp, relevantError(err, *apiError)
}

// Matches takes a Team ID and returns matches played by a team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D~1matches%2Fget
func (s *TeamService) Matches(teamID int64) ([]TeamMatch, *http.Response, error) {
	matches := new([]TeamMatch)
	apiError := new(APIError)
	path := strconv.Itoa(int(teamID)) + "/matches"
	resp, err := s.sling.New().Get(path).Receive(matches, apiError)
	return *matches, resp, relevantError(err, *apiError)
}

// Players takes a Team ID and returns the players that played on a team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D~1players%2Fget
func (s *TeamService) Players(teamID int64) ([]TeamPlayers, *http.Response, error) {
	players := new([]TeamPlayers)
	apiError := new(APIError)
	path := strconv.Itoa(int(teamID)) + "/players"
	resp, err := s.sling.New().Get(path).Receive(players, apiError)
	return *players, resp, relevantError(err, *apiError)
}

// Team takes a Team ID and returns data about that team.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams~1%7Bteam_id%7D%2Fget
func (s *TeamService) Team(teamID int64) (Team, *http.Response, error) {
	team := new(Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(int(teamID))).Receive(team, apiError)
	return *team, resp, relevantError(err, *apiError)
}

// Teams returns data about all teams.
// https://docs.opendota.com/#tag/teams%2Fpaths%2F~1teams%2Fget
func (s *TeamService) Teams() ([]Team, *http.Response, error) {
	teams := new([]Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(teams, apiError)
	return *teams, resp, relevantError(err, *apiError)
}
