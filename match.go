package opendota

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

func newMatchService(sling *sling.Sling) *MatchService {
	return &MatchService{
		sling: sling.Path("matches/"),
	}
}

// MatchService provides methods for accessing teams
// endpoints.
type MatchService struct {
	sling *sling.Sling
}

// Match is a collection of information about a specific Dota 2
// match.
type Match struct {
	MatchID               int64          `json:"match_id"`
	BarracksStatusDire    int            `json:"barracks_status_dire"`
	BarracksStatusRadiant int            `json:"barracks_status_radiant"`
	Chat                  []chat         `json:"chat"`
	Cluster               int            `json:"cluster"`
	Cosmetics             map[string]int `json:"cosmetics"`
	DireScore             int            `json:"dire_score"`
	Duration              int            `json:"duration"`
	Engine                int            `json:"engine"`
	FirstBloodTime        int            `json:"first_blood_time"`
	GameMode              int            `json:"game_mode"`
	HumanPlayers          int            `json:"human_players"`
	LeagueID              int            `json:"leagueid"`
	LobbyType             int            `json:"lobby_type"`
	MatchSeqNum           int64          `json:"match_seq_num"`
	NegativeVotes         int            `json:"negative_votes"`
	Objectives            []objective    `json:"objectives"`
	PicksBans             []pickbans     `json:"picks_bans"`
	PositiveVotes         int            `json:"positive_votes"`
	RadiantGoldAdv        []int          `json:"radiant_gold_adv"`
	RadiantScore          int            `json:"radiant_score"`
	RadiantWin            bool           `json:"radiant_win"`
	RadiantXpAdv          []int          `json:"radiant_xp_adv"`
	Skill                 int            `json:"skill"`
	StartTime             int            `json:"start_time"`
	Teamfights            []teamfights   `json:"teamfights"`
	TowerStatusDire       int            `json:"tower_status_dire"`
	TowerStatusRadiant    int            `json:"tower_status_radiant"`
	Version               int            `json:"version"`
	ReplaySalt            int            `json:"replay_salt"`
	SeriesID              int            `json:"series_id"`
	SeriesType            int            `json:"series_type"`
	League                league         `json:"league"`
	RadiantTeam           matchteam      `json:"radiant_team"`
	DireTeam              matchteam      `json:"dire_team"`
	Players               []matchplayer  `json:"players"`
	Patch                 int            `json:"patch"`
	Region                int            `json:"region"`
	AllWordCounts         map[string]int `json:"all_word_counts"`
	MyWordCounts          map[string]int `json:"my_word_counts"`
	Throw                 int            `json:"throw"`
	Loss                  int            `json:"loss"`
	ReplayURL             string         `json:"replay_url"`
}

type benchmarks struct {
	GoldPerMin        rawpct `json:"gold_per_min"`
	XpPerMin          rawpct `json:"xp_per_min"`
	KillsPerMin       rawpct `json:"kills_per_min"`
	LastHitsPerMin    rawpct `json:"last_hits_per_min"`
	HeroDamagePerMin  rawpct `json:"hero_damage_per_min"`
	HeroHealingPerMin rawpct `json:"hero_healing_per_min"`
	TowerDamage       rawpct `json:"tower_damage"`
}

type buybacklog struct {
	Time       int    `json:"time"`
	Slot       int    `json:"slot"`
	Type       string `json:"type"`
	PlayerSlot int    `json:"player_slot"`
}

type chat struct {
	Time       int    `json:"time"`
	Type       string `json:"type"`
	Unit       string `json:"unit,omitempty"`
	Key        string `json:"key"`
	Slot       int    `json:"slot"`
	PlayerSlot int    `json:"player_slot"`
}

type cosmetics struct {
	ItemID          int    `json:"item_id"`
	Name            string `json:"name"`
	Prefab          string `json:"prefab"`
	CreationDate    string `json:"creation_date"`
	ImageInventory  string `json:"image_inventory"`
	ImagePath       string `json:"image_path"`
	ItemDescription string `json:"item_description"`
	ItemName        string `json:"item_name"`
	ItemRarity      string `json:"item_rarity"`
	ItemTypeName    string `json:"item_type_name"`
	UsedByHeroes    string `json:"used_by_heroes"`
}

type league struct {
	LeagueID int    `json:"leagueid"`
	Ticket   string `json:"ticket"`
	Banner   string `json:"banner"`
	Tier     string `json:"tier"`
	Name     string `json:"name"`
}

type log struct {
	Time int    `json:"time"`
	Key  string `json:"key"`
}

