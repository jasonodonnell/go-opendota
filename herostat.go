package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newHeroStatService(sling *sling.Sling) *HeroStatService {
	return &HeroStatService{
		sling: sling.Path("heroStats"),
	}
}

// HeroStatService provides a method for accessing hero stats.
type HeroStatService struct {
	sling *sling.Sling
}

// HeroStat represents a collection of statistics about a hero.
type HeroStat struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	LocalizedName   string   `json:"localized_name"`
	PrimaryAttr     string   `json:"primary_attr"`
	AttackType      string   `json:"attack_type"`
	Roles           []string `json:"roles"`
	Img             string   `json:"img"`
	Icon            string   `json:"icon"`
	BaseHealth      int      `json:"base_health"`
	BaseHealthRegen float64  `json:"base_health_regen"`
	BaseMana        int      `json:"base_mana"`
	BaseManaRegen   float64  `json:"base_mana_regen"`
	BaseArmor       float64  `json:"base_armor"`
	BaseMr          int      `json:"base_mr"`
	BaseAttackMin   int      `json:"base_attack_min"`
	BaseAttackMax   int      `json:"base_attack_max"`
	BaseStr         int      `json:"base_str"`
	BaseAgi         int      `json:"base_agi"`
	BaseInt         int      `json:"base_int"`
	StrGain         float64  `json:"str_gain"`
	AgiGain         float64  `json:"agi_gain"`
	IntGain         float64  `json:"int_gain"`
	AttackRange     int      `json:"attack_range"`
	ProjectileSpeed int      `json:"projectile_speed"`
	AttackRate      float64  `json:"attack_rate"`
	MoveSpeed       int      `json:"move_speed"`
	TurnRate        float64  `json:"turn_rate"`
	CmEnabled       bool     `json:"cm_enabled"`
	Legs            int      `json:"legs"`
	ProWin          int      `json:"pro_win,omitempty"`
	ProPick         int      `json:"pro_pick,omitempty"`
	HeroID          int      `json:"hero_id,omitempty"`
	ProBan          int      `json:"pro_ban,omitempty"`
	NullPick        int      `json:"null_pick"`
	NullWin         int      `json:"null_win"`
	OnePick         int      `json:"1_pick"`
	OneWin          int      `json:"1_win"`
	TwoPick         int      `json:"2_pick"`
	TwoWin          int      `json:"2_win"`
	ThreePick       int      `json:"3_pick"`
	ThreeWin        int      `json:"3_win"`
	FourPick        int      `json:"4_pick"`
	FourWin         int      `json:"4_win"`
	FivePick        int      `json:"5_pick"`
	FiveWin         int      `json:"5_win"`
	SixPick         int      `json:"6_pick"`
	SixWin          int      `json:"6_win"`
	SevenPick       int      `json:"7_pick"`
	SevenWin        int      `json:"7_win"`
	EightPick       int      `json:"8_pick"`
	EightWin        int      `json:"8_win"`
}

// HeroStats returns a collection of stats about all heroes.
// https://docs.opendota.com/#tag/hero-stats%2Fpaths%2F~1heroStats%2Fget
func (s *HeroStatService) HeroStats() ([]HeroStat, *http.Response, error) {
	herostats := new([]HeroStat)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(herostats, apiError)
	return *herostats, resp, relevantError(err, *apiError)
}
