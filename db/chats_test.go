package db

import (
	"path"
	"reflect"
	"testing"
)

func TestChatDBGet(t *testing.T) {
	expected := []Chat{
		{Id: 1, Title: "chat1"},
		{Id: 2, Title: "chat2"},
		{Id: 3, Title: "chat3"},
	}

	db := ChatDB{
		database[Chat]{
			data: expected,
		},
	}

	chats := db.Get()

	if !reflect.DeepEqual(chats, expected) {
		t.Errorf("Expected chats %v, got %v", expected, chats)
	}
}

func TestChatDBGetIdByOrdinal(t *testing.T) {
	db := ChatDB{
		database[Chat]{
			data: []Chat{
				{Id: 1, Title: "chat1"},
				{Id: 2, Title: "chat2"},
				{Id: 3, Title: "chat3"},
			},
		},
	}

	tests := []struct {
		num          int
		expectedId   int64
		expectedBool bool
	}{
		{1, 1, true},
		{2, 2, true},
		{3, 3, true},
		{4, 0, false},
		{0, 0, false},
	}

	for _, test := range tests {
		id, ok := db.GetIdByOrdinal(test.num)
		if id != test.expectedId || ok != test.expectedBool {
			t.Errorf("Expected id = %d, ok = %t for #%d, but got id = %d, ok = %t",
				test.expectedId, test.expectedBool, test.num, id, ok)
		}
	}
}

func TestChatDBAdd(t *testing.T) {
	db := ChatDB{
		database[Chat]{
			filename: path.Join(t.TempDir(), "test-chats.json"),
			data: []Chat{
				{Id: 1, Title: "chat1"},
				{Id: 2, Title: "chat2"},
				{Id: 3, Title: "chat3"},
			},
		},
	}

	newChat := Chat{Id: 4, Title: "chat4"}
	expected := []Chat{
		{Id: 1, Title: "chat1"},
		{Id: 2, Title: "chat2"},
		{Id: 3, Title: "chat3"},
		{Id: 4, Title: "chat4"},
	}

	db.Add(newChat)

	chats := db.Get()

	if !reflect.DeepEqual(chats, expected) {
		t.Errorf("Expected chats %v, got %v", expected, chats)
	}

	// Test that duplicate chats with same ID are not added
	db.Add(newChat)
	chats = db.Get()
	if !reflect.DeepEqual(chats, expected) {
		t.Errorf("Expected chats %v, got %v", expected, chats)
	}
}
