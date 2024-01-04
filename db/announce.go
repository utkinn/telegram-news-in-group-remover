package db

type announcementSubscription struct {
	UserName string
	ChatId   int64
}

type AnnouncementSubscriptionsDB struct {
	database[announcementSubscription]
}

var announcementsDb = AnnouncementSubscriptionsDB{database[announcementSubscription]{filename: "announcement-subscriptions.json"}}

func init() {
	announcementsDb.load()
}

func GetAnnouncementSubscriptionsDB() *AnnouncementSubscriptionsDB {
	return &announcementsDb
}

func (db *AnnouncementSubscriptionsDB) GetChatIdsOfAdminsSubscribedToAnnouncements() []int64 {
	ids := make([]int64, len(db.data))
	for i, sub := range db.data {
		ids[i] = sub.ChatId
	}
	return ids
}

func (db *AnnouncementSubscriptionsDB) SubscribeToAnnouncements(chatId int64, userName string) {
	db.addNoDupe(announcementSubscription{
		UserName: userName,
		ChatId:   chatId,
	}, func(a, b announcementSubscription) bool { return a.ChatId == b.ChatId })
}

func (db *AnnouncementSubscriptionsDB) UnsubscribeFromAnnouncements(chatId int64) {
	db.filterInPlace(func(ann announcementSubscription) bool { return ann.ChatId != chatId })
}
