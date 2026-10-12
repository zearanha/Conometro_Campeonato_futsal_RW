package match

import (
	"testing"
	"time"
)

func TestPenaltyExpiresAfterTwoMinutes(t *testing.T) {
	m := New()
	m.Toggle()
	m.AddPenalty(Left)

	m.Tick(119 * time.Second)
	if len(m.Penalties(Left)) != 1 {
		t.Fatal("a expulsão deveria continuar ativa")
	}
	m.Tick(2 * time.Second)
	if len(m.Penalties(Left)) != 0 {
		t.Fatal("a expulsão deveria ter acabado")
	}
}

func TestPausedClockDoesNotAdvance(t *testing.T) {
	m := New()
	m.Tick(time.Minute)
	if m.Remaining() != Duration {
		t.Fatal("o tempo não deveria correr pausado")
	}
}

func TestEndOfMatchClearsPenalties(t *testing.T) {
	m := New()
	m.Toggle()
	m.AddPenalty(Right)
	m.Tick(Duration + time.Second)

	if !m.Finished() || m.Running() || len(m.Penalties(Right)) != 0 {
		t.Fatal("fim de jogo deveria parar o relógio e limpar as expulsões")
	}
}