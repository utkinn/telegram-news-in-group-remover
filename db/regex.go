package db

import (
	"github.com/dlclark/regexp2"
)

type BannedRegexDB struct{ database[string] }

var bannedRegexesDb = BannedRegexDB{database[string]{
	filename: "regexes.json",
}}

func init() {
	bannedRegexesDb.load()
}

func GetBannedRegexDB() *BannedRegexDB {
	return &bannedRegexesDb
}

func (db *BannedRegexDB) Ban(regex string) error {
	_, err := regexp2.Compile(regex, regexp2.None)
	if err == nil {
		db.add(regex)
	}
	return err
}

func (db *BannedRegexDB) Get() []string {
	return db.data
}

func (db *BannedRegexDB) Unban(regex string) {
	db.filterInPlace(func(r string) bool { return regex != r })
}
