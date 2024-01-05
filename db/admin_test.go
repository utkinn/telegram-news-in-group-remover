package db

import "testing"

func TestIsAdmin(t *testing.T) {
	adminDB := AdminDB{database[string]{data: []string{"SuperAdminUserName", "RegularAdminUserName"}}}

	if !adminDB.IsAdmin("RegularAdminUserName") {
		t.Fatal("It thinks that RegularAdminUserName is not an admin")
	}

	if adminDB.IsAdmin("RandomDude") {
		t.Fatal("It thinks that RandomDude is an admin")
	}
}

func TestIsSuperAdmin(t *testing.T) {
	adminDB := AdminDB{database[string]{data: []string{"SuperAdminUserName", "RegularAdminUserName"}}}

	if !adminDB.IsSuperAdmin("SuperAdminUserName") {
		t.Fatal("It thinks that SuperAdminUserName is not a super admin")
	}

	if adminDB.IsSuperAdmin("RegularAdminUserName") {
		t.Fatal("It thinks that RegularAdminUserName is a super admin")
	}

	if adminDB.IsSuperAdmin("RandomDude") {
		t.Fatal("It thinks that RandomDude is a super admin")
	}
}
