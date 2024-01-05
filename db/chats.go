package db

type Chat struct {
	ID    int64 `json:"Id"`
	Title string
}

type ChatDB struct{ database[Chat] }

var chatsDB = ChatDB{database[Chat]{
	filename: "group-chats.json",
}}

func init() {
	chatsDB.load()
}

func GetChatsDB() *ChatDB {
	return &chatsDB
}

func (db *ChatDB) Get() []Chat {
	return db.data
}

func (db *ChatDB) GetIDByOrdinal(num int) (int64, bool) {
	if num < 1 || num > len(db.data) {
		return 0, false
	}

	return db.data[num-1].ID, true
}

func (db *ChatDB) Add(chat Chat) {
	db.data = removeDuplicateChatsByID(append(db.data, chat))
	db.write()
}

func removeDuplicateChatsByID(sliceList []Chat) []Chat {
	allKeys := make(map[int64]bool)
	list := []Chat{}

	for _, item := range sliceList {
		if _, value := allKeys[item.ID]; !value {
			allKeys[item.ID] = true

			list = append(list, item)
		}
	}

	return list
}