type matchplayer struct {
	MatchID                 int64                     `json:"match_id"`
	PlayerSlot              int                       `json:"player_slot"`
	AbilityUpgradesArr      []int                     `json:"ability_upgrades_arr"`
	AbilityUses             map[string]int            `json:"ability_uses"`
	AccountID               int                       `json:"account_id"`
	Actions                 map[string]int            `json:"actions"`
	Assists                 int                       `json:"assists"`
	Backpack0               int                       `json:"backpack_0"`
	Backpack1               int                       `json:"backpack_1"`
	Backpack2               int                       `json:"backpack_2"`
	BuybackLog              []buybacklog              `json:"buyback_log"`
	CampsStacked            int                       `json:"camps_stacked"`
	CreepsStacked           int                       `json:"creeps_stacked"`
	Damage                  map[string]int            `json:"damage"`
	DamageInflictor         map[string]int            `json:"damage_inflictor"`
	DamageInflictorReceived map[string]int            `json:"damage_inflictor_received"`
	DamageTaken             map[string]int            `json:"damage_taken"`
	Deaths                  int                       `json:"deaths"`
	Denies                  int                       `json:"denies"`
	DnT                     []int                     `json:"dn_t"`
	FirstbloodClaimed       int                       `json:"firstblood_claimed"`
	Gold                    int                       `json:"gold"`
	GoldPerMin              int                       `json:"gold_per_min"`
	GoldReasons             map[string]int            `json:"gold_reasons"`
	GoldSpent               int                       `json:"gold_spent"`
	GoldT                   []int                     `json:"gold_t"`
	HeroDamage              int                       `json:"hero_damage"`
	HeroHealing             int                       `json:"hero_healing"`
	HeroHits                map[string]int            `json:"hero_hits"`
	HeroID                  int                       `json:"hero_id"`
	Item0                   int                       `json:"item_0"`
	Item1                   int                       `json:"item_1"`
	Item2                   int                       `json:"item_2"`
	Item3                   int                       `json:"item_3"`
	Item4                   int                       `json:"item_4"`
	Item5                   int                       `json:"item_5"`
	ItemUses                map[string]int            `json:"item_uses"`
	KillStreaks             map[string]int            `json:"kill_streaks"`
	Killed                  map[string]int            `json:"killed"`
	KilledBy                map[string]int            `json:"killed_by"`
	Kills                   int                       `json:"kills"`
	KillsLog                []log                     `json:"kills_log"`
	LanePos                 map[string]map[string]int `json:"lane_pos"`
	LastHits                int                       `json:"last_hits"`
	LeaverStatus            int                       `json:"leaver_status"`
	Level                   int                       `json:"level"`
	LhT                     []int                     `json:"lh_t"`
	LifeState               map[string]int            `json:"life_state"`
	MaxHeroHit              maxherohit                `json:"max_hero_hit"`
	MultiKills              map[string]int            `json:"multi_kills"`
	Obs                     map[string]map[string]int `json:"obs"`
	ObsLeftLog              []obslog                  `json:"obs_left_log"`
	ObsLog                  []obslog                  `json:"obs_log"`
	ObsPlaced               int                       `json:"obs_placed"`
	PartyID                 int                       `json:"party_id"`
	PartySize               int                       `json:"party_size"`
	Pings                   int                       `json:"pings"`
	PredVict                bool                      `json:"pred_vict"`
	Purchase                map[string]int            `json:"purchase"`
	PurchaseLog             []log                     `json:"purchase_log"`
	Randomed                bool                      `json:"randomed"`
	RoshansKilled           int                       `json:"roshans_killed"`
	RunePickups             int                       `json:"rune_pickups"`
	Runes                   map[string]int            `json:"runes"`
	RunesLog                []runeslog                `json:"runes_log"`
	Sen                     map[string]map[string]int `json:"sen"`
	SenLeftLog              []obslog                  `json:"sen_left_log"`
	SenLog                  []obslog                  `json:"sen_log"`
	SenPlaced               int                       `json:"sen_placed"`
	Stuns                   float64                   `json:"stuns"`
	TeamfightParticipation  float64                   `json:"teamfight_participation"`
	Times                   []int                     `json:"times"`
	TowerDamage             int                       `json:"tower_damage"`
	TowersKilled            int                       `json:"towers_killed"`
	XpPerMin                int                       `json:"xp_per_min"`
	XpReasons               map[string]int            `json:"xp_reasons"`
	XpT                     []int                     `json:"xp_t"`
	Personaname             string                    `json:"personaname"`
	Name                    string                    `json:"name"`
	LastLogin               string                    `json:"last_login"`
	RadiantWin              bool                      `json:"radiant_win"`
	StartTime               int                       `json:"start_time"`
	Duration                int                       `json:"duration"`
	Cluster                 int                       `json:"cluster"`
	LobbyType               int                       `json:"lobby_type"`
	GameMode                int                       `json:"game_mode"`
	Patch                   int                       `json:"patch"`
	Region                  int                       `json:"region"`
	IsRadiant               bool                      `json:"isRadiant"`
	Win                     int                       `json:"win"`
	Lose                    int                       `json:"lose"`
	TotalGold               int                       `json:"total_gold"`
	TotalXp                 int                       `json:"total_xp"`
	KillsPerMin             float64                   `json:"kills_per_min"`
	Kda                     int                       `json:"kda"`
	Abandons                int                       `json:"abandons"`
	NeutralKills            int                       `json:"neutral_kills"`
	TowerKills              int                       `json:"tower_kills"`
	CourierKills            int                       `json:"courier_kills"`
	LaneKills               int                       `json:"lane_kills"`
	HeroKills               int                       `json:"hero_kills"`
	ObserverKills           int                       `json:"observer_kills"`
	SentryKills             int                       `json:"sentry_kills"`
	RoshanKills             int                       `json:"roshan_kills"`
	NecronomiconKills       int                       `json:"necronomicon_kills"`
	AncientKills            int                       `json:"ancient_kills"`
	BuybackCount            int                       `json:"buyback_count"`
	ObserverUses            int                       `json:"observer_uses"`
	SentryUses              int                       `json:"sentry_uses"`
	LaneEfficiency          float64                   `json:"lane_efficiency"`
	LaneEfficiencyPct       int                       `json:"lane_efficiency_pct"`
	Lane                    int                       `json:"lane"`
	LaneRole                int                       `json:"lane_role"`
	IsRoaming               bool                      `json:"is_roaming"`
	PurchaseTime            map[string]int            `json:"purchase_time"`
	FirstPurchaseTime       map[string]int            `json:"first_purchase_time"`
	ItemWin                 map[string]int            `json:"item_win"`
	ItemUsage               map[string]int            `json:"item_usage"`
	PurchaseTpscroll        int                       `json:"purchase_tpscroll"`
	ActionsPerMin           int                       `json:"actions_per_min"`
	LifeStateDead           int                       `json:"life_state_dead"`
	SoloCompetitiveRank     int                       `json:"solo_competitive_rank"`
	Cosmetics               []cosmetics               `json:"cosmetics"`
	Benchmarks              benchmarks                `json:"benchmarks"`
	PurchaseWardObserver    int                       `json:"purchase_ward_observer,omitempty"`
	PurchaseWardSentry      int                       `json:"purchase_ward_sentry,omitempty"`
	PurchaseGem             int                       `json:"purchase_gem,omitempty"`
}

