package db

import (
	"math/rand"
	"time"
)

type mute struct {
	UserName       string
	StartAt, EndAt time.Time
	IsAnnounced    bool
}

var muteDb = database[mute]{filename: "mute.json"}

func init() {
	muteDb.load()
}

func MuteUser(userName string) {
	muteDb.add(mute{
		UserName: userName,
		StartAt:  time.Now(),
		EndAt:    time.Now().Add(randomMuteDuration()),
	})
}

func UnmuteUser(userName string) {
	muteDb.removeNotMatching(func(item mute) bool { return item.UserName != userName })
}

func IsUserMuted(userName string) (muted, announced bool) {
	muteDb.removeNotMatching(func(item mute) bool { return !item.EndAt.Before(time.Now()) })
	for _, item := range muteDb.data {
		if item.UserName == userName {
			return true, item.IsAnnounced
		}
	}
	return false, false
}

func MarkMuteAnnounced(userName string) {
	for i, item := range muteDb.data {
		if item.UserName == userName {
			muteDb.data[i].IsAnnounced = true
		}
	}
	muteDb.write()
}

const minMuteDuration = time.Minute * 10
const maxMuteDuration = time.Minute * 60

func randomMuteDuration() time.Duration {
	return time.Duration(rand.Int63n(int64(maxMuteDuration) - int64(minMuteDuration) + int64(minMuteDuration)))
}
