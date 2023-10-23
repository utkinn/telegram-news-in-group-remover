package db

type announcementSubscription struct {
	UserName string
	ChatId   int64
}

var announcementsDb = database[announcementSubscription]{filename: "announcement-subscriptions.json"}

func GetChatIdsOfAdminsSubscribedToAnnouncements() []int64 {
	ids := make([]int64, len(announcementsDb.data))
	for i, sub := range announcementsDb.data {
		ids[i] = sub.ChatId
	}
	return ids
}

func SubscribeToAnnouncements(chatId int64, userName string) {
	announcementsDb.add(announcementSubscription{
		UserName: userName,
		ChatId:   chatId,
	})
}

func UnsubscribeFromAnnouncements(chatId int64) {
	announcementsDb.filterInPlaceAndWrite(func(ann announcementSubscription) bool { return ann.ChatId != chatId })
}
