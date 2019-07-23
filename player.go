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

// PlayerService provides methods for accessing information about
// players.
type PlayerService struct {
	sling *sling.Sling
}

// PlayerParam is used for customizing Player queries.
type PlayerParam struct {
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

// GameWins represents how many games are won for a player.
type GameWins struct {
	Games int `json:"games"`
	Win   int `json:"win"`
}

// MmrEstimate represents the estimated MMR score for a player.
type MmrEstimate struct {
	Estimate int `json:"estimate"`
}

// Player represents stats about a player.
type Player struct {
	TrackedUntil        string      `json:"tracked_until,omitempty"`
	SoloCompetitiveRank int         `json:"solo_competitive_rank,omitempty"`
	MmrEstimate         MmrEstimate `json:"mmr_estimate"`
	Profile             Profile     `json:"profile"`
	CompetitiveRank     int         `json:"competitive_rank,omitempty"`
	RankTier            int         `json:"rank_tier"`
}

// PlayerCounts represents the counts of wins for a player for various
// stats.
type PlayerCounts struct {
	LeaverStatus map[string]GameWins `json:"leaver_status"`
	GameMode     map[string]GameWins `json:"game_mode"`
	LobbyType    map[string]GameWins `json:"lobby_type"`
	LaneRole     map[string]GameWins `json:"lane_role"`
	Region       map[string]GameWins `json:"region"`
	Patch        map[string]GameWins `json:"patch"`
	IsRadiant    map[string]GameWins `json:"is_radiant"`
}

// PlayerHero represents the stats of a hero played by a player.
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

// PlayerHistogram represents a distribution of data for a player.
type PlayerHistogram struct {
	X     int `json:"x"`
	Games int `json:"games"`
	Win   int `json:"win"`
}

// PlayerMatch represents match data for a player.
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

// PlayerPeers represents data about peers a player has played with.
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

// PlayerPros represents data about pro players a player has played with.
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

// PlayerRankings represents the ranking of a player.
type PlayerRankings struct {
	HeroID      int     `json:"hero_id"`
	Score       float64 `json:"score"`
	PercentRank int     `json:"percent_rank"`
	Card        int     `json:"card"`
}

// PlayerRatings represents the ratings of a player.
type PlayerRatings struct {
	AccountID           int    `json:"account_id"`
	MatchID             int64  `json:"match_id"`
	SoloCompetitiveRank int    `json:"solo_competitive_rank"`
	CompetitiveRank     int    `json:"competitive_rank"`
	Time                string `json:"time"`
}

// PlayerTotals represents totals in different fields for a player.
type PlayerTotals struct {
	Field string `json:"field"`
	N     int    `json:"n"`
	Sum   int    `json:"sum"`
}

// PlayerWardMap represents observer and sentry wards placed by a player.
type PlayerWardMap struct {
	Obs map[string]map[string]int `json:"obs"`
	Sen map[string]map[string]int `json:"sen"`
}

// PlayerWordCloud represents the words said by a player in chat.
type PlayerWordCloud struct {
	MyWordCounts  map[string]int `json:"my_word_counts"`
	AllWordCounts map[string]int `json:"all_word_counts"`
}

// Profile represents a player's profile.
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

// WinLoss represents the totals of wins and losses for a player.
type WinLoss struct {
	Win  int `json:"win"`
	Lose int `json:"lose"`
}

// Counts takes an Account ID and optional params returns the counts
// of categories for a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1counts%2Fget
func (s *PlayerService) Counts(accountID int64, params *PlayerParam) (PlayerCounts, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	counts := new(PlayerCounts)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/counts", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(counts, apiError)
	return *counts, resp, relevantError(err, *apiError)
}

// Heroes takes an Account ID and optional params and returns information
// about heroes played by a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1heroes%2Fget
func (s *PlayerService) Heroes(accountID int64, params *PlayerParam) ([]PlayerHero, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	playerheroes := new([]PlayerHero)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/heroes", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(playerheroes, apiError)
	return *playerheroes, resp, relevantError(err, *apiError)
}

// Histograms takes an Account ID, Field and optional params and returns a
// distribution of matches of a player for that field.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1histograms~1%7Bfield%7D%2Fget
func (s *PlayerService) Histograms(accountID int64, field string, params *PlayerParam) ([]PlayerHistogram, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	histograms := new([]PlayerHistogram)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/histograms/%s", strconv.Itoa(int(accountID)), field)
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(histograms, apiError)
	return *histograms, resp, relevantError(err, *apiError)
}

// Matches takes an Account ID and optional params and returns recent matches for a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1matches%2Fget
func (s *PlayerService) Matches(accountID int64, params *PlayerParam) ([]PlayerMatch, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	playermatches := new([]PlayerMatch)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/matches", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(playermatches, apiError)
	return *playermatches, resp, relevantError(err, *apiError)
}

// Peers takes an Account ID and optional params and returns information about games
// played with other players.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1peers%2Fget
func (s *PlayerService) Peers(accountID int64, params *PlayerParam) ([]PlayerPeers, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	peers := new([]PlayerPeers)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/peers", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(peers, apiError)
	return *peers, resp, relevantError(err, *apiError)
}

// Player takes an account id and returns information about a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D%2Fget
func (s *PlayerService) Player(accountID int64) (Player, *http.Response, error) {
	player := new(Player)
	apiError := new(APIError)
	path := fmt.Sprintf("%s", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).Receive(player, apiError)
	return *player, resp, relevantError(err, *apiError)
}

// Pros takes an Account ID and optional params and returns information about
// games played with other pro players.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1pros%2Fget
func (s *PlayerService) Pros(accountID int64, params *PlayerParam) ([]PlayerPros, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	pros := new([]PlayerPros)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/pros", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(pros, apiError)
	return *pros, resp, relevantError(err, *apiError)
}

// Rankings takes an Account ID and returns ranking history for a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1rankings%2Fget
func (s *PlayerService) Rankings(accountID int64) ([]PlayerRankings, *http.Response, error) {
	rankings := new([]PlayerRankings)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/rankings", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).Receive(rankings, apiError)
	return *rankings, resp, relevantError(err, *apiError)
}

// Ratings takes an Account ID and returns rating history for a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1ratings%2Fget
func (s *PlayerService) Ratings(accountID int64) ([]PlayerRatings, *http.Response, error) {
	ratings := new([]PlayerRatings)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/ratings", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).Receive(ratings, apiError)
	return *ratings, resp, relevantError(err, *apiError)
}

// RecentMatches takes an Account ID and returns recent matches for a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1recentMatches%2Fget
func (s *PlayerService) RecentMatches(accountID int64) ([]PlayerMatch, *http.Response, error) {
	playermatches := new([]PlayerMatch)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/recentMatches", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).Receive(playermatches, apiError)
	return *playermatches, resp, relevantError(err, *apiError)
}

// Totals takes an Account ID and optional params and returns totals for a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1totals%2Fget
func (s *PlayerService) Totals(accountID int64, params *PlayerParam) ([]PlayerTotals, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	totals := new([]PlayerTotals)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/totals", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(totals, apiError)
	return *totals, resp, relevantError(err, *apiError)
}

// WardMap takes an Account ID and optional params and returns wards placed
// in matches by a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1wardmap%2Fget
func (s *PlayerService) WardMap(accountID int64, params *PlayerParam) (PlayerWardMap, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	wardmap := new(PlayerWardMap)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/wardmap", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(wardmap, apiError)
	return *wardmap, resp, relevantError(err, *apiError)
}

// WinLoss takes an Account ID and optional params and returns the
// win/loss count for a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1wl%2Fget
func (s *PlayerService) WinLoss(accountID int64, params *PlayerParam) (WinLoss, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	winloss := new(WinLoss)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/wl", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(winloss, apiError)
	return *winloss, resp, relevantError(err, *apiError)
}

// WordCloud takes an Account ID and optional params and returns
// words said in matches by a player.
// https://docs.opendota.com/#tag/players%2Fpaths%2F~1players~1%7Baccount_id%7D~1wordcloud%2Fget
func (s *PlayerService) WordCloud(accountID int64, params *PlayerParam) (PlayerWordCloud, *http.Response, error) {
	if params == nil {
		params = &PlayerParam{}
	}
	wordcloud := new(PlayerWordCloud)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/wordcloud", strconv.Itoa(int(accountID)))
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(wordcloud, apiError)
	return *wordcloud, resp, relevantError(err, *apiError)
}
