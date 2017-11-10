package opendota

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

func newPlayersService(sling *sling.Sling) *PlayersService {
	return &PlayersService{
		sling: sling.Path("players/"),
	}
}

// PlayersService provides methods for accessing player
// endpoints.
type PlayersService struct {
	sling *sling.Sling
}

// PlayersParam is the parameter for specifying a player.
type PlayersParam struct {
	AccountID     int64    `url:"account_id"`
	Limit         int      `url:"limit,omitempty"`
	Offset        int      `url:"offset,omitempty"`
	Win           int      `url:"win,omitempty"`
	Patch         int      `url:"patch,omitempty"`
	GameMode      int      `url:"game_mode,omitempty"`
	LobbyType     int      `url:"lobby_type,omitempty"`
	Region        int      `url:"region,omitempty"`
	Date          int      `url:"date,omitempty"`
	LaneRole      int      `url:"lane_role,omitempty"`
	HeroID        int      `url:"hero_id,omitempty"`
	IsRadiant     int      `url:"is_radiant,omitempty"`
	IncAccountID  int64    `url:"included_account_id,omitempty"`
	ExcAccountID  int64    `url:"excluded_account_id,omitempty"`
	WithHeroID    int      `url:"with_hero_id,omitempty"`
	AgainstHeroID int      `url:"against_hero_id,omitempty"`
	Significant   int      `url:"significant,omitempty"`
	Sort          string   `url:"sort,omitempty"`
	Project       []string `url:"project,omitempty"`
}

// PlayerMatch is a collection about a match for a specific player.
type PlayerMatch struct {
	MatchID      int64 `json:"match_id"`
	PlayerSlot   int   `json:"player_slot"`
	RadiantWin   bool  `json:"radiant_win"`
	Duration     int   `json:"duration"`
	GameMode     int   `json:"game_mode"`
	LobbyType    int   `json:"lobby_type"`
	HeroID       int   `json:"hero_id"`
	StartTime    int   `json:"start_time"`
	Version      int   `json:"version"`
	Kills        int   `json:"kills"`
	Deaths       int   `json:"deaths"`
	Assists      int   `json:"assists"`
	Skill        int   `json:"skill,omitempty"`
	XpPerMin     int   `json:"xp_per_min,omitempty"`
	GoldPerMin   int   `json:"gold_per_min,omitempty"`
	HeroDamage   int   `json:"hero_damage,omitempty"`
	TowerDamage  int   `json:"tower_damage,omitempty"`
	HeroHealing  int   `json:"hero_healing,omitempty"`
	LastHits     int   `json:"last_hits,omitempty"`
	Lane         int   `json:"lane,omitempty"`
	LaneRole     int   `json:"lane_role,omitempty"`
	IsRoaming    bool  `json:"is_roaming,omitempty"`
	Cluster      int   `json:"cluster,omitempty"`
	LeaverStatus int   `json:"leaver_status,omitempty"`
	PartySize    int   `json:"party_size"`
}

// Player is a collection of stats about a specific player.
type Player struct {
	TrackedUntil        string      `json:"tracked_until,omitempty"`
	SoloCompetitiveRank string      `json:"solo_competitive_rank,omitempty"`
	MmrEstimate         MmrEstimate `json:"mmr_estimate"`
	Profile             Profile     `json:"profile"`
	CompetitiveRank     string      `json:"competitive_rank,omitempty"`
}

// MmrEstimate is an estimate MMR score for a player.
type MmrEstimate struct {
	Estimate int `json:"estimate"`
}

// Profile is a collection of account information about a player.
type Profile struct {
	AccountID      int    `json:"account_id"`
	Personaname    string `json:"personaname"`
	Name           string `json:"name"`
	Cheese         int    `json:"cheese"`
	Steamid        string `json:"steamid"`
	Avatar         string `json:"avatar"`
	Avatarmedium   string `json:"avatarmedium"`
	Avatarfull     string `json:"avatarfull"`
	Profileurl     string `json:"profileurl"`
	LastLogin      string `json:"last_login,omitempty"`
	Loccountrycode string `json:"loccountrycode"`
}

// WinLoss is a collection of wins and loses for a player.
type WinLoss struct {
	Win  int `json:"win"`
	Lose int `json:"lose"`
}

// PlayerHero is a collection about heroes played for a specific player.
type PlayerHero struct {
	HeroID       string `json:"hero_id"`
	LastPlayed   int    `json:"last_played"`
	Games        int    `json:"games"`
	Win          int    `json:"win"`
	WithGames    int    `json:"with_games"`
	WithWin      int    `json:"with_win"`
	AgainstGames int    `json:"against_games"`
	AgainstWin   int    `json:"against_win"`
}

type PlayerPeers struct {
	AccountID    int    `json:"account_id"`
	LastPlayed   int    `json:"last_played"`
	Win          int    `json:"win"`
	Games        int    `json:"games"`
	WithWin      int    `json:"with_win"`
	WithGames    int    `json:"with_games"`
	AgainstWin   int    `json:"against_win"`
	AgainstGames int    `json:"against_games"`
	WithGpmSum   int    `json:"with_gpm_sum"`
	WithXpmSum   int    `json:"with_xpm_sum"`
	Personaname  string `json:"personaname"`
	LastLogin    string `json:"last_login"`
	Avatar       string `json:"avatar"`
	Avatarfull   string `json:"avatarfull"`
}

// Player returns information about a specific player.
func (s *PlayersService) Player(params *PlayersParam) (Player, *http.Response, error) {
	player := new(Player)
	apiError := new(APIError)
	path := fmt.Sprintf("%s", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).Receive(player, apiError)
	return *player, resp, relevantError(err, *apiError)
}

// WinLoss returns the win/loss count for a specific player.
func (s *PlayersService) WinLoss(params *PlayersParam) (WinLoss, *http.Response, error) {
	winloss := new(WinLoss)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/wl", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(winloss, apiError)
	return *winloss, resp, relevantError(err, *apiError)
}

// RecentMatches returns recent matches played by a specific player.
func (s *PlayersService) RecentMatches(params *PlayersParam) ([]PlayerMatch, *http.Response, error) {
	playermatches := new([]PlayerMatch)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/recentMatches", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).Receive(playermatches, apiError)
	return *playermatches, resp, relevantError(err, *apiError)
}

// Matches returns recent matches played by a specific player, can be
// queried to tune results.
func (s *PlayersService) Matches(params *PlayersParam) ([]PlayerMatch, *http.Response, error) {
	playermatches := new([]PlayerMatch)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/matches", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(playermatches, apiError)
	return *playermatches, resp, relevantError(err, *apiError)
}

// Heroes returns information about heroes played for a specific player.
func (s *PlayersService) Heroes(params *PlayersParam) ([]PlayerHero, *http.Response, error) {
	playerheroes := new([]PlayerHero)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/heroes", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(playerheroes, apiError)
	return *playerheroes, resp, relevantError(err, *apiError)
}

// Peers returns information about games played with other players.
func (s *PlayersService) Peers(params *PlayersParam) ([]PlayerPeers, *http.Response, error) {
	peers := new([]PlayerPeers)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/peers", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(peers, apiError)
	return *peers, resp, relevantError(err, *apiError)
}
