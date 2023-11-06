package db

import (
	"encoding/json"
	"log"
	"os"
)

type database[T any] struct {
	filename string
	data     []T
}

func (db *database[T]) load() {
	if db.data != nil {
		log.Printf("Warning: unnecessary loading of already loaded DB %v. Check init() functions in db source dir.\n", db.filename)
	}

	var err error
	content, err := os.ReadFile(db.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Failed to read %s: %s", db.filename, err.Error())
	}

	if err = json.Unmarshal(content, &db.data); err != nil {
		log.Fatalf("Failed to unmarshal the contents of %s: %s", db.filename, err.Error())
	}
}

func (db *database[T]) write() {
	content, err := json.Marshal(db.data)
	if err != nil {
		log.Fatalf("Failed to marshal banned channels: %s", err.Error())
	}

	if err = os.WriteFile(db.filename, content, 0644); err != nil {
		log.Fatalf("Failed to write %s: %s", db.filename, err.Error())
	}
}

func (db *database[T]) add(item T) {
	db.data = append(db.data, item)
	db.write()
}

// addNoDupe acts like add, but does not insert the item if there is at least one item in the database that causes
// the equal callback applied to all database items to return true at least once.
func (db *database[T]) addNoDupe(item T, equal func(a, b T) bool) {
	for _, x := range db.data {
		if equal(x, item) {
			return
		}
	}
	db.add(item)
}

// addOrReplace acts like add, but if the equal callback, being applied to all items in the database, returns true on
// some item, this item gets replace by the supplied one.
//
// All items that may cause the equal callback to return true after the first one are ignored.
func (db *database[T]) addOrReplace(item T, equal func(a, b T) bool) {
	for i, x := range db.data {
		if equal(x, item) {
			db.data[i] = item
			db.write()
			return
		}
	}
	db.add(item)
}

func (db *database[T]) removeNotMatching(cb func(item T) bool) bool {
	removedSomething := false
	newData := make([]T, 0, len(db.data))
	for _, item := range db.data {
		if cb(item) {
			newData = append(newData, item)
		} else {
			removedSomething = true
		}
	}
	db.data = newData
	db.write()

	return removedSomething
}

func (db *database[T]) any(cb func(item T) bool) bool {
	for _, item := range db.data {
		if cb(item) {
			return true
		}
	}
	return false
}

func (db *database[T]) first(predicate func(item T) bool) (T, bool) {
	for _, item := range db.data {
		if predicate(item) {
			return item, true
		}
	}
	return *new(T), false
}
