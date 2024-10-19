package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type FontSizeType int

const (
	FontSmall FontSizeType = iota
	FontMedium
	FontLarge
)

// getFontSize returns the actual font size based on the enum
func getFontSize(size FontSizeType) float64 {
	switch size {
	case FontSmall:
		return 12.0
	case FontMedium:
		return 16.0
	case FontLarge:
		return 24.0
	default:
		return 16.0
	}
}

func drawText(screen *ebiten.Image, char string, x, y float64, fontSize FontSizeType) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = getFontSize(fontSize) * 1.2
	text.Draw(screen, char, &text.GoTextFace{
		Source: faceSource,
		Size:   getFontSize(fontSize),
	}, op)
}
