package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadataService_Metadata(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/metadata", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"banner":null,"cheese":{"cheese":"130","goal":"100"}}`)
	})

	expected := Metadata{
		Cheese: Cheese{
			Cheese: "130",
			Goal:   "100",
		},
	}

	client := NewClient(httpClient)
	metadata, _, err := client.MetadataService.Metadata()
	assert.Nil(t, err)
	assert.Equal(t, expected, metadata)
}
