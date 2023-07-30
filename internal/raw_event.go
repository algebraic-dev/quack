// This module is useful for parsing and extracting information from the log files.

package internal

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
)

/// A single quake event that is parsed from the log.
type RawEvent struct {
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
var eventRegex = regexp.MustCompile(`\s*(\d+):(\d+)\s([a-zA-Z]+):\s*([^\n]*)`)

// This function validates a single event and returns an RawEvent struct. It's a first pass in order to
// validate the log file and get some events without too many details. 
func Validate(log string) (RawEvent, error) {
	matches := eventRegex.FindStringSubmatch(log)

	if len(matches) != 5 {
		return RawEvent {}, errors.New(fmt.Sprintf("invalid event: '%s'", log))
	}

	hour, _ := strconv.ParseInt(matches[1], 10, 32)
	minutes, _ := strconv.ParseInt(matches[2], 10, 32)

	event := RawEvent { 
		Hour: int(hour),
		Minutes: int(minutes),
		Type: matches[3],
		Description: matches[4],
	}

	return event, nil
}