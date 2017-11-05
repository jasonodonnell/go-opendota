package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

const openDotaAPI = "https://api.opendota.com/api/"

// Client for making Open Dota API requests
type Client struct {
	sling          *sling.Sling
	MatchService   *MatchService
	PlayersService *PlayersService
	TeamService    *TeamService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(openDotaAPI)

	return &Client{
		sling:          base,
		MatchService:   newMatchService(base.New()),
		PlayersService: newPlayersService(base.New()),
		TeamService:    newTeamService(base.New()),
	}
}
