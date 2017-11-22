package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newExplorerService(sling *sling.Sling) *ExplorerService {
	return &ExplorerService{
		sling: sling.Path("explorer"),
	}
}

// ExplorerService provides methods for exploring OpenDota's
// PostgreSQL database.
type ExplorerService struct {
	sling *sling.Sling
}

// ExplorerParam is the parameter for specifying the sql query
// to run against the PostgreSQL database.  It should be encoded
// for use in URLs
type ExplorerParam struct {
	SQL string `url:"sql"`
}

// QueryResult is a collection returned by a sql query
// against the PostgreSQL database.
type QueryResult struct {
	Command    string                   `json:"command"`
	RowCount   int                      `json:"rowCount"`
	Oid        int                      `json:"oid"`
	Rows       []map[string]interface{} `json:"rows"`
	Fields     []field                  `json:"fields"`
	RowAsArray bool                     `json:"rowAsArray"`
	Err        interface{}              `json:"err"`
}

type field struct {
	Name             string `json:"name"`
	TableID          int    `json:"tableID"`
	ColumnID         int    `json:"columnID"`
	DataTypeID       int    `json:"dataTypeID"`
	DataTypeSize     int    `json:"dataTypeSize"`
	DataTypeModifier int    `json:"dataTypeModifier"`
	Format           string `json:"format"`
}

// Explore returns a collection for a specific sql query.
func (s *ExplorerService) Explore(param *ExplorerParam) (QueryResult, *http.Response, error) {
	queryresult := new(QueryResult)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(param).Receive(queryresult, apiError)
	return *queryresult, resp, relevantError(err, *apiError)
}
