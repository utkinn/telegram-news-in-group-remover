package db

import (
	"path"
	"reflect"
	"testing"
)

func TestSubscribe(t *testing.T) {
	database := AnnouncementSubscriptionDB{
		database[AnnouncementSubscription]{
			filename: path.Join(t.TempDir(), "test.json"),
		},
	}

	database.Subscribe(123, "user1")

	if len(database.data) != 1 {
		t.Errorf("Expected 1 subscription, got %d", len(database.data))
	}

	// Subscribing again should not add a new subscription
	database.Subscribe(123, "user1")

	if len(database.data) != 1 {
		t.Errorf("Expected 1 subscription, got %d", len(database.data))
	}

	database.Subscribe(456, "user2")

	if len(database.data) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(database.data))
	}
}

func TestUnsubscribe(t *testing.T) {
	database := AnnouncementSubscriptionDB{
		database[AnnouncementSubscription]{
			filename: path.Join(t.TempDir(), "test.json"),
			data: []AnnouncementSubscription{
				{ChatID: 123, UserName: "user1"},
				{ChatID: 456, UserName: "user2"},
				{ChatID: 789, UserName: "user3"},
			},
		},
	}

	database.Unsubscribe(456)

	expected := []AnnouncementSubscription{
		{ChatID: 123, UserName: "user1"},
		{ChatID: 789, UserName: "user3"},
	}
	if !reflect.DeepEqual(database.data, expected) {
		t.Errorf("Expected data %v, got %v", expected, database.data)
	}
}

func TestUnsubscribeMissing(t *testing.T) {
	database := AnnouncementSubscriptionDB{
		database[AnnouncementSubscription]{
			filename: path.Join(t.TempDir(), "test.json"),
			data: []AnnouncementSubscription{
				{ChatID: 123, UserName: "user1"},
				{ChatID: 456, UserName: "user2"},
				{ChatID: 789, UserName: "user3"},
			},
		},
	}

	database.Unsubscribe(404)

	if len(database.data) != 3 {
		t.Errorf("Expected 3 subscriptions, got %d", len(database.data))
	}
}

func TestGetChatIdsOfSubscribedAdmins(t *testing.T) {
	database := AnnouncementSubscriptionDB{
		database[AnnouncementSubscription]{
			data: []AnnouncementSubscription{
				{ChatID: 123, UserName: "user1"},
				{ChatID: 456, UserName: "user2"},
				{ChatID: 789, UserName: "user3"},
			},
		},
	}

	ids := database.GetChatIDsOfSubscribedAdmins()

	expected := []int64{123, 456, 789}
	if !reflect.DeepEqual(ids, expected) {
		t.Errorf("Expected chat IDs %v, got %v", expected, ids)
	}
}
