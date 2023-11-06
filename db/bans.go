package db

type Channel struct {
	Id    int64
	Title string
}

var bannedChannelsDb = database[Channel]{
	filename: "banned-channels.json",
}

func init() {
	bannedChannelsDb.load()
}

func GetBannedChannels() []Channel {
	return bannedChannelsDb.data
}

func BanChannel(ch Channel) {
	bannedChannelsDb.data = removeDuplicateChannelsById(append(bannedChannelsDb.data, ch))
	bannedChannelsDb.write()
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

func IsChannelIdBanned(channelId int64) bool {
	for _, ch := range GetBannedChannels() {
		if ch.Id == channelId {
			return true
		}
	}
	return false
}

func ClearBannedChannels() {
	bannedChannelsDb.data = make([]Channel, 0)
	bannedChannelsDb.write()
}
