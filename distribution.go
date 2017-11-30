package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newDistributionService(sling *sling.Sling) *DistributionService {
	return &DistributionService{
		sling: sling.Path("distributions"),
	}
}

// DistributionService provides a method for accesing
// distribution data for Matchmaking Ranking (MMR).
type DistributionService struct {
	sling *sling.Sling
}

// Distribution holds distributions of Matchmaking Ranking (MMR)
// data for each region.
type Distribution struct {
	Mmr        mmr        `json:"mmr"`
	CountryMmr countryMmr `json:"country_mmr"`
}

type countryMmr struct {
	Command    string          `json:"command"`
	RowCount   int             `json:"rowCount"`
	Oid        int             `json:"oid"`
	Rows       []countryMmrRow `json:"rows"`
	Fields     []field         `json:"fields"`
	RowAsArray bool            `json:"rowAsArray"`
}

type countryMmrRow struct {
	Loccountrycode string `json:"loccountrycode"`
	Count          int    `json:"count"`
	Avg            string `json:"avg"`
	Common         string `json:"common"`
}

type mmr struct {
	Command    string   `json:"command"`
	RowCount   int      `json:"rowCount"`
	Oid        int      `json:"oid"`
	Rows       []mmrRow `json:"rows"`
	Fields     []field  `json:"fields"`
	RowAsArray bool     `json:"rowAsArray"`
	Sum        sum      `json:"sum"`
}

type mmrRow struct {
	Bin           int `json:"bin"`
	BinName       int `json:"bin_name"`
	Count         int `json:"count"`
	CumulativeSum int `json:"cumulative_sum"`
}

type sum struct {
	Count int `json:"count"`
}

// Distributions returns a collection of Matchmaking Ranking (MMR) distributions
// for each region.
// https://docs.opendota.com/#tag/distributions%2Fpaths%2F~1distributions%2Fget
func (s *DistributionService) Distributions() (Distribution, *http.Response, error) {
	distribution := new(Distribution)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(distribution, apiError)
	return *distribution, resp, relevantError(err, *apiError)
}
