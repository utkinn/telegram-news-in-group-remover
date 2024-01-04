package filters

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerFilter(&screenshotFilter{})
}

type screenshotFilter struct{}

func (s *screenshotFilter) IsMessageAllowed(ctx helpers.ResponseContext) bool {
	if ctx.Message.Photo == nil {
		return true
	}

	img, err := s.downloadScreenshot(ctx)
	if err != nil {
		log.Printf("Screenshot filter failure: %v. Allowing this message to pass through.\n", err)
		return true
	}

	return !isScreenshot(*img)
}

func (s *screenshotFilter) ScrutinyModeOnly() bool {
	return true
}

func (s *screenshotFilter) ShouldSuppressMock() bool {
	return false
}

func (s *screenshotFilter) Description() Description {
	return Description{
		ID:   "screenshots",
		Name: "Скриншоты",
		Desc: "Удаляет скриншоты. Пытается отсеять скриншоты новостей и удалять только их." + unstableNotice,
	}
}

func (s *screenshotFilter) downloadScreenshot(ctx helpers.ResponseContext) (*image.Image, error) {
	largestPhotoSize := ctx.Message.Photo[0]
	screenshotFile, err := ctx.Bot.GetFile(tgbotapi.FileConfig{FileID: largestPhotoSize.FileID})
	if err != nil {
		return nil, fmt.Errorf("failed to get screenshot Telegram file: %v", err)
	}

	screenshotLink := screenshotFile.Link(ctx.Bot.Token)
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
