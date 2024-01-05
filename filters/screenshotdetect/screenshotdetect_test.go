package screenshotdetect_test

import (
	"fmt"
	"image/jpeg"
	"os"
	"testing"

	"github.com/utkinn/telegram-news-in-group-remover/filters/screenshotdetect"
)

func TestIsScreenshotWithScreenshotWithText(t *testing.T) {
	testScreenshotRecognition(t, "text", true)
}

func TestIsScreenshotWithGoogleImageSearch(t *testing.T) {
	testScreenshotRecognition(t, "images", true)
}

func TestIsNotScreenshotRandomPhotos(t *testing.T) {
	testScreenshotRecognition(t, "projector", false)
	testScreenshotRecognition(t, "bald-dude", true)
}

//nolint:revive
func testScreenshotRecognition(t *testing.T, screenshotName string, expectedResult bool) {
	screenshotFile, err := os.Open(fmt.Sprintf("testdata/screenshots/%v.jpeg", screenshotName))
	if err != nil {
		t.Fatal("Failed to open the screenshot file")
	}

	screenshot, err := jpeg.Decode(screenshotFile)
	if err != nil {
		t.Fatalf("Failed to decode screenshot jpeg: %v", err)
	}

	if screenshotdetect.IsScreenshot(screenshot) != expectedResult {
		t.Fatal("IsScreenshot() returned wrong result")
	}
}
