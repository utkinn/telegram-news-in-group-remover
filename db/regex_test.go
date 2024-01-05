package db

import (
	"path"
	"reflect"
	"testing"
)

func TestBannedRegexDBBan(t *testing.T) {
	db := BannedRegexDB{database[string]{filename: path.Join(t.TempDir(), "test-banned-regex.json")}}

	err := db.Ban(`\d+`)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(db.data) != 1 {
		t.Fatalf("Expected 1 regex, got %v", len(db.data))
	}

	err = db.Ban(`invalid regex: [a-z+`)
	if err == nil {
		t.Fatal("Expected error for invalid regex")
	}
}

func TestBannedRegexDBGet(t *testing.T) {
	expected := []string{"regex1", "regex2", "regex3"}
	db := BannedRegexDB{
		database[string]{
			filename: path.Join(t.TempDir(), "test-banned-regex.json"),
			data:     expected,
		},
	}

	actual := db.Get()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestBannedRegexDBUnban(t *testing.T) {
	db := BannedRegexDB{
		database[string]{
			filename: path.Join(t.TempDir(), "test-banned-regex.json"),
			data:     []string{"regex1", "regex2", "regex3"},
		},
	}

	db.Unban("regex2")

	expected := []string{"regex1", "regex3"}
	actual := db.Get()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