type matchteam struct {
	TeamID  int    `json:"team_id"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	LogoURL string `json:"logo_url"`
}

type maxherohit struct {
	Type       string `json:"type"`
	Time       int    `json:"time"`
	Max        bool   `json:"max"`
	Inflictor  string `json:"inflictor"`
	Unit       string `json:"unit"`
	Key        string `json:"key"`
	Value      int    `json:"value"`
	Slot       int    `json:"slot"`
	PlayerSlot int    `json:"player_slot"`
}

type objective struct {
	Time       int    `json:"time"`
	Type       string `json:"type"`
	Team       int    `json:"team,omitempty"`
	Slot       int    `json:"slot,omitempty"`
	PlayerSlot int    `json:"player_slot,omitempty"`
	Unit       string `json:"unit,omitempty"`
}

type obslog struct {
	Time       int    `json:"time"`
	Type       string `json:"type"`
	Key        string `json:"key"`
	Slot       int    `json:"slot"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Z          int    `json:"z"`
	Entityleft bool   `json:"entityleft"`
	Ehandle    int    `json:"ehandle"`
	PlayerSlot int    `json:"player_slot"`
}

type pickbans struct {
	IsPick  bool  `json:"is_pick"`
	HeroID  int   `json:"hero_id"`
	Team    int   `json:"team"`
	Order   int   `json:"order"`
	Ord     int   `json:"ord"`
	MatchID int64 `json:"match_id"`
}

type rawpct struct {
	Raw float64 `json:"raw"`
	Pct float64 `json:"pct"`
}

type runeslog struct {
	Time int `json:"time"`
	Key  int `json:"key"`
}

type teamfightplayers struct {
	AbilityUses map[string]int `json:"ability_uses"`
	ItemUses    map[string]int `json:"item_uses"`
	Killed      map[string]int `json:"killed"`
	Deaths      int            `json:"deaths"`
	Buybacks    int            `json:"buybacks"`
	Damage      int            `json:"damage"`
	Healing     int            `json:"healing"`
	GoldDelta   int            `json:"gold_delta"`
	XpDelta     int            `json:"xp_delta"`
	XpStart     int            `json:"xp_start"`
	XpEnd       int            `json:"xp_end"`
}

type teamfights struct {
	Start     int                `json:"start"`
	End       int                `json:"end"`
	LastDeath int                `json:"last_death"`
	Deaths    int                `json:"deaths"`
	Players   []teamfightplayers `json:"players"`
}

// Match returns a collection for a specific match.
// https://docs.opendota.com/#tag/matches
func (s *MatchService) Match(matchID int64) (Match, *http.Response, error) {
	match := new(Match)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(int(matchID))).Receive(match, apiError)
	return *match, resp, relevantError(err, *apiError)
}
