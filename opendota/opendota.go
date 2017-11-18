package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

const openDotaAPI = "https://api.opendota.com/api/"

// Client for making Open Dota API requests
type Client struct {
	sling              *sling.Sling
	ExplorerService    *ExplorerService
	MatchService       *MatchService
	PlayerService      *PlayerService
	ProMatchService    *ProMatchService
	ProPlayerService   *ProPlayerService
	PublicMatchService *PublicMatchService
	TeamService        *TeamService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(openDotaAPI)

	return &Client{
		sling:              base,
		ExplorerService:    newExplorerService(base.New()),
		MatchService:       newMatchService(base.New()),
		PlayerService:      newPlayerService(base.New()),
		ProMatchService:    newProMatchService(base.New()),
		ProPlayerService:   newProPlayerService(base.New()),
		PublicMatchService: newPublicMatchService(base.New()),
		TeamService:        newTeamService(base.New()),
	}
}
