package db

import (
	"time"
)

type mute struct {
	UserName       string
	StartAt, EndAt time.Time
	IsAnnounced    bool
}

type clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

type MuteDB struct {
	database[mute]
	clock clock
}

var muteDb = MuteDB{database[mute]{filename: "mute.json"}, realClock{}}

func init() {
	muteDb.load()
}

func GetMuteDB() *MuteDB {
	return &muteDb
}

func (db *MuteDB) MuteUser(userName string, duration time.Duration) {
	db.add(mute{
		UserName: userName,
		StartAt:  db.clock.Now(),
		EndAt:    db.clock.Now().Add(duration),
	})
}

func (db *MuteDB) UnmuteUser(userName string) {
	db.filterInPlace(func(item mute) bool { return item.UserName != userName })
}

func (db *MuteDB) GetStatusForUser(userName string) (muted, announced bool) {
	db.cleanUpExpiredMutes()
	for _, item := range db.data {
		if item.UserName == userName {
			return true, item.IsAnnounced
		}
	}
	return false, false
}

func (db *MuteDB) cleanUpExpiredMutes() {
	db.filterInPlace(func(item mute) bool { return !item.EndAt.Before(time.Now()) })
}

func (db *MuteDB) MarkMuteAnnounced(userName string) {
	for i, item := range db.data {
		if item.UserName == userName {
			db.data[i].IsAnnounced = true
		}
	}
	db.write()
}
