package db

import (
	"encoding/json"
	"log"
	"os"
)

type Channel struct {
	Id    int64
	Title string
}

const bannedChannelsDatabaseFile = "banned-channels.json"

func loadBannedChannelsDb() {
	content, err := os.ReadFile(bannedChannelsDatabaseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Failed to read %s: %s", bannedChannelsDatabaseFile, err.Error())
	}

	if err = json.Unmarshal(content, &bannedChannels); err != nil {
		log.Fatalf("Failed to unmarshal the contents of %s: %s", bannedChannelsDatabaseFile, err.Error())
	}
}

func writeDatabase() {
	content, err := json.Marshal(bannedChannels)
	if err != nil {
		log.Fatalf("Failed to marshal banned channels: %s", err.Error())
	}

	if err = os.WriteFile(bannedChannelsDatabaseFile, content, 0644); err != nil {
		log.Fatalf("Failed to write %s: %s", bannedChannelsDatabaseFile, err.Error())
	}
}

var bannedChannels []Channel

func GetBannedChannels() []Channel {
	return bannedChannels
}

func BanChannel(ch Channel) {
	bannedChannels = removeDuplicateChannels(append(bannedChannels, ch))
	writeDatabase()
}

func removeDuplicateChannels(sliceList []Channel) []Channel {
	allKeys := make(map[Channel]bool)
	list := []Channel{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
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
	bannedChannels = make([]Channel, 0)
	writeDatabase()
}
