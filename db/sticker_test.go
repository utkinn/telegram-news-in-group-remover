package db

import "testing"

type fakeRandom struct{}

func (fakeRandom) Intn(int) int { return 1 }

func TestStickerDB_GetRandomMockStickerFileId(t *testing.T) {
	database := StickerDB{
		database[string]{
			data: []string{
				"sticker1",
				"sticker2",
				"sticker3",
			},
		},
		fakeRandom{},
	}

	fileID := database.GetRandomMockStickerFileID()

	if fileID != "sticker2" {
		t.Errorf("Expected file ID to be sticker2, got %v", fileID)
	}
}
