package opendota

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

// PlayersService provides methods for accessing player
// endpoints.
type PlayersService struct {
	sling *sling.Sling
}

// PlayersParam is the parameter for specifying a player.
type PlayersParam struct {
	AccountID int64 `url:"account_id"`
}

func newPlayersService(sling *sling.Sling) *PlayersService {
	return &PlayersService{
		sling: sling.Path("Players/"),
	}
}

type Player struct {
	TrackedUntil        interface{} `json:"tracked_until"`
	SoloCompetitiveRank interface{} `json:"solo_competitive_rank"`
	MmrEstimate         MmrEstimate `json:"mmr_estimate"`
	Profile             Profile     `json:"profile"`
	CompetitiveRank     interface{} `json:"competitive_rank"`
}

type MmrEstimate struct {
	Estimate int `json:"estimate"`
}

type Profile struct {
	AccountID      int         `json:"account_id"`
	Personaname    string      `json:"personaname"`
	Name           string      `json:"name"`
	Cheese         int         `json:"cheese"`
	Steamid        string      `json:"steamid"`
	Avatar         string      `json:"avatar"`
	Avatarmedium   string      `json:"avatarmedium"`
	Avatarfull     string      `json:"avatarfull"`
	Profileurl     string      `json:"profileurl"`
	LastLogin      interface{} `json:"last_login"`
	Loccountrycode string      `json:"loccountrycode"`
}

// Player returns a specific player info.
func (s *PlayersService) Player(params *PlayersParam) (Player, *http.Response, error) {
	player := new(Player)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(int(params.AccountID))).Receive(player, apiError)
	return *player, resp, relevantError(err, *apiError)
}
