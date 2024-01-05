package db

type AnnouncementSubscription struct {
	UserName string
	ChatId   int64
}

type AnnouncementSubscriptionDB struct {
	database[AnnouncementSubscription]
}

var announcementsDb = AnnouncementSubscriptionDB{database[AnnouncementSubscription]{filename: "announcement-subscriptions.json"}}

func init() {
	announcementsDb.load()
}

func GetAnnouncementSubscriptionDB() *AnnouncementSubscriptionDB {
	return &announcementsDb
}

func (db *AnnouncementSubscriptionDB) GetChatIdsOfSubscribedAdmins() []int64 {
	ids := make([]int64, len(db.data))
	for i, sub := range db.data {
		ids[i] = sub.ChatId
	}
	return ids
}

func (db *AnnouncementSubscriptionDB) Subscribe(chatId int64, userName string) {
	db.addNoDupe(AnnouncementSubscription{
		UserName: userName,
		ChatId:   chatId,
	}, func(a, b AnnouncementSubscription) bool { return a.ChatId == b.ChatId })
}

func (db *AnnouncementSubscriptionDB) Unsubscribe(chatId int64) {
	db.filterInPlace(func(ann AnnouncementSubscription) bool { return ann.ChatId != chatId })
}
