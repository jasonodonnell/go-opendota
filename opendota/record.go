package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

// RecordService provides methods for accessing information
// about top performances in a stat.
type RecordService struct {
	sling *sling.Sling
}

func newRecordService(sling *sling.Sling) *RecordService {
	return &RecordService{
		sling: sling.Path("records/"),
	}
}

// Record is a collection of information about the performance
// for a stat.
type Record struct {
	MatchID   string `json:"match_id"`
	StartTime string `json:"start_time"`
	HeroID    string `json:"hero_id"`
	Score     string `json:"score"`
}

// RecordParam provides the ability to query records by a
// field.
type RecordParam struct {
	Field string `url:"field"`
}

// Records returns a collection of top performance stats for a
// specific field.
func (s *RecordService) Records(param *RecordParam) ([]Record, *http.Response, error) {
	record := new([]Record)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(param.Field).Receive(record, apiError)
	return *record, resp, relevantError(err, *apiError)
}
