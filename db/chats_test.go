package db

import (
	"path"
	"reflect"
	"testing"
)

func TestChatDBGet(t *testing.T) {
	expected := []Chat{
		{ID: 1, Title: "chat1"},
		{ID: 2, Title: "chat2"},
		{ID: 3, Title: "chat3"},
	}

	database := ChatDB{
		database[Chat]{
			data: expected,
		},
	}

	if chats := database.Get(); !reflect.DeepEqual(chats, expected) {
		t.Errorf("Expected chats %v, got %v", expected, chats)
	}
}

func TestChatDBGetIdByOrdinal(t *testing.T) {
	database := ChatDB{
		database[Chat]{
			data: []Chat{
				{ID: 1, Title: "chat1"},
				{ID: 2, Title: "chat2"},
				{ID: 3, Title: "chat3"},
			},
		},
	}

	tests := []struct {
		num          int
		expectedID   int64
		expectedBool bool
	}{
		{1, 1, true},
		{2, 2, true},
		{3, 3, true},
		{4, 0, false},
		{0, 0, false},
	}

	for _, test := range tests {
		id, ok := database.GetIDByOrdinal(test.num)
		if id != test.expectedID || ok != test.expectedBool {
			t.Errorf("Expected id = %d, ok = %t for #%d, but got id = %d, ok = %t",
				test.expectedID, test.expectedBool, test.num, id, ok)
		}
	}
}

func TestChatDBAdd(t *testing.T) {
	database := ChatDB{
		database[Chat]{
			filename: path.Join(t.TempDir(), "test-chats.json"),
			data: []Chat{
				{ID: 1, Title: "chat1"},
				{ID: 2, Title: "chat2"},
				{ID: 3, Title: "chat3"},
			},
		},
	}

	newChat := Chat{ID: 4, Title: "chat4"}
	expected := []Chat{
		{ID: 1, Title: "chat1"},
		{ID: 2, Title: "chat2"},
		{ID: 3, Title: "chat3"},
		{ID: 4, Title: "chat4"},
	}

	database.Add(newChat)

	chats := database.Get()

	if !reflect.DeepEqual(chats, expected) {
		t.Errorf("Expected chats %v, got %v", expected, chats)
	}

	// Test that duplicate chats with same ID are not added
	database.Add(newChat)
	chats = database.Get()

	if !reflect.DeepEqual(chats, expected) {
		t.Errorf("Expected chats %v, got %v", expected, chats)
	}
}
