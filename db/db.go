package db

import (
	"encoding/json"
	"log"
	"os"
)

func Load() {
	adminsDb.load()
	bannedChannelsDb.load()
	stickersDb.load()
	nameReplacementsDb.load()
}

type database[T any] struct {
	filename string
	data     []T
}

func (db *database[T]) load() {
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
