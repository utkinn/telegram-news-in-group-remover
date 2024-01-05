package db

import (
	"path"
	"reflect"
	"testing"
)

func TestFilterToggleDBIsFilterEnabled(t *testing.T) {
	database := FilterToggleDB{
		database[filterToggle]{
			data: []filterToggle{
				{ID: "enabled", Enabled: true},
				{ID: "disabled", Enabled: false},
			},
		},
	}

	tests := []struct {
		id       string
		expected bool
	}{
		{"enabled", true},
		{"disabled", false},
		{"nonexisting", true}, // Test case for non-existing filter ID
	}

	for _, test := range tests {
		result := database.IsFilterEnabled(test.id)
		if result != test.expected {
			t.Errorf("Expected IsFilterEnabled(%q) to be %t, but got %t", test.id, test.expected, result)
		}
	}
}

func TestFilterToggleDBSetFilterEnabled(t *testing.T) {
	database := FilterToggleDB{
		database[filterToggle]{
			filename: path.Join(t.TempDir(), "test-filter-toggles.json"),
			data: []filterToggle{
				{ID: "enabled", Enabled: true},
				{ID: "disabled", Enabled: false},
			},
		},
	}

	tests := []struct {
		id               string
		enabled          bool
		expectedSnapshot []filterToggle
	}{
		{"enabled", false, []filterToggle{
			{ID: "enabled", Enabled: false},
			{ID: "disabled", Enabled: false},
		}},
		{"disabled", true, []filterToggle{
			{ID: "enabled", Enabled: false},
			{ID: "disabled", Enabled: true},
		}},
		{"nonexisting", true, []filterToggle{
			{ID: "enabled", Enabled: false},
			{ID: "disabled", Enabled: true},
			{ID: "nonexisting", Enabled: true},
		}},
	}

	for _, test := range tests {
		database.SetFilterEnabled(test.id, test.enabled)

		if !reflect.DeepEqual(database.data, test.expectedSnapshot) {
			t.Errorf("Expected db.data to be %v, but got %v", test.expectedSnapshot, database.data)
		}
	}
}
