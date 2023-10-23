package db

import "strings"

var scrutinyDb = database[string]{
	filename: "scrutiny.json",
}

func AddToScrutiny(nick string) {
	scrutinyDb.add(normalizeNick(nick))
}

func RemoveFromScrutiny(nick string) bool {
	nick = normalizeNick(nick)
	return scrutinyDb.filterInPlaceAndWrite(func(n string) bool { return n != nick })
}

func IsUnderScrutiny(nick string) bool {
	return scrutinyDb.any(func(n string) bool { return n == nick })
}

func normalizeNick(nick string) string {
	return strings.Replace(nick, "@", "", 1)
}
