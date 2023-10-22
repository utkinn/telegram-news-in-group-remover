package db

var scrutinyDb = database[string]{
	filename: "scrutiny.json",
}

func AddToScrutiny(nick string) {
	scrutinyDb.add(nick)
}

func RemoveFromScrutiny(nick string) bool {
	return scrutinyDb.removeByCallback(func(n string) bool { return n != nick })
}

func IsUnderScrutiny(nick string) bool {
	return scrutinyDb.anyByCallback(func(n string) bool { return n == nick })
}
