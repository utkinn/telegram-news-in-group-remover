package db

import "testing"

func TestIsAdmin(t *testing.T) {
	adminsDb.data = []string{"SuperAdminUserName", "RegularAdminUserName"}

	if !IsAdmin("RegularAdminUserName") {
		t.Fatal("It thinks that RegularAdminUserName is not an admin")
	}
	if IsAdmin("RandomDude") {
		t.Fatal("It thinks that RandomDude is an admin")
	}
}

func TestIsSuperAdmin(t *testing.T) {
	adminsDb.data = []string{"SuperAdminUserName", "RegularAdminUserName"}

	if !IsSuperAdmin("SuperAdminUserName") {
		t.Fatal("It thinks that SuperAdminUserName is not a super admin")
	}
	if IsSuperAdmin("RegularAdminUserName") {
		t.Fatal("It thinks that RegularAdminUserName is a super admin")
	}
	if IsSuperAdmin("RandomDude") {
		t.Fatal("It thinks that RandomDude is a super admin")
	}
}
