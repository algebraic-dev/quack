package internal

import (
	"testing"
)

func TestLogParsingSuccessful(t *testing.T) {
	event, err := Validate("  21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT")

	if err != nil {
		t.Error(err)
	}

	if event.Hour != 21 {
		t.Error("Expected hour to be 21")
	}

	if event.Minutes != 42 {
		t.Error("Expected minutes to be 42")
	}

	if event.Type != "Kill" {
		t.Error("Expected type to be Kill")
	}

	if event.Description != "1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT" {
		t.Error("Expected description to be <world> killed Isgalamido by MOD_TRIGGER_HURT")
	}
}
