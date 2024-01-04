package db

type AdminDB database[string]

var adminDB *AdminDB

func init() {
	adminDB = &AdminDB{
		filename: "admins.json",
	}
	(*database[string])(adminDB).load()
}

func GetAdminDB() *AdminDB {
	return adminDB
}

func NewAdminDBForTesting() *AdminDB {
	return &AdminDB{data: []string{"SuperAdmin", "Admin"}}
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
