package db

type Chat struct {
	Id    int64
	Title string
}

type ChatDB struct{ database[Chat] }

var chatsDb = ChatDB{database[Chat]{
	filename: "group-chats.json",
}}

func init() {
	chatsDb.load()
}

func GetChatsDB() *ChatDB {
	return &chatsDb
}

func (db *ChatDB) Get() []Chat {
	return db.data
}

func (db *ChatDB) GetIdByOrdinal(num int) (int64, bool) {
	if num < 1 || num > len(db.data) {
		return 0, false
	}
	return db.data[num-1].Id, true
}

func (db *ChatDB) Add(chat Chat) {
	db.data = removeDuplicateChatsById(append(db.data, chat))
	db.write()
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
