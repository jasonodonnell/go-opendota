package opendota

/* func TestGameService_Top(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/games/top", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"limit": "10", "offset": "0"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"total": 1157, "top": [""]`) // TODO FIX THIS
	})

	expected := &TopGames{
		Total: 1157,
		Top: []top{{
			Channels: 953,
			Viewers:  171708,
			Game: game{
				ID: 32399,
				Box: image{
					Large:    "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-272x380.jpg",
					Medium:   "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-136x190.jpg",
					Small:    "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-52x72.jpg",
					Template: "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-{width}x{height}.jpg",
				},
				GiantbombID: 36113,
				Logo: image{
					Large:    "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-272x380.jpg",
					Medium:   "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-136x190.jpg",
					Small:    "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-52x72.jpg",
					Template: "https://static-cdn.jtvnw.net/ttv-boxart/Counter-Strike:%20Global%20Offensive-{width}x{height}.jpg",
				},
				Name:       "Counter-Strike: Global Offensive",
				Popularity: 170487,
			},
		}},
	}
	client := NewClient("0", httpClient)
	params := &TopGamesParams{
		Limit:  10,
		Offset: 0,
	}
	topGames, _, err := client.GameService.Top(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, topGames)
}
*/
