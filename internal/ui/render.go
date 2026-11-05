package ui

import (
	"bytes"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/gobold"

	"cronometro-futsal/internal/match"
)

var fontSource *text.GoTextFaceSource

func init() {
	src, err := text.NewGoTextFaceSource(bytes.NewReader(gobold.TTF))
	if err != nil {
		panic(err)
	}
	fontSource = src
}

type align int

const (
	alignLeft align = iota
	alignCenter
	alignRight
)

func formatDuration(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	// arredonda para cima: só mostra 00:00 quando realmente acabou
	secs := int((d + time.Second - 1) / time.Second)
	return fmt.Sprintf("%02d:%02d", secs/60, secs%60)
}

func drawText(screen *ebiten.Image, s string, size, x, y float64, a align, clr color.Color) {
	face := &text.GoTextFace{Source: fontSource, Size: size}
	w, _ := text.Measure(s, face, 0)
	switch a {
	case alignCenter:
		x -= w / 2
	case alignRight:
		x -= w
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(clr)
	text.Draw(screen, s, face, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colorBG)

	drawText(screen, fmt.Sprintf("%d", g.match.Fouls(match.Left)),
		sizeFouls, foulsMargin, foulsY, alignLeft, colorWhite)
	drawText(screen, fmt.Sprintf("%d", g.match.Fouls(match.Right)),
		sizeFouls, screenW-foulsMargin, foulsY, alignRight, colorWhite)
	drawText(screen, formatDuration(g.match.Remaining()), sizeMain, screenW/2, mainClockY, alignCenter, colorOrange)

	g.drawPenalties(screen, match.Left, penaltyMargin, alignLeft)
	g.drawPenalties(screen, match.Right, screenW-penaltyMargin, alignRight)

	if msg := g.statusMessage(); msg != "" {
		drawText(screen, msg, sizeStatus, screenW/2, 20, alignCenter, colorHint)
	}
	drawText(screen, "Espaco: iniciar/pausar   R: reiniciar   E: expulsao   F: falta   F11: tela cheia",
		sizeHelp, screenW/2, screenH-35, alignCenter, colorHint)
}

func (g *Game) drawPenalties(screen *ebiten.Image, t match.Team, x float64, a align) {
	for i, p := range g.match.Penalties(t) {
		if i >= maxPenaltyRows {
			break
		}
		y := penaltyTopY + float64(i)*penaltyStepY
		drawText(screen, formatDuration(p), sizePenalty, x, y, a, colorWhite)
	}
}

func (g *Game) statusMessage() string {
	switch {
	case g.selection == expulsionSelection:
		return "EXPULSAO: aperte < ou > para escolher o time (Esc cancela)"
	case g.selection == foulSelection:
		return "FALTA: aperte < ou > para escolher o time (Esc cancela)"
	case g.match.Finished():
		return "FIM DE JOGO  -  R para reiniciar"
	case !g.match.Running():
		return "PAUSADO  -  Espaco para iniciar"
	}
	return ""
}
