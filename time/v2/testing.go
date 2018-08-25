package poker

import (
	"fmt"
	"testing"
	"time"
)

// StubPlayerStore implements PlayerStore for testing purposes
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

// GetPlayerScore returns a score from Scores
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin will record a win to WinCalls
func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

// GetLeague returns League
func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

// AssertPlayerWin allows you to spy on the store's calls to RecordWin
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got '%s' want '%s'", store.WinCalls[0], winner)
	}
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}
