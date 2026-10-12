package ui

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"cronometro-futsal/internal/input"
	"cronometro-futsal/internal/match"
)

type Game struct {
	match        *match.Match
	last         time.Time
	awaitingTeam bool
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
		g.awaitingTeam = false
	}
	return nil
}

func (g *Game) handle(cmd input.Command) {
	switch cmd {
	case input.ToggleClock:
		g.match.Toggle()
	case input.ResetMatch:
		g.match.Reset()
		g.awaitingTeam = false
	case input.ToggleFullscreen:
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	case input.StartExpulsion:
		if !g.match.Finished() {
			g.awaitingTeam = !g.awaitingTeam
		}
	case input.PickLeft:
		g.pick(match.Left)
	case input.PickRight:
		g.pick(match.Right)
	case input.Cancel:
		g.awaitingTeam = false
	}
}

func (g *Game) pick(t match.Team) {
	if !g.awaitingTeam {
		return
	}
	g.match.AddPenalty(t)
	g.awaitingTeam = false
}

func (g *Game) Layout(outW, outH int) (int, int) { return screenW, screenH }