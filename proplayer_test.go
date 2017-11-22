package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProPlayerService_Players(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/proPlayers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"account_id":34505203,"steamid":"76561197994770931","avatar":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg","avatarmedium":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg","avatarfull":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg","profileurl":"http://steamcommunity.com/id/MinD_ContRoL/","personaname":"$a$uk3","last_login":null,"full_history_time":"2017-10-11T16:03:56.626Z","cheese":0,"fh_unavailable":true,"loccountrycode":"BG","last_match_time":"2017-11-10T21:45:00.000Z","name":"MinD_ContRoL","country_code":"","fantasy_role":0,"team_id":2163,"team_name":"Team Liquid","team_tag":"Liquid","is_locked":true,"is_pro":true,"locked_until":1533081600}]`)
	})

	expected := []ProPlayer{
		ProPlayer{
			AccountID:       34505203,
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
			Name:            "MinD_ContRoL",
			CountryCode:     "",
			FantasyRole:     0,
			TeamID:          2163,
			TeamName:        "Team Liquid",
			TeamTag:         "Liquid",
			IsLocked:        true,
			IsPro:           true,
			LockedUntil:     1533081600,
		},
	}

	client := NewClient(httpClient)
	proplayers, _, err := client.ProPlayerService.Players()
	assert.Nil(t, err)
	assert.Equal(t, expected, proplayers)
}
