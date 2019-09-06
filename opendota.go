package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

const openDotaAPI = "https://api.opendota.com/api/"

// Client for making Open Dota API requests
type Client struct {
	sling               *sling.Sling
	BenchmarkService    *BenchmarkService
	DistributionService *DistributionService
	ExplorerService     *ExplorerService
	HealthService       *HealthService
	HeroService         *HeroService
	HeroStatService     *HeroStatService
	LeagueService       *LeagueService
	LiveService         *LiveService
	MatchService        *MatchService
	MetadataService     *MetadataService
	PlayerService       *PlayerService
	ProMatchService     *ProMatchService
	ProPlayerService    *ProPlayerService
	PublicMatchService  *PublicMatchService
	RankingService      *RankingService
	RecordService       *RecordService
	ReplayService       *ReplayService
	ScenariosService    *ScenariosService
	SchemaService       *SchemaService
	SearchService       *SearchService
	StatusService       *StatusService
	TeamService         *TeamService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(openDotaAPI)

	return &Client{
		sling:               base,
		BenchmarkService:    newBenchmarkService(base.New()),
		DistributionService: newDistributionService(base.New()),
		ExplorerService:     newExplorerService(base.New()),
		HealthService:       newHealthService(base.New()),
		HeroService:         newHeroService(base.New()),
		HeroStatService:     newHeroStatService(base.New()),
		LeagueService:       newLeagueService(base.New()),
		LiveService:         newLiveService(base.New()),
		MatchService:        newMatchService(base.New()),
		MetadataService:     newMetadataService(base.New()),
		PlayerService:       newPlayerService(base.New()),
		ProMatchService:     newProMatchService(base.New()),
		ProPlayerService:    newProPlayerService(base.New()),
		PublicMatchService:  newPublicMatchService(base.New()),
		RankingService:      newRankingService(base.New()),
		RecordService:       newRecordService(base.New()),
		ReplayService:       newReplayService(base.New()),
		ScenariosService:    newScenariosService(base.New()),
		SchemaService:       newSchemaService(base.New()),
		SearchService:       newSearchService(base.New()),
		StatusService:       newStatusService(base.New()),
		TeamService:         newTeamService(base.New()),
	}
}
