package db

import (
	"log"
	"os"
	"path"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	testDBFilePath := path.Join(t.TempDir(), "db.json")
	if err := os.WriteFile(testDBFilePath, []byte("[0,1,2]"), 0o600); err != nil {
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

func TestWarningOnUnnecessaryLoad(t *testing.T) {
	oldWriter := log.Writer()
	logBuf := strings.Builder{}

	defer log.SetOutput(oldWriter)
	log.SetOutput(&logBuf)

	testDBFilePath := path.Join(t.TempDir(), "db.json")
	testDB := database[int]{filename: testDBFilePath, data: []int{0, 1, 2}}

	if err := os.WriteFile(testDBFilePath, []byte("[0,1,2]"), 0o600); err != nil {
		t.Fatalf("Failed to write to %v: %v", testDBFilePath, err)
	}

	testDB.load()

	if !strings.Contains(logBuf.String(), "Warning: unnecessary loading of already loaded DB") {
		t.Fatalf("Unexpected log output: %v", logBuf.String())
	}
}

func TestLoadPanicsOnReadFail(t *testing.T) {
	defer func() {
		rec := recover()
		if rec == nil {
			t.Fatal("Expected database.load() to panic")
		}

		if !strings.Contains(rec.(string), "Failed to read") {
			t.Fatalf("Unexpected panic message: %v", rec)
		}
	}()

	funkyPermissionsFilePath := path.Join(t.TempDir(), "funky-permissions.json")
	if err := os.WriteFile(funkyPermissionsFilePath, []byte("[0,1,2]"), 0o000); err != nil {
		t.Fatalf("Failed to create funky-permissions.json: %v", err)
	}

	testDB := database[int]{filename: funkyPermissionsFilePath}
	testDB.load()
}

func TestLoadGarbage(t *testing.T) {
	defer func() {
		rec := recover()
		if rec == nil {
			t.Fatal("Expected database.load() to panic")
		}

		if !strings.Contains(rec.(string), "Failed to unmarshal") {
			t.Fatalf("Unexpected panic message: %v", rec)
		}
	}()

	testDBFilePath := path.Join(t.TempDir(), "db.json")
	if err := os.WriteFile(testDBFilePath, []byte("garbage"), 0o600); err != nil {
		t.Fatalf("Failed to write to %v: %v", testDBFilePath, err)
	}

	testDB := database[int]{filename: testDBFilePath}
	testDB.load()
}

func TestWrite(t *testing.T) {
	testMutatingMethod(t, func(testDB *database[int]) {
		testDB.data = []int{0, 1, 2}
		testDB.write()
	}, "[0,1,2]")
}

type cyclicData struct {
	V int
	C *cyclicData
}

func TestWritePanic(t *testing.T) {
	var expectedPanicSubstring string

	handlePanic := func() {
		rec := recover() //nolint:revive
		if rec == nil {
			t.Fatal("Expected database.write() to panic")
		}

		if !strings.Contains(rec.(string), expectedPanicSubstring) {
			t.Fatalf("Unexpected panic message: %v", rec)
		}
	}

	t.Run("unspecified file name", func(t *testing.T) {
		expectedPanicSubstring = "Database file name is not specified"
		defer handlePanic()

		testDB := database[int]{data: []int{0, 1, 2}}
		testDB.write()
	})

	t.Run("cyclic data", func(t *testing.T) {
		expectedPanicSubstring = "Failed to marshal data"
		defer handlePanic()

		data := cyclicData{V: 1} //nolint:exhaustruct
		data.C = &data
		testDB := database[cyclicData]{filename: path.Join(t.TempDir(), "db.json"), data: []cyclicData{data}}
		testDB.write()
	})

	t.Run("unwritable file", func(t *testing.T) {
		expectedPanicSubstring = "Failed to write"
		defer handlePanic()

		unwritableFilePath := path.Join(t.TempDir(), "unwritable.json")
		if err := os.WriteFile(unwritableFilePath, []byte("[0,1,2]"), 0o000); err != nil {
			t.Fatalf("Failed to create unwritable.json: %v", err)
		}

		testDB := database[int]{filename: unwritableFilePath, data: []int{0, 1, 2}}
		testDB.write()
	})
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
	callback := func(item int) bool { return item == 1 }

	if no1.any(callback) {
		t.Fatal("database.any() returned true incorrectly")
	}

	if !has1.any(callback) {
		t.Fatal("database.any() returned false incorrectly")
	}
}

func TestFirst(t *testing.T) {
	database := database[int]{data: []int{1, 1, 4, 6, 1, 1}}

	result, found := database.first(func(item int) bool { return item%2 == 0 })
	if result != 4 || !found {
		t.Fatalf("want (4, true), got (%v, %v)", result, found)
	}

	result, found = database.first(func(item int) bool { return item == -1 })
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
