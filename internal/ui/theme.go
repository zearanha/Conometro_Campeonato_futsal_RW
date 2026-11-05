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
	sizeMain    = 340.0
	sizePenalty = 64.0
	sizeStatus  = 28.0
	sizeHelp    = 22.0
	sizeFouls   = 100.0
	mainClockY  = 100.0
	foulsY      = 20.0
	foulsMargin = 20.0

	penaltyTopY    = 470.0
	penaltyStepY   = 65.0
	penaltyMargin  = 50.0
	maxPenaltyRows = 3
)
