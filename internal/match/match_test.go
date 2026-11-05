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

func TestFoulsStopAtFivePerTeam(t *testing.T) {
	m := New()
	for i := 0; i < MaxFouls+2; i++ {
		m.AddFoul(Left)
	}
	m.AddFoul(Right)

	if got := m.Fouls(Left); got != MaxFouls {
		t.Fatalf("faltas do time esquerdo = %d, esperava %d", got, MaxFouls)
	}
	if got := m.Fouls(Right); got != 1 {
		t.Fatalf("faltas do time direito = %d, esperava 1", got)
	}
}

func TestAddingPenaltyAlsoAddsFoul(t *testing.T) {
	m := New()
	m.AddPenalty(Right)

	if got := m.Fouls(Right); got != 1 {
		t.Fatalf("faltas do time direito após expulsão = %d, esperava 1", got)
	}
	if got := len(m.Penalties(Right)); got != 1 {
		t.Fatalf("expulsões do time direito = %d, esperava 1", got)
	}
}

func TestResetClearsFouls(t *testing.T) {
	m := New()
	m.AddFoul(Left)
	m.Reset()

	if got := m.Fouls(Left); got != 0 {
		t.Fatalf("faltas após reiniciar = %d, esperava 0", got)
	}
}
