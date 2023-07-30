package internal

import (
	"strings"
)

type MatchAnalytics struct {
	TotalKills int

	// Poor man's HashSet
	Players map[string]bool
	
	/// Kills scoreboard for players and the world
	Kills map[string]int

	/// Kills grouped by the reason
	KillsByMeans map[string]int
}

// This function parses a list of logs into a list of events that can be used to extract information.
func ToEvents(logs string) ([]Event, error) {
	lines := strings.Split(logs, "\n")
	
	var events []Event

	for _, line := range lines {
		event, err := Validate(line)

		if err != nil {
			continue
		}

		result, err := Parse(event)

		if err != nil {
			return nil, err

		}

		events = append(events, result)
	}

	return events, nil
}

func onKill(event KillEvent, analytics *MatchAnalytics) {
	if event.Killer == "<world>" || event.Killer == event.Victim {
		analytics.Kills[event.Victim] -= 1
	} else  {
		analytics.Kills[event.Killer] += 1	
	}
	analytics.KillsByMeans[event.Cause] += 1
	analytics.TotalKills += 1
}

func onInitGame(event InitGameEvent, analytics *MatchAnalytics) {
	*analytics = MatchAnalytics {
		Kills: make(map[string]int),
		Players: make(map[string]bool),
		KillsByMeans: make(map[string]int),
	}
}

func onClientUserinfoChanged(event ClientUserinfoChangedEvent, analytics *MatchAnalytics) {
	analytics.Players[event.Player] = true
	analytics.Kills[event.Player] = 0
}

// This function collects the data from the events and returns a list of analytics. This function
// is following some rules for the killing scoreboard:
//
//  - If the killer is the world, the victim loses a point.
//  - If the killer is a player, the killer gains a point.
//  - If the killer is the victim, the killer loses a point.
//
func CollectData(logs []Event) []MatchAnalytics {
	var result []MatchAnalytics
	var analytics MatchAnalytics
	
	first := true

	for _, event := range logs {
		switch t := event.(type) {
			case ClientUserinfoChangedEvent: onClientUserinfoChanged(t, &analytics)
			case KillEvent: onKill(t, &analytics)
			case InitGameEvent:
				if !first {
					result = append(result, analytics)
				}
				first = false
				onInitGame(t, &analytics)
			default:
				continue
		}
	}

	return result
}