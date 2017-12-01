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

// ExplorerService provides a method for exploring
// OpenDota's PostgreSQL database.
type ExplorerService struct {
	sling *sling.Sling
}

type explorerParam struct {
	SQL string `url:"sql"`
}

// QueryResult represents the results returned by OpenDota's
// PostgreSQL database.
type QueryResult struct {
	Command    string                   `json:"command"`
	RowCount   int                      `json:"rowCount"`
	Oid        int                      `json:"oid"`
	Rows       []map[string]interface{} `json:"rows"`
	Fields     []Field                  `json:"fields"`
	RowAsArray bool                     `json:"rowAsArray"`
	Err        interface{}              `json:"err"`
}

type Field struct {
	Name             string `json:"name"`
	TableID          int    `json:"tableID"`
	ColumnID         int    `json:"columnID"`
	DataTypeID       int    `json:"dataTypeID"`
	DataTypeSize     int    `json:"dataTypeSize"`
	DataTypeModifier int    `json:"dataTypeModifier"`
	Format           string `json:"format"`
}

// Explore takes a SQL query as an argument and returns query results.
// https://docs.opendota.com/#tag/explorer%2Fpaths%2F~1explorer%2Fget
func (s *ExplorerService) Explore(query string) (QueryResult, *http.Response, error) {
	param := &explorerParam{}
	param.SQL = query
	queryresult := new(QueryResult)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(param).Receive(queryresult, apiError)
	return *queryresult, resp, relevantError(err, *apiError)
}
