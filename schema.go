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

// SchemaService provides methods for accessing information
// about the database schema.
type SchemaService struct {
	sling *sling.Sling
}

// Schema is a collection of information about the database schema.
type Schema struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
}

// Schema returns a collection about the database schema.
// https://docs.opendota.com/#tag/schema%2Fpaths%2F~1schema%2Fget
func (s *SchemaService) Schema() ([]Schema, *http.Response, error) {
	schema := new([]Schema)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(schema, apiError)
	return *schema, resp, relevantError(err, *apiError)
}
