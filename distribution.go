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
	Mmr        Mmr        `json:"mmr"`
	CountryMmr CountryMmr `json:"country_mmr"`
}

type CountryMmr struct {
	Command    string          `json:"command"`
	RowCount   int             `json:"rowCount"`
	Oid        int             `json:"oid"`
	Rows       []CountryMmrRow `json:"rows"`
	Fields     []Field         `json:"fields"`
	RowAsArray bool            `json:"rowAsArray"`
}

type CountryMmrRow struct {
	Loccountrycode string `json:"loccountrycode"`
	Count          int    `json:"count"`
	Avg            string `json:"avg"`
	Common         string `json:"common"`
}

type Mmr struct {
	Command    string   `json:"command"`
	RowCount   int      `json:"rowCount"`
	Oid        int      `json:"oid"`
	Rows       []MmrRow `json:"rows"`
	Fields     []Field  `json:"fields"`
	RowAsArray bool     `json:"rowAsArray"`
	Sum        Sum      `json:"sum"`
}

type MmrRow struct {
	Bin           int `json:"bin"`
	BinName       int `json:"bin_name"`
	Count         int `json:"count"`
	CumulativeSum int `json:"cumulative_sum"`
}

type Sum struct {
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
