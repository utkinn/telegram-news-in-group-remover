package main

import "image"

func isScreenshot(img image.Image) bool {
	return aspectRatioSeemsLikeScreenshot(img) && lotsOfWhitePixels(img)
}

func aspectRatioSeemsLikeScreenshot(img image.Image) bool {
	bounds := img.Bounds()
	return float32(bounds.Dy())/float32(bounds.Dx()) > 2
}

const pixelColorMaxValue = 0xffff

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

	return float32(otherOccurrences)/float32(maxValueOccurrences) < 3.5
}
