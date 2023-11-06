package filters

import (
	"fmt"
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

type screenshotFilter struct{ bot *tgbotapi.BotAPI }

func (s *screenshotFilter) IsMessageAllowed(message *tgbotapi.Message) bool {
	if message.Photo == nil {
		return true
	}

	img, err := s.downloadScreenshot(message)
	if err != nil {
		log.Printf("Screenshot filter failure: %v. Allowing this message to pass through.\n", err)
		return true
	}

	return !isScreenshot(*img)
}

func (s *screenshotFilter) downloadScreenshot(message *tgbotapi.Message) (*image.Image, error) {
	largestPhotoSize := message.Photo[0]
	screenshotFile, err := s.bot.GetFile(tgbotapi.FileConfig{FileID: largestPhotoSize.FileID})
	if err != nil {
		return nil, fmt.Errorf("failed to get screenshot Telegram file: %v", err)
	}

	screenshotLink := screenshotFile.Link(s.bot.Token)
	resp, err := http.Get(screenshotLink)
	if err != nil {
		return nil, fmt.Errorf("failed to download screenshot: %v", err)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode screenshot http response: %v", err)
	}

	return &img, nil
}

func (s *screenshotFilter) ScrutinyModeOnly() bool {
	return true
}

func (s *screenshotFilter) ShouldSuppressMock() bool {
	return false
}

func isScreenshot(img image.Image) bool {
	return aspectRatioSeemsLikeScreenshot(img) && lotsOfWhitePixels(img)
}

func aspectRatioSeemsLikeScreenshot(img image.Image) bool {
	bounds := img.Bounds()
	return float32(bounds.Dy())/float32(bounds.Dx()) > 2
}

const pixelColorMaxValue = 0xffff
const whiteTolerance = 5.5

func lotsOfWhitePixels(img image.Image) bool {
	var maxValueOccurrences, otherOccurrences int

	addOccurrence := func(value uint32) {
		if value == pixelColorMaxValue {
			maxValueOccurrences++
		} else {
			otherOccurrences++
		}
	}

	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			addOccurrence(r)
			addOccurrence(g)
			addOccurrence(b)
		}
	}

	return float32(otherOccurrences)/float32(maxValueOccurrences) < whiteTolerance
}
