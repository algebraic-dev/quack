package internal

import (
	"strings"
	"errors"
)

type EventType int64

const (
	Kill EventType = iota
	InitGame
	ShutdownGame 
	Irrelevant
)

/// Interface that all events must implement. It's used to look like a algebraic data type.
type Event interface {
	GetEventType() EventType
}

type KillEvent struct {
	Killer string
	Victim string
	Cause string
}

func (e KillEvent) GetEventType() EventType {
	return Kill
}

type InitGameEvent struct {}

func (e InitGameEvent) GetEventType() EventType {
	return InitGame
}

type ShutdownGameEvent struct {}

func (e ShutdownGameEvent) GetEventType() EventType {
	return ShutdownGame
}

type IrrelevantEvent struct {}

func (e IrrelevantEvent) GetEventType() EventType {
	return Irrelevant
}


func parseKillEvent(description string) (KillEvent, error) {
	words := strings.Split(description, " ")

	var info [3]string

	/// This variable us used to know which string to add to in the info array.
	phase := 0

	for _, word := range words[3:] {
		switch word {
			case "killed":
				phase = 1
			case "by":
				phase = 2
			default:
				info[phase] += word + " "
		}
	}

	if phase != 2 {
		return KillEvent {}, errors.New("invalid kill event")
	}

	event := KillEvent {
		Killer: strings.Trim(info[0], " "),
		Victim: strings.Trim(info[1], " "),
		Cause: strings.Trim(info[2], " "),
	}

	return event, nil
}


// This is the main function of the module. It translates a raw event into a event that implements
// the Event interface.
func Parse(raw RawEvent) (Event, error) {
	switch raw.Type {
		case "Kill":
			return parseKillEvent(raw.Description)
		case "InitGame":
			return InitGameEvent {}, nil
		case "ShutdownGame":
			return ShutdownGameEvent {}, nil
		default:
			return IrrelevantEvent {}, nil
	}
}