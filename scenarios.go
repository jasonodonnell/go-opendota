package opendota

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

func newScenariosService(sling *sling.Sling) *ScenariosService {
	return &ScenariosService{
		sling: sling.Path("scenarios/"),
	}
}

// ScenariosService provides a method for accessing OpenDota
// scenario information.
type ScenariosService struct {
	sling *sling.Sling
}

// ItemTimingsParam is used for customizing item timing queries.
type ItemTimingsParam struct {
	HeroID int    `url:"hero_id,omitempty"`
	Item   string `url:"item,omitempty"`
}

// ItemTimings represents item timing information.
type ItemTimings struct {
	HeroID int    `json:"hero_id"`
	Item   string `json:"item"`
	Time   int    `json:"time"`
	Games  string `json:"games"`
	Wins   string `json:"wins"`
}

// LaneRolesParam is used for customizing lane role queries.
type LaneRolesParam struct {
	HeroID   int    `url:"hero_id,omitempty"`
	LaneRole string `url:"lane_role,omitempty"`
}

// LaneRoles represents lane role information.
type LaneRoles struct {
	HeroID   int    `json:"hero_id"`
	LaneRole int    `json:"lane_role"`
	Time     int    `json:"time"`
	Games    string `json:"games"`
	Wins     string `json:"wins"`
}

// MiscParam is used for creating a misc scenario query.
type MiscParam struct {
	Scenario string `url:"scenario,omitempty"`
}

// MiscQueryResults contains results for the scenario query.
type MiscQueryResults struct {
	Scenario  string `json:"scenario"`
	IsRadiant bool   `json:"is_radiant"`
	Region    int    `json:"region"`
	Games     string `json:"games"`
	Wins      string `json:"wins"`
}

// ItemTimings returns information about item timings and winrates.
// https://docs.opendota.com/#tag/scenarios%2Fpaths%2F~1scenarios~1itemTimings%2Fget
func (s *ScenariosService) ItemTimings(params *ItemTimingsParam) ([]ItemTimings, *http.Response, error) {
	if params == nil {
		params = &ItemTimingsParam{}
	}

	itemTimings := new([]ItemTimings)
	apiError := new(APIError)
	path := "itemTimings"
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(itemTimings, apiError)
	return *itemTimings, resp, relevantError(err, *apiError)
}

// LaneRoles returns information about hero lanes and winrates.
// https://docs.opendota.com/#tag/scenarios%2Fpaths%2F~1scenarios~1laneRoless%2Fget
func (s *ScenariosService) LaneRoles(params *LaneRolesParam) ([]LaneRoles, *http.Response, error) {
	if params == nil {
		params = &LaneRolesParam{}
	}

	laneRoles := new([]LaneRoles)
	apiError := new(APIError)
	path := "laneRoles"
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(laneRoles, apiError)
	return *laneRoles, resp, relevantError(err, *apiError)
}

// Misc returns information about winrates given a query.
// https://docs.opendota.com/#tag/scenarios%2Fpaths%2F~1scenarios~1misc%2Fget
func (s *ScenariosService) Misc(params *MiscParam) ([]MiscQueryResults, *http.Response, error) {
	if params == nil {
		params = &MiscParam{}
	}

	results := new([]MiscQueryResults)
	apiError := new(APIError)
	path := "misc"
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(results, apiError)
	fmt.Printf("%v\n%v", resp.Request, err)
	return *results, resp, relevantError(err, *apiError)
}
