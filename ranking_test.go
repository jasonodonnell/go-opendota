package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankingService_Rankings(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/rankings", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"hero_id":36,"rankings":[{"account_id":40586005,"score":1171.07373076237,"personaname":"wwd","name":"WWD","avatar":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/fb/fbb125a490522eac46c427e5cd2f6c1a161bab68.jpg","last_login":"2015-09-21T20:02:09.044Z","solo_competitive_rank":7532}]}`)
	})

	expected := Ranking{
		HeroID: 36,
		Rankings: []ranking{
			ranking{
				AccountID:           40586005,
				Score:               1171.07373076237,
				Personaname:         "wwd",
				Name:                "WWD",
				Avatar:              "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/fb/fbb125a490522eac46c427e5cd2f6c1a161bab68.jpg",
				LastLogin:           "2015-09-21T20:02:09.044Z",
				SoloCompetitiveRank: 7532,
			},
		},
	}

	param := &RankingParam{
		HeroID: "36",
	}

	client := NewClient(httpClient)
	ranking, _, err := client.RankingService.Rankings(param)
	assert.Nil(t, err)
	assert.Equal(t, expected, ranking)
}
