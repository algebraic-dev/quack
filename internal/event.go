// This module is useful for parsing and extracting information from the log files.

package internal

import (
	"errors"
	"regexp"
	"strconv"
)

/// A single quake event that is parsed from the log.
type Event struct {
	// The event time in hours
	Hour int

	// The minutes of the event
	Minutes int

	// The event type
	Type string

	// The event description
	Description string
}

// The regex for parsing a single event in the log.
var eventRegex = regexp.MustCompile(`\s*(\d+):(\d+)\s([a-zA-Z]+):\s*([^\n]*)$`)

// This function validates a single event and returns an Event struct. It's a first pass in order to
// validate the log file and get some events without too many details. 
func Validate(log string) (Event, error) {
	matches := eventRegex.FindStringSubmatch(log)

	if eventRegex.NumSubexp() != 4 {
		return Event {}, errors.New("invalid event")
	}

	hour, _ := strconv.ParseInt(matches[1], 10, 32)
	minutes, _ := strconv.ParseInt(matches[2], 10, 32)

	event := Event { 
		Hour: int(hour),
		Minutes: int(minutes),
		Type: matches[3],
		Description: matches[4],
	}

	return event, nil
}