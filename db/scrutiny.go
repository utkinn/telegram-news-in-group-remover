package db

import "strings"

type ScrutinyDB struct{ database[string] }

var scrutinyDB = ScrutinyDB{database[string]{
	filename: "scrutiny.json",
}}

func GetScrutinyDB() *ScrutinyDB {
	return &scrutinyDB
}

func init() {
	scrutinyDB.load()
}

func (db *ScrutinyDB) Add(nick string) {
	db.addNoDupe(normalizeNick(nick), func(a, b string) bool { return a == b })
}

func (db *ScrutinyDB) Remove(nick string) bool {
	nick = normalizeNick(nick)
	return db.filterInPlace(func(n string) bool { return n != nick })
}

func (db *ScrutinyDB) IsUnderScrutiny(nick string) bool {
	return db.any(func(n string) bool { return n == nick })
}

func normalizeNick(nick string) string {
	return strings.Replace(nick, "@", "", 1)
}
