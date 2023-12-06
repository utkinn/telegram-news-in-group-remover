package db

var lastMessageChatId int64

func LastMessageChatId() int64 {
	return lastMessageChatId
}

func SetLastMessageChatId(id int64) {
	lastMessageChatId = id
}
