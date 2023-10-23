package db

type announcementSubscription struct {
	userName string
	chatId   int64
}

var announcementsDb = database[announcementSubscription]{filename: "announcement-subscriptions.json"}

func GetChatIdsOfAdminsSubscribedToAnnouncements() []int64 {
	ids := make([]int64, len(announcementsDb.data))
	for i, sub := range announcementsDb.data {
		ids[i] = sub.chatId
	}
	return ids
}

func SubscribeToAnnouncements(chatId int64, userName string) {
	announcementsDb.add(announcementSubscription{
		userName: userName,
		chatId:   chatId,
	})
}

func UnsubscribeFromAnnouncements(chatId int64) {
	announcementsDb.filterInPlaceAndWrite(func(ann announcementSubscription) bool { return ann.chatId != chatId })
}
