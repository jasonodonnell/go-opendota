package main

import (
	"fmt"

	opendota "github.com/jasonodonnell/go-opendota"
)

func main() {

	// // OpenDota Client
	client := opendota.NewClient(nil)
	// param := &opendota.BenchmarkParam{
	// 	HeroID: "36",
	// }

	// // Benchmarks
	// benchmark, _, err := client.BenchmarkService.Benchmarks(param)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(benchmark.Result.GoldPerMin)

	// // MMR Distributions
	// distribution, _, err := client.DistributionService.Distributions()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(distribution)

	// // SQL Explorer
	// param := &opendota.ExplorerParam{
	// 	SQL: "SELECT * FROM public.heroes",
	// }

	// explorer, _, err := client.ExplorerService.Explore(param)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(explorer)

	// param := &opendota.MatchParam{
	// 	MatchID: 3559037317,
	// }

	// match, _, _ := client.MatchService.Match(param)
	// fmt.Println(match.DireTeam.Name, "VS", match.RadiantTeam.Name)
	// for _, player := range match.Players {
	// 	fmt.Println(player.Name, player.AccountID)
	// }

	// OpenDota client
	//client := opendota.NewClient(httpClient)

	// Get Match Data
	match, _, _ := client.MatchService.Match(3559037317)
	fmt.Println(match.DireTeam.Name, "VS", match.RadiantTeam.Name)

	// Get Player Data
	player, _, _ := client.PlayerService.Player(111620041)
	fmt.Println(player.Profile.Name, player.SoloCompetitiveRank)

	// Player Param
	params := &opendota.PlayerParam{
		Win: 0,
	}

	// Get Won Matches For Player
	wins, _, _ := client.PlayerService.Matches(111620041, params)
	for _, game := range wins {
		fmt.Println(game.MatchID, game.HeroID)
	}
}
