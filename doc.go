/*
Package opendota provides a Client for accessing the OpenDota Api.

Here are some examples of requests:

	// OpenDota client
	client := opendota.NewClient(httpClient)

	// Match ID
	matchID := &opendota.MatchParam{MatchID: 3559037317}

	// Get Match Data
	match, res, err := client.MatchService.Match(matchID)
	fmt.Println(match.DireTeam.Name, "VS", match.RadiantTeam.Name)

	// Player ID
	playerID := &opendota.PlayerParam{AccountID: 111620041}

	// Get Player Data
	player, _, _ := client.PlayerService.Player(playerID)
	fmt.Println(player.Profile.Name, player.SoloCompetitiveRank)

*/
package opendota
