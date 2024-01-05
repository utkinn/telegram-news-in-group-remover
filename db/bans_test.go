package db

import (
	"path"
	"reflect"
	"testing"
)

func TestGetBannedChannelDB(t *testing.T) {
	expected := []Channel{
		{ID: 1, Title: "channel1"},
		{ID: 2, Title: "channel2"},
		{ID: 3, Title: "channel3"},
	}

	database := BannedChannelDB{
		database[Channel]{
			data: expected,
		},
	}

	channels := database.Get()

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}
}

func TestBan(t *testing.T) {
	database := BannedChannelDB{
		database[Channel]{
			filename: path.Join(t.TempDir(), "banned-channels.json"),
			data: []Channel{
				{ID: 1, Title: "channel1"},
				{ID: 2, Title: "channel2"},
			},
		},
	}

	callback := Channel{ID: 3, Title: "channel3"}
	database.Ban(callback)

	channels := database.Get()
	expected := []Channel{
		{ID: 1, Title: "channel1"},
		{ID: 2, Title: "channel2"},
		{ID: 3, Title: "channel3"},
	}

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}

	// Test that duplicate channels with same ID are not added...
	database.Ban(callback)
	channels = database.Get()

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}

	// ...but channels with same titles are added
	ch2 := Channel{ID: 4, Title: "channel3"}
	database.Ban(ch2)

	expected = []Channel{
		{ID: 1, Title: "channel1"},
		{ID: 2, Title: "channel2"},
		{ID: 3, Title: "channel3"},
		{ID: 4, Title: "channel3"},
	}
	channels = database.Get()

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}
}

func TestIsBanned(t *testing.T) {
	database := BannedChannelDB{
		database[Channel]{
			data: []Channel{
				{ID: 1, Title: "channel1"},
				{ID: 2, Title: "channel2"},
				{ID: 3, Title: "channel3"},
			},
		},
	}

	if !database.IsBanned(1) {
		t.Errorf("Expected channel 1 to be banned")
	}

	if database.IsBanned(4) {
		t.Errorf("Expected channel 4 to not be banned")
	}
}

func TestClear(t *testing.T) {
	database := BannedChannelDB{
		database[Channel]{
			filename: path.Join(t.TempDir(), "banned-channels.json"),
			data: []Channel{
				{ID: 1, Title: "channel1"},
				{ID: 2, Title: "channel2"},
				{ID: 3, Title: "channel3"},
			},
		},
	}

	database.Clear()

	channels := database.Get()
	expected := []Channel{}

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}
}
