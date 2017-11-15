package main

import (
	"fmt"

	"github.com/jasonodonnell/go-opendota/opendota"
)

func main() {
	client := opendota.NewClient(nil)

	// teamParam := &opendota.TeamParam{
	// 	TeamID: 39,
	// }

	// // // Get Teams
	// teams, _, err := client.TeamService.Teams()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// for _, team := range teams {
	// 	fmt.Println(team.Name)
	// }

	// // Get Specific Team
	// team, _, err := client.TeamService.Team(teamParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(team)

	// // Get Specific Team Played Matches
	// matches, _, err := client.TeamService.Matches(teamParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(matches)

	// // Get Players for a team
	// players, _, err := client.TeamService.Players(teamParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(players)

	// // Get heroes stats for a team
	// heroes, _, err := client.TeamService.Heroes(teamParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(heroes)

	matchParam := &opendota.MatchParam{
		MatchID: 3559037317,
	}

	match, _, err := client.MatchService.Match(matchParam)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(match)
	// playerParam := &opendota.PlayersParam{
	// 	AccountID: 34505203,
	// }
	// winlossParam := &opendota.PlayersWinLossParam{
	// 	AccountID: 34505203,
	// }

	// player, _, err := client.PlayersService.Player(playerParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// winloss, res, err := client.PlayersService.WinLoss(winlossParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(winloss, res)
	// matches, _, err := client.PlayersService.Matches(playerParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(matches[0])
	// counts, res, err := client.PlayersService.Counts(playerParam)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(counts, res)
}
