package storage

import (
	"log/slog"
	"shared/models"
)

type HistoryStore struct {
	entries []models.CalculationResult
	maxSize int
}

func NewHistoryStore(maxSize int) *HistoryStore {
	return &HistoryStore{
		entries: make([]models.CalculationResult, 0, maxSize),
		maxSize: maxSize,
	}
}

func (s *HistoryStore) Add(entry models.CalculationResult) {
	s.entries = append(s.entries, entry)

	if len(s.entries) > s.maxSize {
		slog.Info("history size exceeded, removing oldest entry", "oldest_entry", s.entries[0])
		s.entries = s.entries[1:]
	}
}

func (s *HistoryStore) GetAll() []models.CalculationResult {

	slog.Info("getting records from history store", "count", len(s.entries))
	return s.entries
}

func (s *HistoryStore) Clear() {

	slog.Info("clearing records from history", "count", len(s.entries))
	s.entries = make([]models.CalculationResult, 0, s.maxSize)
}
