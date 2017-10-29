package main

import (
	"fmt"

	"github.com/jasonodonnell/go-opendota/opendota"
)

func main() {
	client := opendota.NewClient(nil)

	teams, _, err := client.TeamService.Teams()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, team := range teams {
		fmt.Println(team.Name)
	}
}
