package main

import (
	"encoding/json"
	"log"
	"os"
)

const bannedChannelsDatabaseFile = "banned-channels.json"

func readBannedChannelsDatabase() {
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

var bannedChannels []channel

func getBannedChannels() []channel {
	return bannedChannels
}

func banChannel(ch channel) {
	bannedChannels = removeDuplicateChannels(append(bannedChannels, ch))
	writeDatabase()
}

func removeDuplicateChannels(sliceList []channel) []channel {
	allKeys := make(map[channel]bool)
	list := []channel{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func isChannelIdBanned(channelId int64) bool {
	for _, ch := range getBannedChannels() {
		if ch.Id == channelId {
			return true
		}
	}
	return false
}

func clearBannedChannels() {
	bannedChannels = make([]channel, 0)
	writeDatabase()
}
