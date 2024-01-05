package db

import (
	"path"
	"reflect"
	"testing"
)

func TestFilterToggleDBIsFilterEnabled(t *testing.T) {
	db := FilterToggleDB{
		database[filterToggle]{
			data: []filterToggle{
				{Id: "enabled", Enabled: true},
				{Id: "disabled", Enabled: false},
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
		result := db.IsFilterEnabled(test.id)
		if result != test.expected {
			t.Errorf("Expected IsFilterEnabled(%q) to be %t, but got %t", test.id, test.expected, result)
		}
	}
}

func TestFilterToggleDBSetFilterEnabled(t *testing.T) {
	db := FilterToggleDB{
		database[filterToggle]{
			filename: path.Join(t.TempDir(), "test-filter-toggles.json"),
			data: []filterToggle{
				{Id: "enabled", Enabled: true},
				{Id: "disabled", Enabled: false},
			},
		},
	}

	tests := []struct {
		id               string
		enabled          bool
		expectedSnapshot []filterToggle
	}{
		{"enabled", false, []filterToggle{
			{Id: "enabled", Enabled: false},
			{Id: "disabled", Enabled: false},
		}},
		{"disabled", true, []filterToggle{
			{Id: "enabled", Enabled: false},
			{Id: "disabled", Enabled: true},
		}},
		{"nonexisting", true, []filterToggle{
			{Id: "enabled", Enabled: false},
			{Id: "disabled", Enabled: true},
			{Id: "nonexisting", Enabled: true},
		}},
	}

	for _, test := range tests {
		db.SetFilterEnabled(test.id, test.enabled)
		if !reflect.DeepEqual(db.data, test.expectedSnapshot) {
			t.Errorf("Expected db.data to be %v, but got %v", test.expectedSnapshot, db.data)
		}
	}
}
