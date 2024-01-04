package db

type announcementSubscription struct {
	UserName string
	ChatId   int64
}

type AnnouncementSubscriptionDB struct {
	database[announcementSubscription]
}

var announcementsDb = AnnouncementSubscriptionDB{database[announcementSubscription]{filename: "announcement-subscriptions.json"}}

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
	db.addNoDupe(announcementSubscription{
		UserName: userName,
		ChatId:   chatId,
	}, func(a, b announcementSubscription) bool { return a.ChatId == b.ChatId })
}

func (db *AnnouncementSubscriptionDB) Unsubscribe(chatId int64) {
	db.filterInPlace(func(ann announcementSubscription) bool { return ann.ChatId != chatId })
}
