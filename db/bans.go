package db

type Channel struct {
	ID    int64 `json:"Id"`
	Title string
}

type BannedChannelDB struct{ database[Channel] }

var bannedChannelsDB = BannedChannelDB{database[Channel]{
	filename: "banned-channels.json",
}}

func init() {
	bannedChannelsDB.load()
}

func GetBannedChannelDB() *BannedChannelDB {
	return &bannedChannelsDB
}

func (db *BannedChannelDB) Get() []Channel {
	return db.data
}

func (db *BannedChannelDB) Ban(ch Channel) {
	db.data = removeDuplicateChannelsByID(append(db.data, ch))
	db.write()
}

func removeDuplicateChannelsByID(sliceList []Channel) []Channel {
	allKeys := make(map[int64]bool)
	list := []Channel{}

	for _, item := range sliceList {
		if _, value := allKeys[item.ID]; !value {
			allKeys[item.ID] = true

			list = append(list, item)
		}
	}

	return list
}

func (db *BannedChannelDB) IsBanned(channelID int64) bool {
	for _, ch := range db.data {
		if ch.ID == channelID {
			return true
		}
	}

	return false
}

func (db *BannedChannelDB) Clear() {
	db.data = make([]Channel, 0)
	db.write()
}
