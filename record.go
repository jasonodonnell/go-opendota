package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newRecordService(sling *sling.Sling) *RecordService {
	return &RecordService{
		sling: sling.Path("records/"),
	}
}

// RecordService provides a method for accessing records for a field.
type RecordService struct {
	sling *sling.Sling
}

// Record represents a record for a field.
type Record struct {
	MatchID   string `json:"match_id"`
	StartTime string `json:"start_time"`
	HeroID    string `json:"hero_id"`
	Score     string `json:"score"`
}

// Records takes a field and returns the records for that field.
// https://docs.opendota.com/#tag/records%2Fpaths%2F~1records~1%7Bfield%7D%2Fget
func (s *RecordService) Records(field string) ([]Record, *http.Response, error) {
	record := new([]Record)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(field).Receive(record, apiError)
	return *record, resp, relevantError(err, *apiError)
}
