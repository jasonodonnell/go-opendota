package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newSchemaService(sling *sling.Sling) *SchemaService {
	return &SchemaService{
		sling: sling.Path("schema"),
	}
}

// SchemaService provides a method for accessing the OpenDota
// database schema.
type SchemaService struct {
	sling *sling.Sling
}

// Schema represents the database schema for OpenDota.
type Schema struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
}

// Schema returns the OpenDota database schema.
// https://docs.opendota.com/#tag/schema%2Fpaths%2F~1schema%2Fget
func (s *SchemaService) Schema() ([]Schema, *http.Response, error) {
	schema := new([]Schema)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(schema, apiError)
	return *schema, resp, relevantError(err, *apiError)
}
