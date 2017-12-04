package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamService_Heroes(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/teams/2163/heroes", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"hero_id":86,"localized_name":"Rubick","games_played":169,"wins":104}]`)
	})

	expected := []TeamHeroes{
		{
			HeroID:        86,
			LocalizedName: "Rubick",
			GamesPlayed:   169,
			Wins:          104,
		},
	}

	client := NewClient(httpClient)
	heroes, _, err := client.TeamService.Heroes(2163)
	assert.Nil(t, err)
	assert.Equal(t, expected, heroes)
}

func TestTeamService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/teams/2163/matches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"match_id":3528450522,"radiant_win":true,"radiant":true,"duration":2024,"start_time":1509224338,"leagueid":5609,"league_name":"ESL One Hamburg 2017","cluster":132}]`)
	})

	expected := []TeamMatch{
		{
			MatchID:    3528450522,
			RadiantWin: true,
			Radiant:    true,
			Duration:   2024,
			StartTime:  1509224338,
			LeagueID:   5609,
			LeagueName: "ESL One Hamburg 2017",
			Cluster:    132,
		},
	}

	client := NewClient(httpClient)
	matches, _, err := client.TeamService.Matches(2163)
	assert.Nil(t, err)
	assert.Equal(t, expected, matches)
}

func TestTeamService_Players(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/teams/2163/players", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"account_id":30237211,"name":"BuLba","games_played":545,"wins":312,"is_current_team_member":false}]`)
	})

	expected := []TeamPlayers{
		{
			AccountID:           30237211,
			Name:                "BuLba",
			GamesPlayed:         545,
			Wins:                312,
			IsCurrentTeamMember: false,
		},
	}

	client := NewClient(httpClient)
	players, _, err := client.TeamService.Players(2163)
	assert.Nil(t, err)
	assert.Equal(t, expected, players)
}

func TestTeamService_Team(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/teams/2163", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"team_id":2163,"rating":1628.74,"wins":582,"losses":368,"last_match_time":1509224338,"name":"Team Liquid","tag":"Liquid","logo_url":"http://cloud-3.steamusercontent.com/ugc/858347654776522964/E70F0E063879154A1982B3C907D6A5DFDA183BF9/"}`)
	})

	expected := Team{
		TeamID:        2163,
		Rating:        1628.74,
		Wins:          582,
		Losses:        368,
		LastMatchTime: 1509224338,
		Name:          "Team Liquid",
		Tag:           "Liquid",
		LogoURL:       "http://cloud-3.steamusercontent.com/ugc/858347654776522964/E70F0E063879154A1982B3C907D6A5DFDA183BF9/",
	}

	client := NewClient(httpClient)
	team, _, err := client.TeamService.Team(2163)
	assert.Nil(t, err)
	assert.Equal(t, expected, team)
}

func TestTeamService_Teams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/teams/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"team_id":2163,"rating":1628.74,"wins":582,"losses":368,"last_match_time":1509224338,"name":"Team Liquid","tag":"Liquid","logo_url":"http://cloud-3.steamusercontent.com/ugc/858347654776522964/E70F0E063879154A1982B3C907D6A5DFDA183BF9/"}]`)
	})

	expected := []Team{
		{
			TeamID:        2163,
			Rating:        1628.74,
			Wins:          582,
			Losses:        368,
			LastMatchTime: 1509224338,
			Name:          "Team Liquid",
			Tag:           "Liquid",
			LogoURL:       "http://cloud-3.steamusercontent.com/ugc/858347654776522964/E70F0E063879154A1982B3C907D6A5DFDA183BF9/",
		},
	}

	client := NewClient(httpClient)
	teams, _, err := client.TeamService.Teams()
	assert.Nil(t, err)
	assert.Equal(t, expected, teams)
}
