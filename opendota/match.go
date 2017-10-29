package opendota

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

// MatchService provides methods for accessing teams
// endpoints.
type MatchService struct {
	sling *sling.Sling
}

type Match struct {
	MatchID               int `json:"match_id"`
	BarracksStatusDire    int `json:"barracks_status_dire"`
	BarracksStatusRadiant int `json:"barracks_status_radiant"`
	Chat                  []struct {
		Time       int    `json:"time"`
		Unit       string `json:"unit"`
		Key        string `json:"key"`
		Slot       int    `json:"slot"`
		PlayerSlot int    `json:"player_slot"`
	} `json:"chat"`
	Cluster   int `json:"cluster"`
	Cosmetics struct {
	} `json:"cosmetics"`
	DireScore      int `json:"dire_score"`
	Duration       int `json:"duration"`
	Engine         int `json:"engine"`
	FirstBloodTime int `json:"first_blood_time"`
	GameMode       int `json:"game_mode"`
	HumanPlayers   int `json:"human_players"`
	Leagueid       int `json:"leagueid"`
	LobbyType      int `json:"lobby_type"`
	MatchSeqNum    int `json:"match_seq_num"`
	NegativeVotes  int `json:"negative_votes"`
	Objectives     struct {
	} `json:"objectives"`
	PicksBans      interface{} `json:"picks_bans"`
	PositiveVotes  int         `json:"positive_votes"`
	RadiantGoldAdv struct {
	} `json:"radiant_gold_adv"`
	RadiantScore int  `json:"radiant_score"`
	RadiantWin   bool `json:"radiant_win"`
	RadiantXpAdv struct {
	} `json:"radiant_xp_adv"`
	StartTime  int `json:"start_time"`
	Teamfights struct {
	} `json:"teamfights"`
	TowerStatusDire    int `json:"tower_status_dire"`
	TowerStatusRadiant int `json:"tower_status_radiant"`
	Version            int `json:"version"`
	ReplaySalt         int `json:"replay_salt"`
	SeriesID           int `json:"series_id"`
	SeriesType         int `json:"series_type"`
	RadiantTeam        struct {
	} `json:"radiant_team"`
	DireTeam struct {
	} `json:"dire_team"`
	League struct {
	} `json:"league"`
	Skill   int `json:"skill"`
	Players []struct {
		MatchID            int   `json:"match_id"`
		PlayerSlot         int   `json:"player_slot"`
		AbilityUpgradesArr []int `json:"ability_upgrades_arr"`
		AbilityUses        struct {
		} `json:"ability_uses"`
		AccountID int `json:"account_id"`
		Actions   struct {
		} `json:"actions"`
		AdditionalUnits interface{} `json:"additional_units"`
		Assists         int         `json:"assists"`
		Backpack0       int         `json:"backpack_0"`
		Backpack1       int         `json:"backpack_1"`
		Backpack2       int         `json:"backpack_2"`
		BuybackLog      []struct {
			Time       int `json:"time"`
			Slot       int `json:"slot"`
			PlayerSlot int `json:"player_slot"`
		} `json:"buyback_log"`
		CampsStacked  int `json:"camps_stacked"`
		CreepsStacked int `json:"creeps_stacked"`
		Damage        struct {
		} `json:"damage"`
		DamageInflictor struct {
		} `json:"damage_inflictor"`
		DamageInflictorReceived struct {
		} `json:"damage_inflictor_received"`
		DamageTaken struct {
		} `json:"damage_taken"`
		Deaths      int   `json:"deaths"`
		Denies      int   `json:"denies"`
		DnT         []int `json:"dn_t"`
		Gold        int   `json:"gold"`
		GoldPerMin  int   `json:"gold_per_min"`
		GoldReasons struct {
		} `json:"gold_reasons"`
		GoldSpent   int   `json:"gold_spent"`
		GoldT       []int `json:"gold_t"`
		HeroDamage  int   `json:"hero_damage"`
		HeroHealing int   `json:"hero_healing"`
		HeroHits    struct {
		} `json:"hero_hits"`
		HeroID   int `json:"hero_id"`
		Item0    int `json:"item_0"`
		Item1    int `json:"item_1"`
		Item2    int `json:"item_2"`
		Item3    int `json:"item_3"`
		Item4    int `json:"item_4"`
		Item5    int `json:"item_5"`
		ItemUses struct {
		} `json:"item_uses"`
		KillStreaks struct {
		} `json:"kill_streaks"`
		Killed struct {
		} `json:"killed"`
		KilledBy struct {
		} `json:"killed_by"`
		Kills    int `json:"kills"`
		KillsLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"kills_log"`
		LanePos struct {
		} `json:"lane_pos"`
		LastHits     int   `json:"last_hits"`
		LeaverStatus int   `json:"leaver_status"`
		Level        int   `json:"level"`
		LhT          []int `json:"lh_t"`
		LifeState    struct {
		} `json:"life_state"`
		MaxHeroHit struct {
		} `json:"max_hero_hit"`
		MultiKills struct {
		} `json:"multi_kills"`
		Obs struct {
		} `json:"obs"`
		ObsLeftLog []struct {
		} `json:"obs_left_log"`
		ObsLog []struct {
		} `json:"obs_log"`
		ObsPlaced      int `json:"obs_placed"`
		PartyID        int `json:"party_id"`
		PermanentBuffs []struct {
		} `json:"permanent_buffs"`
		Pings    int `json:"pings"`
		Purchase struct {
		} `json:"purchase"`
		PurchaseLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"purchase_log"`
		RunePickups int `json:"rune_pickups"`
		Runes       struct {
			Property1 int `json:"property1"`
			Property2 int `json:"property2"`
		} `json:"runes"`
		RunesLog []struct {
			Time int `json:"time"`
			Key  int `json:"key"`
		} `json:"runes_log"`
		Sen struct {
		} `json:"sen"`
		SenLeftLog []struct {
		} `json:"sen_left_log"`
		SenLog []struct {
		} `json:"sen_log"`
		SenPlaced   int   `json:"sen_placed"`
		Stuns       int   `json:"stuns"`
		Times       []int `json:"times"`
		TowerDamage int   `json:"tower_damage"`
		XpPerMin    int   `json:"xp_per_min"`
		XpReasons   struct {
		} `json:"xp_reasons"`
		XpT               []int       `json:"xp_t"`
		Personaname       string      `json:"personaname"`
		Name              string      `json:"name"`
		LastLogin         interface{} `json:"last_login"`
		RadiantWin        bool        `json:"radiant_win"`
		StartTime         int         `json:"start_time"`
		Duration          int         `json:"duration"`
		Cluster           int         `json:"cluster"`
		LobbyType         int         `json:"lobby_type"`
		GameMode          int         `json:"game_mode"`
		Patch             int         `json:"patch"`
		Region            int         `json:"region"`
		IsRadiant         bool        `json:"isRadiant"`
		Win               int         `json:"win"`
		Lose              int         `json:"lose"`
		TotalGold         int         `json:"total_gold"`
		TotalXp           int         `json:"total_xp"`
		KillsPerMin       float64     `json:"kills_per_min"`
		Kda               int         `json:"kda"`
		Abandons          int         `json:"abandons"`
		NeutralKills      int         `json:"neutral_kills"`
		TowerKills        int         `json:"tower_kills"`
		CourierKills      int         `json:"courier_kills"`
		LaneKills         int         `json:"lane_kills"`
		HeroKills         int         `json:"hero_kills"`
		ObserverKills     int         `json:"observer_kills"`
		SentryKills       int         `json:"sentry_kills"`
		RoshanKills       int         `json:"roshan_kills"`
		NecronomiconKills int         `json:"necronomicon_kills"`
		AncientKills      int         `json:"ancient_kills"`
		BuybackCount      int         `json:"buyback_count"`
		ObserverUses      int         `json:"observer_uses"`
		SentryUses        int         `json:"sentry_uses"`
		LaneEfficiency    int         `json:"lane_efficiency"`
		LaneEfficiencyPct int         `json:"lane_efficiency_pct"`
		Lane              int         `json:"lane"`
		LaneRole          int         `json:"lane_role"`
		IsRoaming         bool        `json:"is_roaming"`
		PurchaseTime      struct {
		} `json:"purchase_time"`
		FirstPurchaseTime struct {
		} `json:"first_purchase_time"`
		ItemWin struct {
		} `json:"item_win"`
		ItemUsage struct {
		} `json:"item_usage"`
		PurchaseTpscroll struct {
		} `json:"purchase_tpscroll"`
		ActionsPerMin       int   `json:"actions_per_min"`
		LifeStateDead       int   `json:"life_state_dead"`
		SoloCompetitiveRank int   `json:"solo_competitive_rank"`
		Cosmetics           []int `json:"cosmetics"`
		Benchmarks          struct {
		} `json:"benchmarks"`
	} `json:"players"`
	Patch         int `json:"patch"`
	Region        int `json:"region"`
	AllWordCounts struct {
	} `json:"all_word_counts"`
	MyWordCounts struct {
	} `json:"my_word_counts"`
	Throw     int    `json:"throw"`
	Loss      int    `json:"loss"`
	ReplayURL string `json:"replay_url"`
}

// MatchParam is the parameter for specifying a match by ID.
type MatchParam struct {
	MatchID int64 `url:"match_id,omitempty"`
}

func newMatchService(sling *sling.Sling) *MatchService {
	return &MatchService{
		sling: sling.Path("matches/"),
	}
}

// Match returns a collection for a specific match.
func (s *MatchService) Match(params *MatchParam) (Match, *http.Response, error) {
	match := new(Match)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(int(params.MatchID))).Receive(match, apiError)
	return *match, resp, relevantError(err, *apiError)
}
