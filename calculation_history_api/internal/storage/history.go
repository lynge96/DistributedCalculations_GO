package storage

import (
	"log"
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
		log.Printf("History size exceeded, removing oldest entry: %+v", s.entries[0])
		s.entries = s.entries[1:]
	}
}

func (s *HistoryStore) GetAll() []models.CalculationResult {

	log.Printf("Getting %d records from history store", len(s.entries))
	return s.entries
}

func (s *HistoryStore) Clear() {

	log.Printf("Clearing %d records from history", len(s.entries))
	s.entries = make([]models.CalculationResult, 0, s.maxSize)
}
