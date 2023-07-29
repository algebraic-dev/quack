package internal

import (
	"testing"

)

func TestLogKillParsing(t *testing.T) {
	event, err := Validate("  21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT")

	if err != nil {
		t.Error(err)
		return
	}
	
	res, err := Parse(event)

	if err != nil {
		t.Error(err)
	}

	_, ok := res.(KillEvent)

	if !ok {
		t.Error("Expected event to be KillEvent")
	}
}

func TestLogInitGameParsing(t *testing.T) {
	event, err := Validate("  1:47 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\bot_minplayers\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0")

	if err != nil {
		t.Error(err)
		return
	}

	res, err := Parse(event)

	if err != nil {
		t.Error(err)
	}

	_, ok := res.(InitGameEvent)

	if !ok {
		t.Error("Expected event to be InitGameEvent")
	}
}