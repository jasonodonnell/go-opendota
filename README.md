# go-opendota [![Build Status](https://travis-ci.org/jasonodonnell/go-opendota.png)](https://travis-ci.org/jasonodonnell/go-opendota) [![GoDoc](https://godoc.org/github.com/jasonodonnell/go-opendota?status.png)](https://godoc.org/github.com/jasonodonnell/go-opendota)

Go-Opendota is a simple Go package for accessing the [OpenDota API](https://docs.opendota.com/#).  

Successful queries return native Go structs.

## Services

- [x] Benchmarks
- [x] Distributions
- [x] Explorer
- [x] Health
- [x] Hero Stats
- [x] Heroes
- [x] Leagues
- [x] Live
- [x] Matches
- [x] Metadata
- [x] Players
- [x] Pro Matches
- [x] Pro Players
- [x] Public Matches
- [x] Rankings 
- [x] Records 
- [x] Replays
- [x] Schema
- [x] Search
- [x] Status
- [x] Teams 

## Install

```bash
go get github.com/jasonodonnell/go-opendota/opendota
```

## Examples

Get information about a specific match:

```go
// OpenDota client
client := opendota.NewClient(httpClient)

// Match ID
matchID := &opendota.MatchParam{MatchID: 3559037317}

// Get Match Data
match, res, err := client.MatchService.Match(matchID)
fmt.Println(match.DireTeam.Name, "VS", match.RadiantTeam.Name)
```

Get information about a specific player:

```go
// Player ID
playerID := &opendota.PlayerParam{AccountID: 111620041}

// Get Player Data
player, _, _ := client.PlayerService.Player(playerID)
fmt.Println(player.Profile.Name, player.SoloCompetitiveRank)
```

## License

[MIT License](LICENSE)
