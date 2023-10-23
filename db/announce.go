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
