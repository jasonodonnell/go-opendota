package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExplorerService_Explorer(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/explorer", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"command":"SELECT","rowCount":115,"oid":null,"rows":[{"id":44,"name":"npc_dota_hero_phantom_assassin","localized_name":"Phantom Assassin","primary_attr":"agi","attack_type":"Melee","roles":["Carry","Escape"],"legs":2}],"fields":[{"name":"id","tableID":929536173,"columnID":1,"dataTypeID":23,"dataTypeSize":4,"dataTypeModifier":-1,"format":"text"}],"rowAsArray":false,"err":null}`)
	})

	expected := QueryResult{
		Command:  "SELECT",
		RowCount: 115,
		Rows: []map[string]interface{}{
			map[string]interface{}{
				"id":             float64(44),
				"name":           "npc_dota_hero_phantom_assassin",
				"localized_name": "Phantom Assassin",
				"primary_attr":   "agi",
				"attack_type":    "Melee",
				"roles":          []interface{}{"Carry", "Escape"},
				"legs":           float64(2),
			},
		},
		Fields: []field{
			field{
				Name:             "id",
				TableID:          929536173,
				ColumnID:         1,
				DataTypeID:       23,
				DataTypeSize:     4,
				DataTypeModifier: -1,
				Format:           "text",
			},
		},
		RowAsArray: false,
	}

	query := &ExplorerParam{
		SQL: "SELECT%20*%20FROM%20public.heroes",
	}

	client := NewClient(httpClient)
	queryresult, _, err := client.ExplorerService.Explore(query)
	assert.Nil(t, err)
	assert.Equal(t, expected, queryresult)
}
