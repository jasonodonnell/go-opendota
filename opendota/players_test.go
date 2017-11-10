package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayersService_Player(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"tracked_until":"","solo_competitive_rank":"","profile":{"account_id":34505203,"personaname":"$a$uk3","name":"MinD_ContRoL","cheese":0,"steamid":"76561197994770931","avatar":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg","avatarmedium":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg","avatarfull":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg","profileurl":"http://steamcommunity.com/id/MinD_ContRoL/","last_login":"","loccountrycode":"BG"},"mmr_estimate":{"estimate":8009},"competitive_rank":""}`)
	})

	profile := Profile{
		AccountID:      34505203,
		Personaname:    "$a$uk3",
		Name:           "MinD_ContRoL",
		Cheese:         0,
		Steamid:        "76561197994770931",
		Avatar:         "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg",
		Avatarmedium:   "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg",
		Avatarfull:     "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg",
		Profileurl:     "http://steamcommunity.com/id/MinD_ContRoL/",
		LastLogin:      "",
		Loccountrycode: "BG",
	}

	mmrestimate := MmrEstimate{
		Estimate: 8009,
	}

	expected := Player{
		TrackedUntil:        "",
		SoloCompetitiveRank: "",
		MmrEstimate:         mmrestimate,
		Profile:             profile,
		CompetitiveRank:     "",
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	player, _, err := client.PlayersService.Player(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, player)
}

func TestPlayersService_WinLoss(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/wl", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"win":2059,"lose":1529}`)
	})

	expected := WinLoss{
		Win:  2059,
		Lose: 1529,
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	winloss, _, err := client.PlayersService.WinLoss(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, winloss)
}

func TestPlayersService_RecentMatches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/recentMatches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":3552339511,"player_slot":129,"radiant_win":true,"duration":1801,"game_mode":2,"lobby_type":1,"hero_id":55,"start_time":1510246501,"version":20,"kills":1,"deaths":3,"assists":4,"skill":null,"xp_per_min":337,"gold_per_min":367,"hero_damage":8198,"tower_damage":557,"hero_healing":3309,"last_hits":167,"lane":1,"lane_role":3,"is_roaming":false,"cluster":134,"leaver_status":0,"party_size":10}]`)
	})

	expected := []PlayerMatch{
		PlayerMatch{
			MatchID:      3552339511,
			PlayerSlot:   129,
			RadiantWin:   true,
			Duration:     1801,
			GameMode:     2,
			LobbyType:    1,
			HeroID:       55,
			StartTime:    1510246501,
			Version:      20,
			Kills:        1,
			Deaths:       3,
			Assists:      4,
			Skill:        0,
			XpPerMin:     337,
			GoldPerMin:   367,
			HeroDamage:   8198,
			TowerDamage:  557,
			HeroHealing:  3309,
			LastHits:     167,
			Lane:         1,
			LaneRole:     3,
			IsRoaming:    false,
			Cluster:      134,
			LeaverStatus: 0,
			PartySize:    10,
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	recentMatches, _, err := client.PlayersService.RecentMatches(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, recentMatches)
}

func TestPlayersService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/matches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":3552339511,"player_slot":129,"radiant_win":true,"duration":1801,"game_mode":2,"lobby_type":1,"hero_id":55,"start_time":1510246501,"version":20,"kills":1,"deaths":3,"assists":4,"skill":null,"xp_per_min":337,"gold_per_min":367,"hero_damage":8198,"tower_damage":557,"hero_healing":3309,"last_hits":167,"lane":1,"lane_role":3,"is_roaming":false,"cluster":134,"leaver_status":0,"party_size":10}]`)
	})

	expected := []PlayerMatch{
		PlayerMatch{
			MatchID:      3552339511,
			PlayerSlot:   129,
			RadiantWin:   true,
			Duration:     1801,
			GameMode:     2,
			LobbyType:    1,
			HeroID:       55,
			StartTime:    1510246501,
			Version:      20,
			Kills:        1,
			Deaths:       3,
			Assists:      4,
			Skill:        0,
			XpPerMin:     337,
			GoldPerMin:   367,
			HeroDamage:   8198,
			TowerDamage:  557,
			HeroHealing:  3309,
			LastHits:     167,
			Lane:         1,
			LaneRole:     3,
			IsRoaming:    false,
			Cluster:      134,
			LeaverStatus: 0,
			PartySize:    10,
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	matches, _, err := client.PlayersService.Matches(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, matches)
}

func TestPlayersService_Heroes(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/heroes", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"hero_id":"11","last_played":1453737676,"games":166,"win":105,"with_games":182,"with_win":94,"against_games":273,"against_win":156}]`)
	})

	expected := []PlayerHero{
		PlayerHero{
			HeroID:       "11",
			LastPlayed:   1453737676,
			Games:        166,
			Win:          105,
			WithGames:    182,
			WithWin:      94,
			AgainstGames: 273,
			AgainstWin:   156,
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	heroes, _, err := client.PlayersService.Heroes(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, heroes)
}

func TestPlayersService_Peers(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/peers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"account_id":72312627,"last_played":1510246501,"win":394,"games":606,"with_win":386,"with_games":591,"against_win":8,"against_games":15,"with_gpm_sum":232910,"with_xpm_sum":240982,"personaname":"Notail43","last_login":null,"avatar":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg","avatarfull":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg"}]`)
	})

	expected := []PlayerPeers{
		PlayerPeers{
			AccountID:    72312627,
			LastPlayed:   1510246501,
			Win:          394,
			Games:        606,
			WithWin:      386,
			WithGames:    591,
			AgainstWin:   8,
			AgainstGames: 15,
			WithGpmSum:   232910,
			WithXpmSum:   240982,
			Personaname:  "Notail43",
			LastLogin:    "",
			Avatar:       "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg",
			Avatarfull:   "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg",
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	peers, _, err := client.PlayersService.Peers(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, peers)
}
