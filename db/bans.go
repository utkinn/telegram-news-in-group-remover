package db

type Channel struct {
	Id    int64
	Title string
}

type BannedChannelsDB struct{ database[Channel] }

var bannedChannelsDb = BannedChannelsDB{database[Channel]{
	filename: "banned-channels.json",
}}

func init() {
	bannedChannelsDb.load()
}

func GetBannedChannelsDB() *BannedChannelsDB {
	return &bannedChannelsDb
}

func (db *BannedChannelsDB) GetBannedChannels() []Channel {
	return db.data
}

func (db *BannedChannelsDB) BanChannel(ch Channel) {
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

func (db *BannedChannelsDB) IsChannelIdBanned(channelId int64) bool {
	for _, ch := range db.data {
		if ch.Id == channelId {
			return true
		}
	}
	return false
}

func (db *BannedChannelsDB) ClearBannedChannels() {
	db.data = make([]Channel, 0)
	db.write()
}
