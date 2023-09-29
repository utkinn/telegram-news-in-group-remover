package db

var adminsDb = database[string]{
	filename: "admins.json",
}

func IsAdmin(nick string) bool {
	for _, n := range adminsDb.data {
		if n == nick {
			return true
		}
	}
	return false
}
