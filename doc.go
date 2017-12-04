/*
Package opendota provides a Client for accessing the OpenDota Api.

Here are some examples of requests:

	// OpenDota client
	client := opendota.NewClient(httpClient)

	// Get Match Data
	match, res, err := client.MatchService.Match(3559037317)
	fmt.Println(match.DireTeam.Name, "VS", match.RadiantTeam.Name)

	// Get Player Data
	player, res, err := client.PlayerService.Player(111620041)
	fmt.Println(player.Profile.Name, player.SoloCompetitiveRank)

	// Player Param
	params := &opendota.PlayerParam{
		Win: 0,
	}

	// Get Won Matches For Player
	wins, res, err := client.PlayerService.Matches(111620041, params)
	for _, game := range wins {
		fmt.Println(game.MatchID, game.HeroID)
	}

All required parameters are passed as arguments to functions.  Additional
arguments can be passed via the param objects.
*/
package opendota
