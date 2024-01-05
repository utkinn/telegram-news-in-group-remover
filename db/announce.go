package db

type AnnouncementSubscription struct {
	UserName string
	ChatID   int64 `json:"ChatId"`
}

type AnnouncementSubscriptionDB struct {
	database[AnnouncementSubscription]
}

var announcementsDB = AnnouncementSubscriptionDB{
	database[AnnouncementSubscription]{
		filename: "announcement-subscriptions.json",
	},
}

func init() {
	announcementsDB.load()
}

func GetAnnouncementSubscriptionDB() *AnnouncementSubscriptionDB {
	return &announcementsDB
}

func (db *AnnouncementSubscriptionDB) GetChatIDsOfSubscribedAdmins() []int64 {
	ids := make([]int64, len(db.data))
	for i, sub := range db.data {
		ids[i] = sub.ChatID
	}

	return ids
}

func (db *AnnouncementSubscriptionDB) Subscribe(chatID int64, userName string) {
	db.addNoDupe(AnnouncementSubscription{
		UserName: userName,
		ChatID:   chatID,
	}, func(a, b AnnouncementSubscription) bool { return a.ChatID == b.ChatID })
}

func (db *AnnouncementSubscriptionDB) Unsubscribe(chatID int64) {
	db.filterInPlace(func(ann AnnouncementSubscription) bool { return ann.ChatID != chatID })
}
