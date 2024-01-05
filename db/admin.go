package db

type AdminDB struct{ database[string] }

var adminDB = AdminDB{database[string]{
	filename: "admins.json",
}}

func init() {
	adminDB.load()
}

func GetAdminDB() *AdminDB {
	return &adminDB
}

func NewAdminDBForTesting() *AdminDB {
	return &AdminDB{database[string]{data: []string{"SuperAdmin", "Admin"}}}
}

func (db *AdminDB) IsAdmin(nick string) bool {
	for _, n := range db.data {
		if n == nick {
			return true
		}
	}

	return false
}

// IsSuperAdmin reports whether nick belongs to the Super Admin.
// Super-admin is the first admin in the admin database.
func (db *AdminDB) IsSuperAdmin(nick string) bool {
	return db.data[0] == nick
}
