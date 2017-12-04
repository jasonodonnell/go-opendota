package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaService_Schema(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/schema", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"table_name":"team_rating","column_name":"team_id","data_type":"bigint"},{"table_name":"team_rating","column_name":"rating","data_type":"real"}]`)
	})

	expected := []Schema{
		{
			TableName:  "team_rating",
			ColumnName: "team_id",
			DataType:   "bigint",
		},
		{
			TableName:  "team_rating",
			ColumnName: "rating",
			DataType:   "real",
		},
	}

	client := NewClient(httpClient)
	schema, _, err := client.SchemaService.Schema()
	assert.Nil(t, err)
	assert.Equal(t, expected, schema)
}
