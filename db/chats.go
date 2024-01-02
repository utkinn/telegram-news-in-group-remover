package db

type Chat struct {
	Id    int64
	Title string
}

var chatsDb = database[Chat]{
	filename: "group-chats.json",
}

func init() {
	chatsDb.load()
}

func GetChats() []Chat {
	return chatsDb.data
}

func GetChatIdByNumber(num int) (int64, bool) {
	if num < 1 || num > len(chatsDb.data) {
		return 0, false
	}
	return chatsDb.data[num-1].Id, true
}

func AddChat(chat Chat) {
	chatsDb.data = removeDuplicateChatsById(append(chatsDb.data, chat))
	chatsDb.write()
}

func removeDuplicateChatsById(sliceList []Chat) []Chat {
	allKeys := make(map[int64]bool)
	list := []Chat{}
	for _, item := range sliceList {
		if _, value := allKeys[item.Id]; !value {
			allKeys[item.Id] = true
			list = append(list, item)
		}
	}
	return list
}
