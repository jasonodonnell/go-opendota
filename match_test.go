package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchService_Match(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/matches/3559037317", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"match_id":3559037317,"barracks_status_dire":48,"barracks_status_radiant":63,"chat":[{"time":-41,"type":"chat","unit":"k y le","key":"glhf","slot":5,"player_slot":128}],"cluster":123,"cosmetics":{"5242":3},"dire_score":15,"duration":2231,"engine":1,"first_blood_time":120,"game_mode":2,"human_players":10,"leagueid":5627,"lobby_type":1,"match_seq_num":3095917221,"negative_votes":12,"objectives":[{"time":31,"type":"CHAT_MESSAGE_COURIER_LOST","team":2}],"picks_bans":[{"is_pick":false,"hero_id":60,"team":1,"order":0,"ord":0,"match_id":3559037317}],"positive_votes":112,"radiant_gold_adv":[0],"radiant_score":27,"radiant_win":true,"radiant_xp_adv":[0],"skill":null,"start_time":1510535111,"teamfights":[{"start":266,"end":306,"last_death":291,"deaths":3,"players":[{"deaths_pos":{},"ability_uses":{"puck_illusory_orb":2,"puck_ethereal_jaunt":1},"item_uses":{"tpscroll":1,"flask":1},"deaths":0,"buybacks":0,"damage":225,"healing":295,"gold_delta":224,"xp_delta":311,"xp_start":1570,"xp_end":1881}]}],"tower_status_dire":384,"tower_status_radiant":1830,"version":20,"replay_salt":897752925,"series_id":177752,"series_type":1,"league":{"leagueid":5627,"ticket":"econ/leagues/subscriptions_dreamleague_season_8","banner":"econ/leagues/subscriptions_dreamleague_season_8_ingame","tier":"premium","name":"DreamLeague season 8"},"radiant_team":{"team_id":39,"name":"Evil Geniuses","tag":"EG","logo_url":"http://cloud-3.steamusercontent.com/ugc/142255738559146189/495FE7DC43BFAE03C5237446AB526888CA81827E/"},"dire_team":{"team_id":3,"name":"compLexity Gaming","tag":"coL","logo_url":"http://cloud-3.steamusercontent.com/ugc/398960342836468051/26C1C5F6C3082FF2A1B706435F7DC0FCF5F5D27F/"},"players":[{"match_id":3559037317,"player_slot":0,"ability_upgrades_arr":[5069],"ability_uses":{"puck_dream_coil":9},"account_id":111620041,"actions":{"1":4247},"additional_units":null,"assists":14,"backpack_0":41,"backpack_1":46,"backpack_2":0,"buyback_log":[{"time":2101,"slot":0,"type":"buyback_log","player_slot":0}],"camps_stacked":2,"creeps_stacked":5,"damage":{"npc_dota_creep_badguys_melee":56322},"damage_inflictor":{"cyclone":39},"damage_inflictor_received":{"jakiro_liquid_fire":71},"damage_taken":{"npc_dota_badguys_tower4":1101},"deaths":3,"denies":8,"dn_t":[8],"firstblood_claimed":0,"gold":3339,"gold_per_min":462,"gold_reasons":{"0":1378},"gold_spent":13340,"gold_t":[0],"hero_damage":15234,"hero_healing":0,"hero_hits":{"cyclone":1},"hero_id":13,"item_0":1,"item_1":63,"item_2":100,"item_3":36,"item_4":0,"item_5":108,"item_uses":{"cheese":1},"killed":{"npc_dota_creep_badguys_melee":120},"killed_by":{"npc_dota_hero_skeleton_king":1},"kills":3,"kills_log":[{"time":1218,"key":"npc_dota_hero_jakiro"}],"lane_pos":{"74":{"74":2}},"last_hits":222,"leaver_status":0,"level":22,"lh_t":[0],"life_state":{"0":2291},"max_hero_hit":{"type":"max_hero_hit","time":1895,"max":true,"inflictor":"puck_waning_rift","unit":"npc_dota_hero_puck","key":"npc_dota_hero_skeleton_king","value":242,"slot":0,"player_slot":0},"obs":{"160":{"86":1}},"obs_left_log":[{"time":319,"type":"obs_left_log","key":"[160, 86]","slot":0,"x":160,"y":86,"z":130,"entityleft":true,"ehandle":13091570,"player_slot":0}],"obs_log":[{"time":-47,"type":"obs_log","key":"[160, 86]","slot":0,"x":160,"y":86,"z":130,"entityleft":false,"ehandle":13091570,"player_slot":0}],"obs_placed":1,"party_id":0,"party_size":10,"performance_others":null,"permanent_buffs":null,"pings":37,"pred_vict":false,"purchase":{"ultimate_scepter":1},"purchase_log":[{"time":-89,"key":"circlet"}],"randomed":false,"repicked":null,"roshans_killed":0,"rune_pickups":11,"runes":{"1":1},"runes_log":[{"time":659,"key":5}],"sen":{"108":{"122":1}},"sen_left_log":[{"time":1580,"type":"sen_left_log","key":"[108, 122]","slot":0,"x":108,"y":122,"z":130,"entityleft":true,"ehandle":10011357,"player_slot":0}],"sen_log":[{"time":1334,"type":"sen_log","key":"[108, 122]","slot":0,"x":108,"y":122,"z":130,"entityleft":false,"ehandle":10011357,"player_slot":0}],"sen_placed":1,"stuns":31.331408,"teamfight_participation":0.6296296,"times":[0],"tower_damage":1355,"towers_killed":0,"xp_per_min":557,"xp_reasons":{"0":1069,"1":6912,"2":12440,"3":313},"xp_t":[0],"personaname":"t.Danger","name":"SumaiL","last_login":null,"radiant_win":true,"start_time":1510535111,"duration":2231,"cluster":123,"lobby_type":1,"game_mode":2,"patch":26,"region":2,"isRadiant":true,"win":1,"lose":0,"total_gold":17178,"total_xp":20711,"kills_per_min":0.08068130883012103,"kda":4,"abandons":0,"neutral_kills":54,"tower_kills":0,"courier_kills":0,"lane_kills":166,"hero_kills":5,"observer_kills":2,"sentry_kills":0,"roshan_kills":0,"necronomicon_kills":0,"ancient_kills":0,"buyback_count":1,"observer_uses":0,"sentry_uses":1,"lane_efficiency":0.5364970842549769,"lane_efficiency_pct":53,"lane":2,"lane_role":2,"is_roaming":false,"purchase_time":{"circlet":-89},"first_purchase_time":{"circlet":-89},"item_win":{"circlet":1},"item_usage":{"circlet":1},"purchase_tpscroll":14,"actions_per_min":171,"life_state_dead":86,"solo_competitive_rank":7188,"cosmetics":[{"item_id":6671,"name":"Merry Wanderer's Brush","prefab":"wearable","creation_date":"2014-05-16T00:00:00.000Z","image_inventory":"econ/items/puck/merry_wanderers_brush_tail/merry_wanderers_brush_tail","image_path":"icons/econ/items/puck/merry_wanderers_brush_tail/merry_wanderers_brush_tail.5a6d5cf348f802896085dd52d48e092a52a021c2.png","item_description":"#DOTA_Item_Desc_Merry_Wanderers_Brush","item_name":"#DOTA_Item_Merry_Wanderers_Brush","item_rarity":"immortal","item_type_name":"#DOTA_WearableType_Tail","used_by_heroes":"npc_dota_hero_puck"}],"benchmarks":{"gold_per_min":{"raw":462,"pct":0.6323529411764706},"xp_per_min":{"raw":557,"pct":0.4411764705882353},"kills_per_min":{"raw":0.08068130883012102,"pct":0.17647058823529413},"last_hits_per_min":{"raw":5.970416853428955,"pct":0.9558823529411765},"hero_damage_per_min":{"raw":409.69968623935455,"pct":0.2647058823529412},"hero_healing_per_min":{"raw":0,"pct":1},"tower_damage":{"raw":1355,"pct":0.6323529411764706}}}],"patch":26,"region":2,"all_word_counts":{"glhf":2},"my_word_counts":{"glhf":2},"throw":5685,"loss":13861,"replay_url":"http://replay123.valve.net/570/3559037317_897752925.dem.bz2"}`)
	})

	expected := Match{
		MatchID:               3559037317,
		BarracksStatusDire:    48,
		BarracksStatusRadiant: 63,
		Chat: []Chat{
			Chat{
				Time:       -41,
				Type:       "chat",
				Unit:       "k y le",
				Key:        "glhf",
				Slot:       5,
				PlayerSlot: 128,
			},
		},
		Cluster: 123,
		Cosmetics: map[string]int{
			"5242": 3,
		},
		DireScore:      15,
		Duration:       2231,
		Engine:         1,
		FirstBloodTime: 120,
		GameMode:       2,
		HumanPlayers:   10,
		LeagueID:       5627,
		LobbyType:      1,
		MatchSeqNum:    3095917221,
		NegativeVotes:  12,
		Objectives: []Objective{
			Objective{
				Time: 31,
				Type: "CHAT_MESSAGE_COURIER_LOST",
				Team: 2,
			},
		},
		PicksBans: []PickBans{
			PickBans{
				IsPick:  false,
				HeroID:  60,
				Team:    1,
				Order:   0,
				Ord:     0,
				MatchID: 3559037317,
			},
		},
		PositiveVotes: 112,
		RadiantGoldAdv: []int{
			0,
		},
		RadiantScore: 27,
		RadiantWin:   true,
		RadiantXpAdv: []int{
			0,
		},
		StartTime: 1510535111,
		Teamfights: []Teamfights{
			Teamfights{
				Start:     266,
				End:       306,
				LastDeath: 291,
				Deaths:    3,
				Players: []TeamfightPlayers{
					TeamfightPlayers{
						AbilityUses: map[string]int{
							"puck_illusory_orb":   2,
							"puck_ethereal_jaunt": 1,
						},
						ItemUses: map[string]int{
							"tpscroll": 1,
							"flask":    1,
						},
						Deaths:    0,
						Buybacks:  0,
						Damage:    225,
						Healing:   295,
						GoldDelta: 224,
						XpDelta:   311,
						XpStart:   1570,
						XpEnd:     1881,
					},
				},
			},
		},
		TowerStatusDire:    384,
		TowerStatusRadiant: 1830,
		Version:            20,
		ReplaySalt:         897752925,
		SeriesID:           177752,
		SeriesType:         1,
		League: MatchLeague{
			LeagueID: 5627,
			Ticket:   "econ/leagues/subscriptions_dreamleague_season_8",
			Banner:   "econ/leagues/subscriptions_dreamleague_season_8_ingame",
			Tier:     "premium",
			Name:     "DreamLeague season 8",
		},
		RadiantTeam: MatchTeam{
			TeamID:  39,
			Name:    "Evil Geniuses",
			Tag:     "EG",
			LogoURL: "http://cloud-3.steamusercontent.com/ugc/142255738559146189/495FE7DC43BFAE03C5237446AB526888CA81827E/",
		},
		DireTeam: MatchTeam{
			TeamID:  3,
			Name:    "compLexity Gaming",
			Tag:     "coL",
			LogoURL: "http://cloud-3.steamusercontent.com/ugc/398960342836468051/26C1C5F6C3082FF2A1B706435F7DC0FCF5F5D27F/",
		},
		Players: []MatchPlayer{
			MatchPlayer{
				MatchID:    3559037317,
				PlayerSlot: 0,
				AbilityUpgradesArr: []int{
					5069,
				},
				AbilityUses: map[string]int{
					"puck_dream_coil": 9,
				},
				AccountID: 111620041,
				Actions: map[string]int{
					"1": 4247,
				},
				Assists:   14,
				Backpack0: 41,
				Backpack1: 46,
				Backpack2: 0,
				BuybackLog: []BuybackLog{
					BuybackLog{
						Time:       2101,
						Slot:       0,
						Type:       "buyback_log",
						PlayerSlot: 0,
					},
				},
				CampsStacked:  2,
				CreepsStacked: 5,
				Damage: map[string]int{
					"npc_dota_creep_badguys_melee": 56322,
				},
				DamageInflictor: map[string]int{
					"cyclone": 39,
				},
				DamageInflictorReceived: map[string]int{
					"jakiro_liquid_fire": 71,
				},
				DamageTaken: map[string]int{
					"npc_dota_badguys_tower4": 1101,
				},
				Deaths: 3,
				Denies: 8,
				DnT: []int{
					8,
				},
				FirstbloodClaimed: 0,
				Gold:              3339,
				GoldPerMin:        462,
				GoldReasons: map[string]int{
					"0": 1378,
				},
				GoldSpent: 13340,
				GoldT: []int{
					0,
				},
				HeroDamage:  15234,
				HeroHealing: 0,
				HeroHits: map[string]int{
					"cyclone": 1,
				},
				HeroID: 13,
				Item0:  1,
				Item1:  63,
				Item2:  100,
				Item3:  36,
				Item4:  0,
				Item5:  108,
				ItemUses: map[string]int{
					"cheese": 1,
				},
				Killed: map[string]int{
					"npc_dota_creep_badguys_melee": 120,
				},
				KilledBy: map[string]int{
					"npc_dota_hero_skeleton_king": 1,
				},
				Kills: 3,
				KillsLog: []Log{
					Log{
						Time: 1218,
						Key:  "npc_dota_hero_jakiro",
					},
				},
				LanePos: map[string]map[string]int{
					"74": {
						"74": 2,
					},
				},
				LastHits:     222,
				LeaverStatus: 0,
				Level:        22,
				LhT: []int{
					0,
				},
				LifeState: map[string]int{
					"0": 2291,
				},
				MaxHeroHit: MaxHeroHit{
					Type:       "max_hero_hit",
					Time:       1895,
					Max:        true,
					Inflictor:  "puck_waning_rift",
					Unit:       "npc_dota_hero_puck",
					Key:        "npc_dota_hero_skeleton_king",
					Value:      242,
					Slot:       0,
					PlayerSlot: 0,
				},
				Obs: map[string]map[string]int{
					"160": {
						"86": 1,
					},
				},
				ObsLeftLog: []ObsLog{
					ObsLog{
						Time:       319,
						Type:       "obs_left_log",
						Key:        "[160, 86]",
						Slot:       0,
						X:          160,
						Y:          86,
						Z:          130,
						Entityleft: true,
						Ehandle:    13091570,
						PlayerSlot: 0,
					},
				},
				ObsLog: []ObsLog{
					ObsLog{
						Time:       -47,
						Type:       "obs_log",
						Key:        "[160, 86]",
						Slot:       0,
						X:          160,
						Y:          86,
						Z:          130,
						Entityleft: false,
						Ehandle:    13091570,
						PlayerSlot: 0,
					},
				},
				ObsPlaced: 1,
				PartyID:   0,
				PartySize: 10,
				Pings:     37,
				PredVict:  false,
				Purchase: map[string]int{
					"ultimate_scepter": 1,
				},
				PurchaseLog: []Log{
					Log{
						Time: -89,
						Key:  "circlet",
					},
				},
				Randomed:      false,
				RoshansKilled: 0,
				RunePickups:   11,
				Runes: map[string]int{
					"1": 1,
				},
				RunesLog: []RunesLog{
					RunesLog{
						Time: 659,
						Key:  5,
					},
				},
				Sen: map[string]map[string]int{
					"108": {
						"122": 1,
					},
				},
				SenLeftLog: []ObsLog{
					ObsLog{
						Time:       1580,
						Type:       "sen_left_log",
						Key:        "[108, 122]",
						Slot:       0,
						X:          108,
						Y:          122,
						Z:          130,
						Entityleft: true,
						Ehandle:    10011357,
						PlayerSlot: 0,
					},
				},
				SenLog: []ObsLog{
					ObsLog{
						Time:       1334,
						Type:       "sen_log",
						Key:        "[108, 122]",
						Slot:       0,
						X:          108,
						Y:          122,
						Z:          130,
						Entityleft: false,
						Ehandle:    10011357,
						PlayerSlot: 0,
					},
				},
				SenPlaced: 1,
				Stuns:     31.331408,
				TeamfightParticipation: 0.6296296,
				Times: []int{
					0,
				},
				TowerDamage: 1355,
				XpPerMin:    557,
				XpReasons: map[string]int{
					"0": 1069,
					"1": 6912,
					"2": 12440,
					"3": 313,
				},
				XpT: []int{
					0,
				},
				Personaname:       "t.Danger",
				Name:              "SumaiL",
				RadiantWin:        true,
				StartTime:         1510535111,
				Duration:          2231,
				Cluster:           123,
				LobbyType:         1,
				GameMode:          2,
				Patch:             26,
				Region:            2,
				IsRadiant:         true,
				Win:               1,
				Lose:              0,
				TotalGold:         17178,
				TotalXp:           20711,
				KillsPerMin:       0.08068130883012103,
				Kda:               4,
				Abandons:          0,
				NeutralKills:      54,
				TowerKills:        0,
				CourierKills:      0,
				LaneKills:         166,
				HeroKills:         5,
				ObserverKills:     2,
				SentryKills:       0,
				RoshanKills:       0,
				NecronomiconKills: 0,
				AncientKills:      0,
				BuybackCount:      1,
				ObserverUses:      0,
				SentryUses:        1,
				LaneEfficiency:    0.5364970842549769,
				LaneEfficiencyPct: 53,
				Lane:              2,
				LaneRole:          2,
				IsRoaming:         false,
				PurchaseTime: map[string]int{
					"circlet": -89,
				},
				FirstPurchaseTime: map[string]int{
					"circlet": -89,
				},
				ItemWin: map[string]int{
					"circlet": 1,
				},
				ItemUsage: map[string]int{
					"circlet": 1,
				},
				PurchaseTpscroll:    14,
				ActionsPerMin:       171,
				LifeStateDead:       86,
				SoloCompetitiveRank: 7188,
				Cosmetics: []Cosmetics{
					Cosmetics{
						ItemID:          6671,
						Name:            "Merry Wanderer's Brush",
						Prefab:          "wearable",
						CreationDate:    "2014-05-16T00:00:00.000Z",
						ImageInventory:  "econ/items/puck/merry_wanderers_brush_tail/merry_wanderers_brush_tail",
						ImagePath:       "icons/econ/items/puck/merry_wanderers_brush_tail/merry_wanderers_brush_tail.5a6d5cf348f802896085dd52d48e092a52a021c2.png",
						ItemDescription: "#DOTA_Item_Desc_Merry_Wanderers_Brush",
						ItemName:        "#DOTA_Item_Merry_Wanderers_Brush",
						ItemRarity:      "immortal",
						ItemTypeName:    "#DOTA_WearableType_Tail",
						UsedByHeroes:    "npc_dota_hero_puck",
					},
				},
				Benchmarks: Benchmarks{
					GoldPerMin: RawPCT{
						Raw: 462,
						Pct: 0.6323529411764706,
					},
					XpPerMin: RawPCT{
						Raw: 557,
						Pct: 0.4411764705882353,
					},
					KillsPerMin: RawPCT{
						Raw: 0.08068130883012102,
						Pct: 0.17647058823529413,
					},
					LastHitsPerMin: RawPCT{
						Raw: 5.970416853428955,
						Pct: 0.9558823529411765,
					},
					HeroDamagePerMin: RawPCT{
						Raw: 409.69968623935455,
						Pct: 0.2647058823529412,
					},
					HeroHealingPerMin: RawPCT{
						Raw: 0,
						Pct: 1,
					},
					TowerDamage: RawPCT{
						Raw: 1355,
						Pct: 0.6323529411764706,
					},
				},
			},
		},
		Patch:  26,
		Region: 2,
		AllWordCounts: map[string]int{
			"glhf": 2,
		},
		MyWordCounts: map[string]int{
			"glhf": 2,
		},
		Throw:     5685,
		Loss:      13861,
		ReplayURL: "http://replay123.valve.net/570/3559037317_897752925.dem.bz2",
	}

	client := NewClient(httpClient)
	match, _, err := client.MatchService.Match(3559037317)
	assert.Nil(t, err)
	assert.Equal(t, expected, match)
}
