package db

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

type BannedRegexDB struct{ database[string] }

var bannedRegexesDB = BannedRegexDB{database[string]{
	filename: "regexes.json",
}}

func init() {
	bannedRegexesDB.load()
}

func GetBannedRegexDB() *BannedRegexDB {
	return &bannedRegexesDB
}

func (db *BannedRegexDB) Ban(regex string) error {
	_, err := regexp2.Compile(regex, regexp2.None)
	if err != nil {
		return fmt.Errorf("failed to ban regex %v: %w", regex, err)
	}

	db.add(regex)

	return nil
}

func (db *BannedRegexDB) Get() []string {
	return db.data
}

func (db *BannedRegexDB) Unban(regex string) {
	db.filterInPlace(func(r string) bool { return regex != r })
}
