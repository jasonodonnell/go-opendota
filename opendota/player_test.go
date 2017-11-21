package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerService_Counts(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/counts", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"leaver_status":{"0":{"games":3441,"win":2025}},"game_mode":{"1":{"games":865,"win":469}},"lobby_type":{"0":{"games":416,"win":243}},"lane_role":{"0":{"games":2322,"win":1312}},"region":{"0":{"games":1,"win":1}},"patch":{"7":{"games":100,"win":58}},"is_radiant":{"0":{"games":1746,"win":995}}}`)
	})

	expected := PlayerCounts{
		LeaverStatus: map[string]GameWins{
			"0": {
				Games: 3441,
				Win:   2025,
			},
		},
		GameMode: map[string]GameWins{
			"1": {
				Games: 865,
				Win:   469,
			},
		},
		LobbyType: map[string]GameWins{
			"0": {
				Games: 416,
				Win:   243,
			},
		},
		LaneRole: map[string]GameWins{
			"0": {
				Games: 2322,
				Win:   1312,
			},
		},
		Region: map[string]GameWins{
			"0": {
				Games: 1,
				Win:   1,
			},
		},
		Patch: map[string]GameWins{
			"7": {
				Games: 100,
				Win:   58,
			},
		},
		IsRadiant: map[string]GameWins{
			"0": {
				Games: 1746,
				Win:   995,
			},
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	counts, _, err := client.PlayerService.Counts(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, counts)
}

func TestPlayerService_Heroes(t *testing.T) {
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
	heroes, _, err := client.PlayerService.Heroes(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, heroes)
}

func TestPlayerService_Histograms(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/histograms/kills", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"x":0,"games":389,"win":143}]`)
	})

	expected := []PlayerHistogram{
		PlayerHistogram{
			X:     0,
			Games: 389,
			Win:   143,
		},
	}

	params := &PlayersParam{
		AccountID: 34505203,
		Field:     "kills",
	}

	client := NewClient(httpClient)
	histograms, _, err := client.PlayerService.Histograms(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, histograms)
}

func TestPlayerService_Matches(t *testing.T) {
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
	matches, _, err := client.PlayerService.Matches(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, matches)
}

func TestPlayerService_Peers(t *testing.T) {
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
			AvatarFull:   "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg",
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	peers, _, err := client.PlayerService.Peers(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, peers)
}

func TestPlayerService_Player(t *testing.T) {
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
		SteamID:        "76561197994770931",
		Avatar:         "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg",
		AvatarMedium:   "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg",
		AvatarFull:     "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg",
		ProfileURL:     "http://steamcommunity.com/id/MinD_ContRoL/",
		LastLogin:      "",
		LocCountryCode: "BG",
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
	player, _, err := client.PlayerService.Player(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, player)
}

func TestPlayerService_Pros(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/pros", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"account_id":34505203,"name":"MinD_ContRoL","country_code":"","fantasy_role":0,"team_id":2163,"team_name":"Team Liquid","team_tag":"Liquid","is_locked":true,"is_pro":true,"locked_until":1533081600,"steamid":"76561197994770931","avatar":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg","avatarmedium":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg","avatarfull":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg","profileurl":"http://steamcommunity.com/id/MinD_ContRoL/","personaname":"$a$uk3","last_login":null,"full_history_time":"2017-10-11T16:03:56.626Z","cheese":0,"fh_unavailable":true,"loccountrycode":"BG","last_match_time":"2017-11-10T21:45:00.000Z","last_played":1510350300,"win":2060,"games":3590,"with_win":2060,"with_games":3590,"against_win":0,"against_games":0,"with_gpm_sum":null,"with_xpm_sum":null}]`)
	})

	expected := []PlayerPros{
		PlayerPros{
			AccountID:       34505203,
			Name:            "MinD_ContRoL",
			CountryCode:     "",
			FantasyRole:     0,
			TeamID:          2163,
			TeamName:        "Team Liquid",
			TeamTag:         "Liquid",
			IsLocked:        true,
			IsPro:           true,
			LockedUntil:     1533081600,
			SteamID:         "76561197994770931",
			Avatar:          "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg",
			AvatarMedium:    "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg",
			AvatarFull:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg",
			ProfileURL:      "http://steamcommunity.com/id/MinD_ContRoL/",
			Personaname:     "$a$uk3",
			FullHistoryTime: "2017-10-11T16:03:56.626Z",
			Cheese:          0,
			FhUnavailable:   true,
			LocCountryCode:  "BG",
			LastMatchTime:   "2017-11-10T21:45:00.000Z",
			LastPlayed:      1510350300,
			Win:             2060,
			Games:           3590,
			WithWin:         2060,
			WithGames:       3590,
			AgainstGames:    0,
			AgainstWin:      0,
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	pros, _, err := client.PlayerService.Pros(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, pros)
}

func TestPlayerService_Rankings(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/rankings", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"hero_id":53,"score":439.08488536972,"percent_rank":1,"card":543680}]`)
	})

	expected := []PlayerRankings{
		PlayerRankings{
			HeroID:      53,
			Score:       439.08488536972,
			PercentRank: 1,
			Card:        543680,
		},
	}

	params := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	rankings, _, err := client.PlayerService.Rankings(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, rankings)
}

func TestPlayerService_Ratings(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/111620041/ratings", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"account_id":111620041,"match_id":2067247580,"solo_competitive_rank":7812,"competitive_rank":5656,"time":"2016-01-09T20:48:25.028Z"}]`)
	})

	expected := []PlayerRatings{
		PlayerRatings{
			AccountID:           111620041,
			MatchID:             2067247580,
			SoloCompetitiveRank: 7812,
			CompetitiveRank:     5656,
			Time:                "2016-01-09T20:48:25.028Z",
		},
	}

	params := &PlayersParam{
		AccountID: 111620041,
	}

	client := NewClient(httpClient)
	ratings, _, err := client.PlayerService.Ratings(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, ratings)
}

func TestPlayerService_RecentMatches(t *testing.T) {
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
	recentMatches, _, err := client.PlayerService.RecentMatches(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, recentMatches)
}

func TestPlayerService_Totals(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/totals", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"field":"kills","n":3590,"sum":27722}]`)
	})

	expected := []PlayerTotals{
		PlayerTotals{
			Field: "kills",
			N:     3590,
			Sum:   27722,
		},
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	totals, _, err := client.PlayerService.Totals(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, totals)
}

func TestPlayerService_WinLoss(t *testing.T) {
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
	winloss, _, err := client.PlayerService.WinLoss(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, winloss)
}

func TestPlayerService_WordCloud(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203/wordcloud", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"my_word_counts":{"gg":800},"all_word_counts":{"gg":5289}}`)
	})

	expected := PlayerWordCloud{
		MyWordCounts: map[string]int{
			"gg": 800,
		},
		AllWordCounts: map[string]int{
			"gg": 5289,
		},
	}

	params := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	wordcloud, _, err := client.PlayerService.WordCloud(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, wordcloud)
}
