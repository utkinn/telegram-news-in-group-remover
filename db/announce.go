package db

type announcementSubscription struct {
	UserName string
	ChatId   int64
}

var announcementsDb = database[announcementSubscription]{filename: "announcement-subscriptions.json"}

func init() {
	announcementsDb.load()
}

func GetChatIdsOfAdminsSubscribedToAnnouncements() []int64 {
	ids := make([]int64, len(announcementsDb.data))
	for i, sub := range announcementsDb.data {
		ids[i] = sub.ChatId
	}
	return ids
}

func SubscribeToAnnouncements(chatId int64, userName string) {
	announcementsDb.addNoDupe(announcementSubscription{
		UserName: userName,
		ChatId:   chatId,
	}, func(a, b announcementSubscription) bool { return a.ChatId == b.ChatId })
}

func UnsubscribeFromAnnouncements(chatId int64) {
	announcementsDb.removeNotMatching(func(ann announcementSubscription) bool { return ann.ChatId != chatId })
}
