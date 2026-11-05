package match

import "time"

const (
	Duration        = 10 * time.Minute
	PenaltyDuration = 2 * time.Minute
	MaxFouls        = 5
)

type Team int

const (
	Left Team = iota
	Right
)

type Match struct {
	remaining time.Duration
	running   bool
	penalties [2][]time.Duration
	fouls     [2]int
}

func New() *Match {
	return &Match{remaining: Duration}
}

func (m *Match) Remaining() time.Duration { return m.remaining }
func (m *Match) Running() bool            { return m.running }
func (m *Match) Finished() bool           { return m.remaining == 0 }

// Penalties devolve uma cópia das contagens de expulsão do time.
func (m *Match) Penalties(t Team) []time.Duration {
	return append([]time.Duration(nil), m.penalties[t]...)
}

func (m *Match) Fouls(t Team) int { return m.fouls[t] }

// Toggle inicia ou pausa o jogo. Não faz nada se o jogo já acabou.
func (m *Match) Toggle() {
	if m.Finished() {
		return
	}
	m.running = !m.running
}

// Reset volta tudo ao estado inicial.
func (m *Match) Reset() {
	*m = Match{remaining: Duration}
}

// AddPenalty registra uma expulsão para o time.
func (m *Match) AddPenalty(t Team) {
	if m.Finished() {
		return
	}
	m.penalties[t] = append(m.penalties[t], PenaltyDuration)
	m.AddFoul(t)
}

// AddFoul registra uma falta para o time até o limite permitido.
func (m *Match) AddFoul(t Team) {
	if m.Finished() || m.fouls[t] >= MaxFouls {
		return
	}
	m.fouls[t]++
}

// Tick avança o relógio. Só tem efeito com o jogo rodando.
func (m *Match) Tick(dt time.Duration) {
	if !m.running {
		return
	}
	if dt > m.remaining {
		dt = m.remaining
	}
	m.remaining -= dt

	for t := range m.penalties {
		kept := m.penalties[t][:0]
		for _, p := range m.penalties[t] {
			if p -= dt; p > 0 {
				kept = append(kept, p)
			}
		}
		m.penalties[t] = kept
	}

	if m.remaining == 0 {
		m.running = false
		m.penalties = [2][]time.Duration{}
	}
}
