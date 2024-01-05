package db

import (
	"path"
	"reflect"
	"testing"
)

func TestGetBannedChannelDB(t *testing.T) {
	expected := []Channel{
		{Id: 1, Title: "channel1"},
		{Id: 2, Title: "channel2"},
		{Id: 3, Title: "channel3"},
	}

	db := BannedChannelDB{
		database[Channel]{
			data: expected,
		},
	}

	channels := db.Get()

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}
}

func TestBan(t *testing.T) {
	db := BannedChannelDB{
		database[Channel]{
			filename: path.Join(t.TempDir(), "banned-channels.json"),
			data: []Channel{
				{Id: 1, Title: "channel1"},
				{Id: 2, Title: "channel2"},
			},
		},
	}

	ch := Channel{Id: 3, Title: "channel3"}
	db.Ban(ch)

	channels := db.Get()
	expected := []Channel{
		{Id: 1, Title: "channel1"},
		{Id: 2, Title: "channel2"},
		{Id: 3, Title: "channel3"},
	}

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}

	// Test that duplicate channels with same ID are not added...
	db.Ban(ch)
	channels = db.Get()
	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}

	// ...but channels with same titles are added
	ch2 := Channel{Id: 4, Title: "channel3"}
	db.Ban(ch2)
	expected = []Channel{
		{Id: 1, Title: "channel1"},
		{Id: 2, Title: "channel2"},
		{Id: 3, Title: "channel3"},
		{Id: 4, Title: "channel3"},
	}
	channels = db.Get()
	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}
}

func TestIsBanned(t *testing.T) {
	db := BannedChannelDB{
		database[Channel]{
			data: []Channel{
				{Id: 1, Title: "channel1"},
				{Id: 2, Title: "channel2"},
				{Id: 3, Title: "channel3"},
			},
		},
	}

	if !db.IsBanned(1) {
		t.Errorf("Expected channel 1 to be banned")
	}

	if db.IsBanned(4) {
		t.Errorf("Expected channel 4 to not be banned")
	}
}

func TestClear(t *testing.T) {
	db := BannedChannelDB{
		database[Channel]{
			filename: path.Join(t.TempDir(), "banned-channels.json"),
			data: []Channel{
				{Id: 1, Title: "channel1"},
				{Id: 2, Title: "channel2"},
				{Id: 3, Title: "channel3"},
			},
		},
	}

	db.Clear()

	channels := db.Get()
	expected := []Channel{}

	if !reflect.DeepEqual(channels, expected) {
		t.Errorf("Expected channels %v, got %v", expected, channels)
	}
}
