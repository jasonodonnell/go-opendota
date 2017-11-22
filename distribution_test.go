package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributionService_Distributions(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/distributions", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"mmr":{"command":"SELECT","rowCount":98,"oid":null,"rows":[{"bin":0,"bin_name":0,"count":9786,"cumulative_sum":9786}],"fields":[{"name":"bin","tableID":0,"columnID":0,"dataTypeID":23,"dataTypeSize":4,"dataTypeModifier":-1,"format":"text"}],"rowAsArray":false,"sum":{"count":2444278}},"country_mmr":{"command":"SELECT","rowCount":255,"oid":null,"rows":[{"loccountrycode":"AN","count":1,"avg":"4113","common":"AN"}],"fields":[{"name":"loccountrycode","tableID":16405,"columnID":12,"dataTypeID":1043,"dataTypeSize":-1,"dataTypeModifier":6,"format":"text"}],"rowAsArray":false}}`)
	})

	expected := Distribution{
		Mmr: mmr{
			Command:  "SELECT",
			RowCount: 98,
			Rows: []mmrRow{
				mmrRow{
					Bin:           0,
					BinName:       0,
					Count:         9786,
					CumulativeSum: 9786,
				},
			},
			Fields: []field{
				field{
					Name:             "bin",
					TableID:          0,
					ColumnID:         0,
					DataTypeID:       23,
					DataTypeSize:     4,
					DataTypeModifier: -1,
					Format:           "text",
				},
			},
			RowAsArray: false,
			Sum: sum{
				Count: 2444278,
			},
		},
		CountryMmr: countryMmr{
			Command:  "SELECT",
			RowCount: 255,
			Rows: []countryMmrRow{
				countryMmrRow{
					Loccountrycode: "AN",
					Count:          1,
					Avg:            "4113",
					Common:         "AN",
				},
			},
			Fields: []field{
				field{
					Name:             "loccountrycode",
					TableID:          16405,
					ColumnID:         12,
					DataTypeID:       1043,
					DataTypeSize:     -1,
					DataTypeModifier: 6,
					Format:           "text",
				},
			},
			RowAsArray: false,
		},
	}

	client := NewClient(httpClient)
	distributions, _, err := client.DistributionService.Distributions()
	assert.Nil(t, err)
	assert.Equal(t, expected, distributions)
}
