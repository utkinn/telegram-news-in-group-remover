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

type MuteDB struct{ database[mute] }

var muteDb = MuteDB{database[mute]{filename: "mute.json"}}

func init() {
	muteDb.load()
}

func GetMuteDB() *MuteDB {
	return &muteDb
}

func (db *MuteDB) MuteUser(userName string) {
	db.add(mute{
		UserName: userName,
		StartAt:  time.Now(),
		EndAt:    time.Now().Add(randomMuteDuration()),
	})
}

func (db *MuteDB) UnmuteUser(userName string) {
	db.filterInPlace(func(item mute) bool { return item.UserName != userName })
}

func (db *MuteDB) IsUserMuted(userName string) (muted, announced bool) {
	db.filterInPlace(func(item mute) bool { return !item.EndAt.Before(time.Now()) })
	for _, item := range db.data {
		if item.UserName == userName {
			return true, item.IsAnnounced
		}
	}
	return false, false
}

func (db *MuteDB) MarkMuteAnnounced(userName string) {
	for i, item := range db.data {
		if item.UserName == userName {
			db.data[i].IsAnnounced = true
		}
	}
	db.write()
}

const minMuteDuration = time.Minute * 10
const maxMuteDuration = time.Minute * 60

func randomMuteDuration() time.Duration {
	return time.Duration(rand.Int63n(int64(maxMuteDuration) - int64(minMuteDuration) + int64(minMuteDuration)))
}
