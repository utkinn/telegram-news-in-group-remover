package db

import (
	"os"
	"path"
	"testing"
)

func TestLoad(t *testing.T) {
	testDBFilePath := path.Join(t.TempDir(), "db.json")
	if err := os.WriteFile(testDBFilePath, []byte("[0,1,2]"), 0644); err != nil {
		t.Fatalf("Failed to write to %v: %v", testDBFilePath, err)
	}
	testDB := database[int]{filename: testDBFilePath}

	testDB.load()

	if len(testDB.data) != 3 {
		t.Fatalf("Unexpected len(testDB.data): want 3, got %v", len(testDB.data))
	}
	for i, n := range testDB.data {
		if n != i {
			t.Fatalf("Unexpected member at index %v: want %v, got %v", i, i, n)
		}
	}
}

func TestLoadEmpty(t *testing.T) {
	testDBFilePath := path.Join(t.TempDir(), "db.json")
	testDB := database[int]{filename: testDBFilePath}

	testDB.load()

	if len(testDB.data) != 0 {
		t.Fatalf("Unexpected len(testDB.data): want 0, got %v", len(testDB.data))
	}
}

func TestWrite(t *testing.T) {
	testMutatingMethod(t, func(testDB *database[int]) {
		testDB.data = []int{0, 1, 2}
		testDB.write()
	}, "[0,1,2]")
}

func TestAdd(t *testing.T) {
	testMutatingMethod(t, func(testDB *database[int]) {
		testDB.add(1)
	}, "[1]")
}

func TestAddNoDupe(t *testing.T) {
	testMutatingMethod(t, func(testDB *database[int]) {
		testDB.data = []int{1, 2}
		testDB.addNoDupe(1, func(a, b int) bool { return a == b })
		testDB.addNoDupe(3, func(a, b int) bool { return a == b })
	}, "[1,2,3]")
}

func TestAddOrReplace(t *testing.T) {
	testMutatingMethod(t, func(testDB *database[int]) {
		testDB.data = []int{1, 2, 4, 3, 8}
		testDB.addOrReplace(64, func(dbItem, newItem int) bool { return dbItem%2 == 0 })
	}, "[1,64,4,3,8]")
}

func TestFilterInPlace(t *testing.T) {
	testMutatingMethod(t, func(testDB *database[int]) {
		testDB.data = []int{1, 2, 4, 3, 8}
		testDB.filterInPlace(func(x int) bool { return x%2 == 0 })
	}, "[2,4,8]")
}

func TestAny(t *testing.T) {
	no1 := database[int]{data: []int{2, 3, 4}}
	has1 := database[int]{data: []int{1, 0}}
	cb := func(item int) bool { return item == 1 }

	if no1.any(cb) {
		t.Fatal("database.any() returned true incorrectly")
	}
	if !has1.any(cb) {
		t.Fatal("database.any() returned false incorrectly")
	}
}

func TestFirst(t *testing.T) {
	db := database[int]{data: []int{1, 1, 4, 6, 1, 1}}

	result, found := db.first(func(item int) bool { return item%2 == 0 })
	if result != 4 || !found {
		t.Fatalf("want (4, true), got (%v, %v)", result, found)
	}

	result, found = db.first(func(item int) bool { return item == -1 })
	if found {
		t.Fatalf("want (_, false), got (%v, %v)", result, found)
	}
}

func testMutatingMethod(t *testing.T, act func(testDB *database[int]), wantedContent string) {
	testDBFilePath := path.Join(t.TempDir(), "db.json")
	testDB := database[int]{filename: testDBFilePath}

	act(&testDB)

	contents, err := os.ReadFile(testDBFilePath)
	if err != nil {
		t.Fatalf("Failed to read %v: %v", testDBFilePath, err)
	}
	if string(contents) != wantedContent {
		t.Fatalf("Unexpected json file content: want %v, got %v", wantedContent, contents)
	}
}
