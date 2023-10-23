package db

import "strings"

var scrutinyDb = database[string]{
	filename: "scrutiny.json",
}

func AddToScrutiny(nick string) {
	scrutinyDb.addNoDupe(normalizeNick(nick), func(a, b string) bool { return a == b })
}

func RemoveFromScrutiny(nick string) bool {
	nick = normalizeNick(nick)
	return scrutinyDb.removeNotMatching(func(n string) bool { return n != nick })
}

func IsUnderScrutiny(nick string) bool {
	return scrutinyDb.any(func(n string) bool { return n == nick })
}

func normalizeNick(nick string) string {
	return strings.Replace(nick, "@", "", 1)
}
