package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeroService_Durations(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/heroes/1/durations", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"duration_bin":5700,"games_played":1,"wins":0}]`)
	})

	expected := []HeroDuration{
		HeroDuration{
			DurationBin: 5700,
			GamesPlayed: 1,
			Wins:        0,
		},
	}

	param := &HeroParam{
		HeroID: 1,
	}

	client := NewClient(httpClient)
	durations, _, err := client.HeroService.Durations(param)
	assert.Nil(t, err)
	assert.Equal(t, expected, durations)
}

func TestHeroService_Heroes(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/heroes/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"id":1,"name":"npc_dota_hero_antimage","localized_name":"Anti-Mage","primary_attr":"agi","attack_type":"Melee","roles":["Carry","Escape","Nuker"],"legs":2}]`)
	})

	expected := []Hero{
		Hero{
			ID:            1,
			Name:          "npc_dota_hero_antimage",
			LocalizedName: "Anti-Mage",
			PrimaryAttr:   "agi",
			AttackType:    "Melee",
			Roles:         []string{"Carry", "Escape", "Nuker"},
			Legs:          2,
		},
	}

	client := NewClient(httpClient)
	heroes, _, err := client.HeroService.Heroes()
	assert.Nil(t, err)
	assert.Equal(t, expected, heroes)
}

func TestHeroService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/heroes/1/matches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":3572814592,"start_time":1511183518,"duration":1562,"radiant_win":false,"leagueid":4127,"league_name":"Uprise Champions Cup: Special","radiant":false,"account_id":111637216,"kills":1,"deaths":2,"assists":5}]`)
	})

	expected := []HeroMatch{
		HeroMatch{
			MatchID:    3572814592,
			StartTime:  1511183518,
			Duration:   1562,
			RadiantWin: false,
			LeagueID:   4127,
			LeagueName: "Uprise Champions Cup: Special",
			Radiant:    false,
			AccountID:  111637216,
			Kills:      1,
			Deaths:     2,
			Assists:    5,
		},
	}

	param := &HeroParam{
		HeroID: 1,
	}

	client := NewClient(httpClient)
	matches, _, err := client.HeroService.Matches(param)
	assert.Nil(t, err)
	assert.Equal(t, expected, matches)
}

func TestHeroService_Matchups(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/heroes/1/matchups", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"hero_id":55,"games_played":95,"wins":54}]`)
	})

	expected := []HeroMatchup{
		HeroMatchup{
			HeroID:      55,
			GamesPlayed: 95,
			Wins:        54,
		},
	}

	param := &HeroParam{
		HeroID: 1,
	}

	client := NewClient(httpClient)
	matchups, _, err := client.HeroService.Matchups(param)
	assert.Nil(t, err)
	assert.Equal(t, expected, matchups)
}

func TestHeroService_Players(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/heroes/1/players", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"account_id":21289303,"games_played":115,"wins":72}]`)
	})

	expected := []HeroPlayer{
		HeroPlayer{
			AccountID:   21289303,
			GamesPlayed: 115,
			Wins:        72,
		},
	}

	param := &HeroParam{
		HeroID: 1,
	}

	client := NewClient(httpClient)
	players, _, err := client.HeroService.Players(param)
	assert.Nil(t, err)
	assert.Equal(t, expected, players)
}
