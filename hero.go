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

// Durations returns a collection of stats about a specific hero for varying match lengths.
// https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes~1%7Bhero_id%7D~1durations%2Fget
func (s *HeroService) Durations(heroID int) ([]HeroDuration, *http.Response, error) {
	herodurations := new([]HeroDuration)
	apiError := new(APIError)
	path := strconv.Itoa(heroID) + "/durations"
	resp, err := s.sling.New().Get(path).Receive(herodurations, apiError)
	return *herodurations, resp, relevantError(err, *apiError)
}

// Heroes returns a collection of all heroes.
// https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes%2Fget
func (s *HeroService) Heroes() ([]Hero, *http.Response, error) {
	heroes := new([]Hero)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(heroes, apiError)
	return *heroes, resp, relevantError(err, *apiError)
}

// Matches returns a collection of matches played by a specific hero.
// https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes~1%7Bhero_id%7D~1matches%2Fget
func (s *HeroService) Matches(heroID int) ([]HeroMatch, *http.Response, error) {
	heromatches := new([]HeroMatch)
	apiError := new(APIError)
	path := strconv.Itoa(heroID) + "/matches"
	resp, err := s.sling.New().Get(path).Receive(heromatches, apiError)
	return *heromatches, resp, relevantError(err, *apiError)
}

// Matchups returns a collection of how a hero compares against all other heroes.
// https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes~1%7Bhero_id%7D~1matchups%2Fget
func (s *HeroService) Matchups(heroID int) ([]HeroMatchup, *http.Response, error) {
	heromatchups := new([]HeroMatchup)
	apiError := new(APIError)
	path := strconv.Itoa(heroID) + "/matchups"
	resp, err := s.sling.New().Get(path).Receive(heromatchups, apiError)
	return *heromatchups, resp, relevantError(err, *apiError)
}

// Players returns a collection about players for a specific hero.
// https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes~1%7Bhero_id%7D~1players%2Fget
func (s *HeroService) Players(heroID int) ([]HeroPlayer, *http.Response, error) {
	heroplayers := new([]HeroPlayer)
	apiError := new(APIError)
	path := strconv.Itoa(heroID) + "/players"
	resp, err := s.sling.New().Get(path).Receive(heroplayers, apiError)
	return *heroplayers, resp, relevantError(err, *apiError)
}
