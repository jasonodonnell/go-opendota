package opendota

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

func newHeroService(sling *sling.Sling) *HeroService {
	return &HeroService{
		sling: sling.Path("heroes/"),
	}
}

// HeroService provides methods for accessing information
// about heroes.
type HeroService struct {
	sling *sling.Sling
}

// HeroParam is used to speficy a specific hero for queries.
type HeroParam struct {
	HeroID int64 `url:"hero_id"`
}

// Hero is a collection of information about a hero.
type Hero struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	LocalizedName string   `json:"localized_name"`
	PrimaryAttr   string   `json:"primary_attr"`
	AttackType    string   `json:"attack_type"`
	Roles         []string `json:"roles"`
	Legs          int      `json:"legs"`
}

// HeroDuration is a collection about a heroes performance
// over a range of match durations.
type HeroDuration struct {
	DurationBin int `json:"duration_bin"`
	GamesPlayed int `json:"games_played"`
	Wins        int `json:"wins"`
}

// HeroMatch is a collection of information about a hero
// in a specific match.
type HeroMatch struct {
	MatchID    int64  `json:"match_id"`
	StartTime  int    `json:"start_time"`
	Duration   int    `json:"duration"`
	RadiantWin bool   `json:"radiant_win"`
	LeagueID   int    `json:"leagueid"`
	LeagueName string `json:"league_name"`
	Radiant    bool   `json:"radiant"`
	AccountID  int    `json:"account_id"`
	Kills      int    `json:"kills"`
	Deaths     int    `json:"deaths"`
	Assists    int    `json:"assists"`
}

// HeroMatchup is a colleciton of information about how
// a hero matches up against another hero.
type HeroMatchup struct {
	HeroID      int `json:"hero_id"`
	GamesPlayed int `json:"games_played"`
	Wins        int `json:"wins"`
}

// HeroPlayer is a collection of information about players
// playing a specific hero.
type HeroPlayer struct {
	AccountID   int `json:"account_id"`
	GamesPlayed int `json:"games_played"`
	Wins        int `json:"wins"`
}

// Durations returns stats about a specific hero for varying match lengths.
func (s *HeroService) Durations(param *HeroParam) ([]HeroDuration, *http.Response, error) {
	herodurations := new([]HeroDuration)
	apiError := new(APIError)
	path := strconv.Itoa(int(param.HeroID)) + "/durations"
	resp, err := s.sling.New().Get(path).Receive(herodurations, apiError)
	return *herodurations, resp, relevantError(err, *apiError)
}

// Heroes returns a collection of all heroes.
func (s *HeroService) Heroes() ([]Hero, *http.Response, error) {
	heroes := new([]Hero)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(heroes, apiError)
	return *heroes, resp, relevantError(err, *apiError)
}

// Matches returns a collection of matches played by a specific hero.
func (s *HeroService) Matches(param *HeroParam) ([]HeroMatch, *http.Response, error) {
	heromatches := new([]HeroMatch)
	apiError := new(APIError)
	path := strconv.Itoa(int(param.HeroID)) + "/matches"
	resp, err := s.sling.New().Get(path).Receive(heromatches, apiError)
	return *heromatches, resp, relevantError(err, *apiError)
}

// Matchups returns a collection of how a hero compares against all other heroes.
func (s *HeroService) Matchups(param *HeroParam) ([]HeroMatchup, *http.Response, error) {
	heromatchups := new([]HeroMatchup)
	apiError := new(APIError)
	path := strconv.Itoa(int(param.HeroID)) + "/matchups"
	resp, err := s.sling.New().Get(path).Receive(heromatchups, apiError)
	return *heromatchups, resp, relevantError(err, *apiError)
}

// Players returns information about players who play a specific hero.
func (s *HeroService) Players(param *HeroParam) ([]HeroPlayer, *http.Response, error) {
	heroplayers := new([]HeroPlayer)
	apiError := new(APIError)
	path := strconv.Itoa(int(param.HeroID)) + "/players"
	resp, err := s.sling.New().Get(path).Receive(heroplayers, apiError)
	return *heroplayers, resp, relevantError(err, *apiError)
}
