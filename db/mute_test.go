//nolint:exhaustruct
package db

import (
	"path"
	"reflect"
	"testing"
	"time"
)

type fakeClock struct{}

func (fakeClock) Now() time.Time { return time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC) }

func TestMuteDBUnmuteUser(t *testing.T) {
	database := MuteDB{
		database[mute]{
			filename: path.Join(t.TempDir(), "mute.json"),
			data: []mute{
				{UserName: "user1"},
				{UserName: "user2"},
				{UserName: "user3"},
			},
		},
		fakeClock{},
	}

	database.UnmuteUser("user2")

	expected := []mute{
		{UserName: "user1"},
		{UserName: "user3"},
	}

	// Check if the data in the MuteDB instance matches the expected result
	if !reflect.DeepEqual(database.data, expected) {
		t.Errorf("Expected data %v, got %v", expected, database.data)
	}
}

func TestMuteDBGetStatusForUser(t *testing.T) {
	nextHour := time.Now().Add(time.Hour)
	database := MuteDB{
		database[mute]{
			filename: path.Join(t.TempDir(), "mute.json"),
			data: []mute{
				{UserName: "user1", IsAnnounced: true, EndAt: nextHour},
				{UserName: "user2", IsAnnounced: false, EndAt: nextHour},
				{UserName: "expired", IsAnnounced: false, EndAt: time.Now().Add(-time.Hour)},
			},
		},
		fakeClock{},
	}

	tests := []struct {
		userName          string
		expectedMuted     bool
		expectedAnnounced bool
	}{
		{"user1", true, true},
		{"user2", true, false},
		{"user?", false, false},
		{"expired", false, false},
	}

	for _, test := range tests {
		muted, announced := database.GetStatusForUser(test.userName)
		if muted != test.expectedMuted || announced != test.expectedAnnounced {
			t.Errorf("For user %s, expected muted=%v, announced=%v; got muted=%v, announced=%v",
				test.userName, test.expectedMuted, test.expectedAnnounced, muted, announced)
		}
	}

	// Check that the expired mute was removed
	expected := []mute{
		{UserName: "user1", IsAnnounced: true, EndAt: nextHour},
		{UserName: "user2", IsAnnounced: false, EndAt: nextHour},
	}
	if !reflect.DeepEqual(database.data, expected) {
		t.Errorf("Expected data %+v, got %+v", expected, database.data)
	}
}

func TestMuteDBMarkMuteAnnounced(t *testing.T) {
	database := MuteDB{
		database[mute]{
			filename: path.Join(t.TempDir(), "mute.json"),
			data: []mute{
				{UserName: "user1", IsAnnounced: false},
			},
		},
		fakeClock{},
	}

	database.MarkMuteAnnounced("user1")

	expected := []mute{
		{UserName: "user1", IsAnnounced: true},
	}

	// Check if the data in the MuteDB instance matches the expected result
	if !reflect.DeepEqual(database.data, expected) {
		t.Errorf("Expected data %+v, got %+v", expected, database.data)
	}
}

func TestMuteDBMuteUser(t *testing.T) {
	clock := fakeClock{}
	database := MuteDB{
		database[mute]{
			filename: path.Join(t.TempDir(), "mute.json"),
			data:     []mute{},
		},
		clock,
	}

	userName := "user"
	duration := time.Hour

	database.MuteUser(userName, duration)

	expected := []mute{
		{UserName: userName, StartAt: clock.Now(), EndAt: clock.Now().Add(duration)},
	}

	// Check if the data in the MuteDB instance matches the expected result
	if !reflect.DeepEqual(database.data, expected) {
		t.Errorf("Expected data %+v, got %+v", expected, database.data)
	}
}
