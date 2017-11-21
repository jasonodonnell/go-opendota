package opendota

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

func newPlayerService(sling *sling.Sling) *PlayerService {
	return &PlayerService{
		sling: sling.Path("players/"),
	}
}

// PlayerService provides methods for accessing player
// endpoints.
type PlayerService struct {
	sling *sling.Sling
}

// PlayersParam is the parameter for specifying a player.
type PlayersParam struct {
	AccountID     int64    `url:"account_id"`
	Field         string   `url:"field,omitempty"`
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

// GameWins is a collection of how many games are won for
// various stats.
type GameWins struct {
	Games int `json:"games"`
	Win   int `json:"win"`
}

// MmrEstimate is an estimate MMR score for a player.
type MmrEstimate struct {
	Estimate int `json:"estimate"`
}

// Player is a collection of stats about a specific player.
type Player struct {
	TrackedUntil        string      `json:"tracked_until,omitempty"`
	SoloCompetitiveRank string      `json:"solo_competitive_rank,omitempty"`
	MmrEstimate         MmrEstimate `json:"mmr_estimate"`
	Profile             Profile     `json:"profile"`
	CompetitiveRank     string      `json:"competitive_rank,omitempty"`
}

// PlayerCounts is a collection of counts of a specific player
// for various stats.
type PlayerCounts struct {
	LeaverStatus map[string]GameWins `json:"leaver_status"`
	GameMode     map[string]GameWins `json:"game_mode"`
	LobbyType    map[string]GameWins `json:"lobby_type"`
	LaneRole     map[string]GameWins `json:"lane_role"`
	Region       map[string]GameWins `json:"region"`
	Patch        map[string]GameWins `json:"patch"`
	IsRadiant    map[string]GameWins `json:"is_radiant"`
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

// PlayerHistogram is a collection that represents a distribution of
// data for a specifc player.
type PlayerHistogram struct {
	X     int `json:"x"`
	Games int `json:"games"`
	Win   int `json:"win"`
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

// PlayerPeers is a collection about peers that have played with specific player.
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
	AvatarFull   string `json:"avatarfull"`
}

// PlayerPros is a collection about pros that have played
// with a specific player.
type PlayerPros struct {
	AccountID       int    `json:"account_id"`
	Name            string `json:"name"`
	CountryCode     string `json:"country_code"`
	FantasyRole     int    `json:"fantasy_role"`
	TeamID          int    `json:"team_id"`
	TeamName        string `json:"team_name"`
	TeamTag         string `json:"team_tag"`
	IsLocked        bool   `json:"is_locked"`
	IsPro           bool   `json:"is_pro"`
	LockedUntil     int    `json:"locked_until"`
	SteamID         string `json:"steamid"`
	Avatar          string `json:"avatar"`
	AvatarMedium    string `json:"avatarmedium"`
	AvatarFull      string `json:"avatarfull"`
	ProfileURL      string `json:"profileurl"`
	Personaname     string `json:"personaname"`
	Cheese          int    `json:"cheese"`
	FhUnavailable   bool   `json:"fh_unavailable"`
	LocCountryCode  string `json:"loccountrycode"`
	LastPlayed      int    `json:"last_played"`
	FullHistoryTime string `json:"full_history_time"`
	LastMatchTime   string `json:"last_match_time"`
	Win             int    `json:"win"`
	Games           int    `json:"games"`
	WithWin         int    `json:"with_win"`
	WithGames       int    `json:"with_games"`
	AgainstWin      int    `json:"against_win"`
	AgainstGames    int    `json:"against_games"`
	WithGpmSum      int    `json:"with_gpm_sum"`
	WithXpmSum      int    `json:"with_xpm_sum"`
}

// PlayerRankings is a collection of rankings for a specific player.
type PlayerRankings struct {
	HeroID      int     `json:"hero_id"`
	Score       float64 `json:"score"`
	PercentRank int     `json:"percent_rank"`
	Card        int     `json:"card"`
}

// PlayerRatings is a collection of ratings over time for a specific player.
type PlayerRatings struct {
	AccountID           int    `json:"account_id"`
	MatchID             int64  `json:"match_id"`
	SoloCompetitiveRank int    `json:"solo_competitive_rank"`
	CompetitiveRank     int    `json:"competitive_rank"`
	Time                string `json:"time"`
}

// PlayerTotals is a collection of stats about a specific player
// for different fields.
type PlayerTotals struct {
	Field string `json:"field"`
	N     int    `json:"n"`
	Sum   int    `json:"sum"`
}

// PlayerWardMap is a collection of observer and sentry wards placed
// by a specific player.
type PlayerWardMap struct {
	Obs map[string]map[string]int `json:"obs"`
	Sen map[string]map[string]int `json:"sen"`
}

// PlayerWordCloud is a collection of words said by a specific player.
type PlayerWordCloud struct {
	MyWordCounts  map[string]int `json:"my_word_counts"`
	AllWordCounts map[string]int `json:"all_word_counts"`
}

// Profile is a collection of account information about a player.
type Profile struct {
	AccountID      int    `json:"account_id"`
	Personaname    string `json:"personaname"`
	Name           string `json:"name"`
	Cheese         int    `json:"cheese"`
	SteamID        string `json:"steamid"`
	Avatar         string `json:"avatar"`
	AvatarMedium   string `json:"avatarmedium"`
	AvatarFull     string `json:"avatarfull"`
	ProfileURL     string `json:"profileurl"`
	LastLogin      string `json:"last_login,omitempty"`
	LocCountryCode string `json:"loccountrycode"`
}

// WinLoss is a collection of wins and loses for a player.
type WinLoss struct {
	Win  int `json:"win"`
	Lose int `json:"lose"`
}

// Counts returns the count of categories for a specific player.
func (s *PlayerService) Counts(params *PlayersParam) (PlayerCounts, *http.Response, error) {
	counts := new(PlayerCounts)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/counts", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(counts, apiError)
	return *counts, resp, relevantError(err, *apiError)
}

// Heroes returns information about heroes played for a specific player.
func (s *PlayerService) Heroes(params *PlayersParam) ([]PlayerHero, *http.Response, error) {
	playerheroes := new([]PlayerHero)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/heroes", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(playerheroes, apiError)
	return *playerheroes, resp, relevantError(err, *apiError)
}

// Histograms returns a distribution of matches in a single field for a specific
// player.
func (s *PlayerService) Histograms(params *PlayersParam) ([]PlayerHistogram, *http.Response, error) {
	histograms := new([]PlayerHistogram)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/histograms/%s", strconv.Itoa(int(params.AccountID)), params.Field)
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(histograms, apiError)
	return *histograms, resp, relevantError(err, *apiError)
}

// Matches returns recent matches played by a specific player, can be
// queried to tune results.
func (s *PlayerService) Matches(params *PlayersParam) ([]PlayerMatch, *http.Response, error) {
	playermatches := new([]PlayerMatch)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/matches", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(playermatches, apiError)
	return *playermatches, resp, relevantError(err, *apiError)
}

// Peers returns information about games played with other players.
func (s *PlayerService) Peers(params *PlayersParam) ([]PlayerPeers, *http.Response, error) {
	peers := new([]PlayerPeers)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/peers", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(peers, apiError)
	return *peers, resp, relevantError(err, *apiError)
}

// Player returns information about a specific player.
func (s *PlayerService) Player(params *PlayersParam) (Player, *http.Response, error) {
	player := new(Player)
	apiError := new(APIError)
	path := fmt.Sprintf("%s", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).Receive(player, apiError)
	return *player, resp, relevantError(err, *apiError)
}

// Pros returns information about games played with other pro players.
func (s *PlayerService) Pros(params *PlayersParam) ([]PlayerPros, *http.Response, error) {
	pros := new([]PlayerPros)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/pros", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(pros, apiError)
	return *pros, resp, relevantError(err, *apiError)
}

// Rankings returns ranking history for a specific player.
func (s *PlayerService) Rankings(params *PlayersParam) ([]PlayerRankings, *http.Response, error) {
	rankings := new([]PlayerRankings)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/rankings", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).Receive(rankings, apiError)
	return *rankings, resp, relevantError(err, *apiError)
}

// Ratings returns rating history for a specific player.
func (s *PlayerService) Ratings(params *PlayersParam) ([]PlayerRatings, *http.Response, error) {
	ratings := new([]PlayerRatings)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/ratings", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).Receive(ratings, apiError)
	return *ratings, resp, relevantError(err, *apiError)
}

// RecentMatches returns recent matches played by a specific player.
func (s *PlayerService) RecentMatches(params *PlayersParam) ([]PlayerMatch, *http.Response, error) {
	playermatches := new([]PlayerMatch)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/recentMatches", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).Receive(playermatches, apiError)
	return *playermatches, resp, relevantError(err, *apiError)
}

// Totals returns the total in stats for a specific player.
func (s *PlayerService) Totals(params *PlayersParam) ([]PlayerTotals, *http.Response, error) {
	totals := new([]PlayerTotals)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/totals", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(totals, apiError)
	return *totals, resp, relevantError(err, *apiError)
}

// WardMap returns wards placed in matches by a specific player.
func (s *PlayerService) WardMap(params *PlayersParam) (PlayerWardMap, *http.Response, error) {
	wardmap := new(PlayerWardMap)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/wardmap", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(wardmap, apiError)
	return *wardmap, resp, relevantError(err, *apiError)
}

// WinLoss returns the win/loss count for a specific player.
func (s *PlayerService) WinLoss(params *PlayersParam) (WinLoss, *http.Response, error) {
	winloss := new(WinLoss)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/wl", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(winloss, apiError)
	return *winloss, resp, relevantError(err, *apiError)
}

// WordCloud returns words said/read in matches by a player.
func (s *PlayerService) WordCloud(params *PlayersParam) (PlayerWordCloud, *http.Response, error) {
	wordcloud := new(PlayerWordCloud)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/wordcloud", strconv.Itoa(int(params.AccountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(wordcloud, apiError)
	return *wordcloud, resp, relevantError(err, *apiError)
}
