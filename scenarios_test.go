package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScenariosService_ItemTimings(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/scenarios/itemTimings", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"hero_id":1,"item":"butterfly","games":"166","wins":"105","time":1800}]`)
	})

	expected := []ItemTimings{
		{
			HeroID: 1,
			Item:   "butterfly",
			Games:  "166",
			Wins:   "105",
			Time:   1800,
		},
	}

	params := &ItemTimingsParam{}

	client := NewClient(httpClient)
	itemTimings, _, err := client.ScenariosService.ItemTimings(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, itemTimings)
}

func TestScenariosService_LaneRoles(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/scenarios/laneRoles", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"hero_id":1,"lane_role":1,"games":"166","wins":"105","time":1800}]`)
	})

	expected := []LaneRoles{
		{
			HeroID:   1,
			LaneRole: 1,
			Games:    "166",
			Wins:     "105",
			Time:     1800,
		},
	}

	params := &LaneRolesParam{}

	client := NewClient(httpClient)
	laneRoles, _, err := client.ScenariosService.LaneRoles(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, laneRoles)
}

func TestScenariosService_Misc(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/scenarios/misc", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"scenario":"courier_kill","is_radiant":false,"region":1,"games":"840","wins":"442"}]`)
	})

	expected := []MiscQueryResults{
		{
			Scenario:  "courier_kill",
			IsRadiant: false,
			Region:    1,
			Games:     "840",
			Wins:      "442",
		},
	}

	params := &MiscParam{}

	client := NewClient(httpClient)
	misc, _, err := client.ScenariosService.Misc(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, misc)
}
