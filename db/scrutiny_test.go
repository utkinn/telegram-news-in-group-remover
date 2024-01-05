package db

import (
	"path"
	"testing"
)

func TestScrutinyDBAdd(t *testing.T) {
	database := ScrutinyDB{
		database[string]{
			filename: path.Join(t.TempDir(), "scrutiny.json"),
			data:     []string{},
		},
	}

	userName := "testUserName"

	database.Add(userName)

	if !database.IsUnderScrutiny(userName) {
		t.Errorf("Expected username %s to be under scrutiny", userName)
	}

	// Test duplicates
	database.Add(userName)

	if len(database.data) != 1 {
		t.Errorf("Expected only one entry in the database")
	}

	// Test "@" trimming
	database.Add("@" + userName)

	if len(database.data) != 1 {
		t.Errorf("Expected only one entry in the database")
	}
}

func TestScrutinyDBRemove(t *testing.T) {
	database := ScrutinyDB{
		database[string]{
			filename: path.Join(t.TempDir(), "scrutiny.json"),
			data:     []string{"testUserName", "anotherUserName", "yetAnotherUserName"},
		},
	}

	t.Run("Remove existing username", func(t *testing.T) {
		userName := "testUserName"
		removed := database.Remove(userName)
		if !removed {
			t.Errorf("Expected return value of db.Remove(%#v) to be true", userName)
		}
		if database.IsUnderScrutiny(userName) {
			t.Errorf("Expected userName %s to be removed from scrutiny", userName)
		}
	})

	t.Run("Remove non-existing username", func(t *testing.T) {
		userName := "nonExistingUserName"
		removed := database.Remove(userName)
		if removed {
			t.Errorf("Expected return value of db.Remove(%#v) to be false", userName)
		}
	})

	t.Run("Remove username with '@' prefix", func(t *testing.T) {
		userName := "@anotherUserName"
		removed := database.Remove(userName)
		if !removed {
			t.Errorf("Expected return value of db.Remove(%#v) to be true", userName)
		}
		if database.IsUnderScrutiny(userName) {
			t.Errorf("Expected userName %s to be removed from scrutiny", userName)
		}
	})
}
