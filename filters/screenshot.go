package filters

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/filters/screenshotdetect"
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

	return !screenshotdetect.IsScreenshot(*img)
}

func (*screenshotFilter) ScrutinyModeOnly() bool {
	return true
}

func (*screenshotFilter) ShouldSuppressMock() bool {
	return false
}

func (*screenshotFilter) Description() Description {
	return Description{
		ID:   "screenshots",
		Name: "Скриншоты",
		Desc: "Удаляет скриншоты. Пытается отсеять скриншоты новостей и удалять только их." + unstableNotice,
	}
}

func (*screenshotFilter) downloadScreenshot(ctx helpers.ResponseContext) (*image.Image, error) {
	largestPhotoSize := ctx.Message.Photo[0]

	screenshotFile, err := ctx.Bot.GetFile(tgbotapi.FileConfig{FileID: largestPhotoSize.FileID})
	if err != nil {
		return nil, fmt.Errorf("failed to get screenshot Telegram file: %w", err)
	}

	screenshotLink := screenshotFile.Link(ctx.Bot.Token)

	resp, err := http.Get(screenshotLink) //nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("failed to download screenshot: %w", err)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode screenshot http response: %w", err)
	}

	return &img, nil
}
