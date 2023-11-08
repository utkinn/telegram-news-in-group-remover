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

// IsSuperAdmin reports whether nick belongs to the Super Admin.
// Super-admin is the first admin in the admin database.
func IsSuperAdmin(nick string) bool {
	return adminsDb.data[0] == nick
}

func SetAdminsForTesting() {
	adminsDb.data = []string{"SuperAdmin", "Admin"}
}
