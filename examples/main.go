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

	playerID := &opendota.PlayerParam{AccountID: 111620041}
	player, _, err := client.PlayerService.Player(playerID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(player.Profile.Name, player.SoloCompetitiveRank)
}
