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
