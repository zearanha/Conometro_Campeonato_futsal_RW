package ui

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"cronometro-futsal/internal/input"
	"cronometro-futsal/internal/match"
)

type teamSelection int

const (
	noSelection teamSelection = iota
	expulsionSelection
	foulSelection
)

type Game struct {
	match     *match.Match
	last      time.Time
	selection teamSelection
}

func NewGame() *Game {
	return &Game{match: match.New(), last: time.Now()}
}

func (g *Game) Update() error {
	now := time.Now()
	dt := now.Sub(g.last)
	g.last = now

	for _, cmd := range input.Poll() {
		g.handle(cmd)
	}
	g.match.Tick(dt)

	if g.match.Finished() {
		g.selection = noSelection
	}
	return nil
}

func (g *Game) handle(cmd input.Command) {
	switch cmd {
	case input.ToggleClock:
		g.match.Toggle()
	case input.ResetMatch:
		g.match.Reset()
		g.selection = noSelection
	case input.ToggleFullscreen:
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	case input.StartExpulsion:
		if !g.match.Finished() {
			if g.selection == expulsionSelection {
				g.selection = noSelection
			} else {
				g.selection = expulsionSelection
			}
		}
	case input.StartFoul:
		if !g.match.Finished() {
			if g.selection == foulSelection {
				g.selection = noSelection
			} else {
				g.selection = foulSelection
			}
		}
	case input.PickLeft:
		g.pick(match.Left)
	case input.PickRight:
		g.pick(match.Right)
	case input.Cancel:
		g.selection = noSelection
	}
}

func (g *Game) pick(t match.Team) {
	switch g.selection {
	case expulsionSelection:
		g.match.AddPenalty(t)
	case foulSelection:
		g.match.AddFoul(t)
	default:
		return
	}
	g.selection = noSelection
}

func (g *Game) Layout(outW, outH int) (int, int) { return screenW, screenH }
