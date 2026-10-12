package ui

import "image/color"

const (
	screenW = 1280
	screenH = 720
)

var (
	colorBG     = color.RGBA{0, 0, 0, 255}
	colorOrange = color.RGBA{255, 140, 0, 255}
	colorWhite  = color.RGBA{255, 255, 255, 255}
	colorHint   = color.RGBA{110, 110, 110, 255}
)

const (
	sizeMain    = 280.0
	sizePenalty = 80.0
	sizeStatus  = 28.0
	sizeHelp    = 22.0

	penaltyTopY    = 440.0
	penaltyStepY   = 80.0
	penaltyMargin  = 50.0
	maxPenaltyRows = 3
)