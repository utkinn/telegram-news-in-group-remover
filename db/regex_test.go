package db

import (
	"path"
	"reflect"
	"testing"
)

func TestBannedRegexDBBan(t *testing.T) {
	database := BannedRegexDB{database[string]{filename: path.Join(t.TempDir(), "test-banned-regex.json")}}

	err := database.Ban(`\d+`)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(database.data) != 1 {
		t.Fatalf("Expected 1 regex, got %v", len(database.data))
	}

	err = database.Ban(`invalid regex: [a-z+`)
	if err == nil {
		t.Fatal("Expected error for invalid regex")
	}
}

func TestBannedRegexDBGet(t *testing.T) {
	expected := []string{"regex1", "regex2", "regex3"}
	database := BannedRegexDB{
		database[string]{
			filename: path.Join(t.TempDir(), "test-banned-regex.json"),
			data:     expected,
		},
	}

	actual := database.Get()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestBannedRegexDBUnban(t *testing.T) {
	database := BannedRegexDB{
		database[string]{
			filename: path.Join(t.TempDir(), "test-banned-regex.json"),
			data:     []string{"regex1", "regex2", "regex3"},
		},
	}

	database.Unban("regex2")

	expected := []string{"regex1", "regex3"}
	actual := database.Get()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
