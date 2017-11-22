package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newProPlayerService(sling *sling.Sling) *ProPlayerService {
	return &ProPlayerService{
		sling: sling.Path("proPlayers"),
	}
}

// ProPlayerService provides methods for accessing pro
// endpoints.
type ProPlayerService struct {
	sling *sling.Sling
}

// ProPlayer is a collectiong about a specific professional
// dota player.
type ProPlayer struct {
	AccountID       int    `json:"account_id"`
	SteamID         string `json:"steamid"`
	Avatar          string `json:"avatar"`
	AvatarMedium    string `json:"avatarmedium"`
	AvatarFull      string `json:"avatarfull"`
	ProfileURL      string `json:"profileurl"`
	Personaname     string `json:"personaname"`
	LastLogin       string `json:"last_login"`
	FullHistoryTime string `json:"full_history_time"`
	Cheese          int    `json:"cheese"`
	FhUnavailable   bool   `json:"fh_unavailable"`
	LocCountryCode  string `json:"loccountrycode"`
	LastMatchTime   string `json:"last_match_time"`
	Name            string `json:"name"`
	CountryCode     string `json:"country_code"`
	FantasyRole     int    `json:"fantasy_role"`
	TeamID          int    `json:"team_id"`
	TeamName        string `json:"team_name"`
	TeamTag         string `json:"team_tag"`
	IsLocked        bool   `json:"is_locked"`
	IsPro           bool   `json:"is_pro"`
	LockedUntil     int    `json:"locked_until"`
}

// Players returns information about pro players.
func (s *ProPlayerService) Players() ([]ProPlayer, *http.Response, error) {
	proplayers := new([]ProPlayer)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(proplayers, apiError)
	return *proplayers, resp, relevantError(err, *apiError)
}
