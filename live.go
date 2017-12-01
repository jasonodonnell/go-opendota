package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newLiveService(sling *sling.Sling) *LiveService {
	return &LiveService{
		sling: sling.Path("live"),
	}
}

// LiveService provides a method for accessing live games.
type LiveService struct {
	sling *sling.Sling
}

// LiveGame represents a live game.
type LiveGame struct {
	ActivateTime   int           `json:"activate_time"`
	DeactivateTime int           `json:"deactivate_time"`
	ServerSteamID  string        `json:"server_steam_id"`
	LobbyID        string        `json:"lobby_id"`
	LeagueID       int           `json:"league_id"`
	LobbyType      int           `json:"lobby_type"`
	GameTime       int           `json:"game_time"`
	Delay          int           `json:"delay"`
	Spectators     int           `json:"spectators"`
	GameMode       int           `json:"game_mode"`
	AverageMmr     int           `json:"average_mmr"`
	SortScore      int           `json:"sort_score"`
	LastUpdateTime int           `json:"last_update_time"`
	RadiantLead    int           `json:"radiant_lead"`
	RadiantScore   int           `json:"radiant_score"`
	DireScore      int           `json:"dire_score"`
	Players        []LivePlayers `json:"players"`
	BuildingState  int           `json:"building_state"`
}

type LivePlayers struct {
	AccountID   int    `json:"account_id"`
	HeroID      int    `json:"hero_id"`
	Name        string `json:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	FantasyRole int    `json:"fantasy_role,omitempty"`
	TeamID      int    `json:"team_id,omitempty"`
	TeamName    string `json:"team_name,omitempty"`
	TeamTag     string `json:"team_tag,omitempty"`
	IsLocked    bool   `json:"is_locked,omitempty"`
	IsPro       bool   `json:"is_pro,omitempty"`
	LockedUntil int    `json:"locked_until,omitempty"`
}

// Live returns a collection of the top live games.
// https://docs.opendota.com/#tag/live%2Fpaths%2F~1live%2Fget
func (s *LiveService) Live() ([]LiveGame, *http.Response, error) {
	livegames := new([]LiveGame)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(livegames, apiError)
	return *livegames, resp, relevantError(err, *apiError)
}
