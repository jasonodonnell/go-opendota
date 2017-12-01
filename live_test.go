package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveService_Live(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/live", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"activate_time":1511274112,"deactivate_time":0,"server_steam_id":"90112032772868097","lobby_id":"25354178522305816","league_id":0,"lobby_type":7,"game_time":332,"delay":120,"spectators":24,"game_mode":22,"average_mmr":7163,"sort_score":7687,"last_update_time":1511274752,"radiant_lead":-10,"radiant_score":3,"dire_score":5,"players":[{"account_id":111473138,"hero_id":112,"name":"Panda","country_code":"id","fantasy_role":1,"team_id":3262512,"team_name":"ThePrime.NND","team_tag":"TP.NND","is_locked":false,"is_pro":true,"locked_until":1493683200}],"building_state":4784201}]`)
	})

	expected := []LiveGame{
		LiveGame{
			ActivateTime:   1511274112,
			DeactivateTime: 0,
			ServerSteamID:  "90112032772868097",
			LobbyID:        "25354178522305816",
			LeagueID:       0,
			LobbyType:      7,
			GameTime:       332,
			Delay:          120,
			Spectators:     24,
			GameMode:       22,
			AverageMmr:     7163,
			SortScore:      7687,
			LastUpdateTime: 1511274752,
			RadiantLead:    -10,
			RadiantScore:   3,
			DireScore:      5,
			Players: []LivePlayers{
				LivePlayers{
					AccountID:   111473138,
					HeroID:      112,
					Name:        "Panda",
					CountryCode: "id",
					FantasyRole: 1,
					TeamID:      3262512,
					TeamName:    "ThePrime.NND",
					TeamTag:     "TP.NND",
					IsLocked:    false,
					IsPro:       true,
					LockedUntil: 1493683200,
				},
			},
			BuildingState: 4784201,
		},
	}

	client := NewClient(httpClient)
	livegames, _, err := client.LiveService.Live()
	assert.Nil(t, err)
	assert.Equal(t, expected, livegames)
}
