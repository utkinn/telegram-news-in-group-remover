package db

type Channel struct {
	Id    int64
	Title string
}

type BannedChannelDB struct{ database[Channel] }

var bannedChannelsDb = BannedChannelDB{database[Channel]{
	filename: "banned-channels.json",
}}

func init() {
	bannedChannelsDb.load()
}

func GetBannedChannelDB() *BannedChannelDB {
	return &bannedChannelsDb
}

func (db *BannedChannelDB) Get() []Channel {
	return db.data
}

func (db *BannedChannelDB) Ban(ch Channel) {
	db.data = removeDuplicateChannelsById(append(db.data, ch))
	db.write()
}

func removeDuplicateChannelsById(sliceList []Channel) []Channel {
	allKeys := make(map[int64]bool)
	list := []Channel{}
	for _, item := range sliceList {
		if _, value := allKeys[item.Id]; !value {
			allKeys[item.Id] = true
			list = append(list, item)
		}
	}
	return list
}

func (db *BannedChannelDB) IsBanned(channelId int64) bool {
	for _, ch := range db.data {
		if ch.Id == channelId {
			return true
		}
	}
	return false
}

func (db *BannedChannelDB) Clear() {
	db.data = make([]Channel, 0)
	db.write()
}
