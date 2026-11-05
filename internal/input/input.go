package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Command int

const (
	ToggleClock Command = iota
	ResetMatch
	StartExpulsion
	StartFoul
	PickLeft
	PickRight
	Cancel
	ToggleFullscreen
)

var bindings = []struct {
	key ebiten.Key
	cmd Command
}{
	{ebiten.KeySpace, ToggleClock},
	{ebiten.KeyR, ResetMatch},
	{ebiten.KeyE, StartExpulsion},
	{ebiten.KeyF, StartFoul},
	{ebiten.KeyArrowLeft, PickLeft},
	{ebiten.KeyArrowRight, PickRight},
	{ebiten.KeyEscape, Cancel},
	{ebiten.KeyF11, ToggleFullscreen},
}

// Poll devolve os comandos acionados neste frame.
func Poll() []Command {
	var cmds []Command
	for _, b := range bindings {
		if inpututil.IsKeyJustPressed(b.key) {
			cmds = append(cmds, b.cmd)
		}
	}
	return cmds
}
