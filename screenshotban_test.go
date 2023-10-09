package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"testing"
)

func TestIsScreenshotWithScreenshotWithText(t *testing.T) {
	testScreenshotRecognition(t, "text", true)
}

func TestIsScreenshotWithGoogleImageSearch(t *testing.T) {
	testScreenshotRecognition(t, "images", true)
}

func TestIsNotScreenshotRandomPhoto(t *testing.T) {
	testScreenshotRecognition(t, "projector", false)
}

func testScreenshotRecognition(t *testing.T, screenshotName string, expectedResult bool) {
	screenshotFile, err := os.Open(fmt.Sprintf("testdata/screenshots/%v.jpeg", screenshotName))
	if err != nil {
		t.Fatal("Failed to open the screenshot file")
	}

	screenshot, err := jpeg.Decode(screenshotFile)
	if err != nil {
		t.Fatalf("Failed to decode screenshot jpeg: %v", err)
	}

	if isScreenshot(screenshot) != expectedResult {
		t.Fatal("isScreenshot() returned wrong result")
	}
}
