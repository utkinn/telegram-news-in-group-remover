package db

import (
	"path"
	"reflect"
	"testing"
)

func TestSubscribe(t *testing.T) {
	db := AnnouncementSubscriptionDB{database[AnnouncementSubscription]{filename: path.Join(t.TempDir(), "test.json")}}

	db.Subscribe(123, "user1")
	if len(db.data) != 1 {
		t.Errorf("Expected 1 subscription, got %d", len(db.data))
	}

	// Subscribing again should not add a new subscription
	db.Subscribe(123, "user1")
	if len(db.data) != 1 {
		t.Errorf("Expected 1 subscription, got %d", len(db.data))
	}

	db.Subscribe(456, "user2")
	if len(db.data) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(db.data))
	}
}

func TestUnsubscribe(t *testing.T) {
	db := AnnouncementSubscriptionDB{
		database[AnnouncementSubscription]{
			filename: path.Join(t.TempDir(), "test.json"),
			data: []AnnouncementSubscription{
				{ChatId: 123, UserName: "user1"},
				{ChatId: 456, UserName: "user2"},
				{ChatId: 789, UserName: "user3"},
			},
		},
	}

	db.Unsubscribe(456)

	expected := []AnnouncementSubscription{
		{ChatId: 123, UserName: "user1"},
		{ChatId: 789, UserName: "user3"},
	}
	if !reflect.DeepEqual(db.data, expected) {
		t.Errorf("Expected data %v, got %v", expected, db.data)
	}
}

func TestUnsubscribeMissing(t *testing.T) {
	db := AnnouncementSubscriptionDB{
		database[AnnouncementSubscription]{
			filename: path.Join(t.TempDir(), "test.json"),
			data: []AnnouncementSubscription{
				{ChatId: 123, UserName: "user1"},
				{ChatId: 456, UserName: "user2"},
				{ChatId: 789, UserName: "user3"},
			},
		},
	}

	db.Unsubscribe(404)

	if len(db.data) != 3 {
		t.Errorf("Expected 3 subscriptions, got %d", len(db.data))
	}
}

func TestGetChatIdsOfSubscribedAdmins(t *testing.T) {
	db := AnnouncementSubscriptionDB{
		database[AnnouncementSubscription]{
			data: []AnnouncementSubscription{
				{ChatId: 123, UserName: "user1"},
				{ChatId: 456, UserName: "user2"},
				{ChatId: 789, UserName: "user3"},
			},
		},
	}

	ids := db.GetChatIdsOfSubscribedAdmins()

	expected := []int64{123, 456, 789}
	if !reflect.DeepEqual(ids, expected) {
		t.Errorf("Expected chat IDs %v, got %v", expected, ids)
	}
}
