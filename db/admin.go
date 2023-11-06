package db

var adminsDb = database[string]{
	filename: "admins.json",
}

func init() {
	adminsDb.load()
}

func IsAdmin(nick string) bool {
	for _, n := range adminsDb.data {
		if n == nick {
			return true
		}
	}
	return false
}

// Super-admin is the first admin in the list
func IsSuperAdmin(nick string) bool {
	return adminsDb.data[0] == nick
}
