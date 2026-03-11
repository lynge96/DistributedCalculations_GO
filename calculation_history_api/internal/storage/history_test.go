package storage

import (
	"fmt"
	"shared/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	store := NewHistoryStore(5)

	result := models.CalculationResult{Expression: "2+4", Result: 6}
	store.Add(result)

	assert.Contains(t, store.GetAll(), result, "expected to find result in store")
}

func TestAdd_RemovesOldestEvent_WhenLimitIsExceeded(t *testing.T) {
	store := NewHistoryStore(2)

	first := models.CalculationResult{ID: uuid.New(), Expression: "1 + 1", Result: 2}
	second := models.CalculationResult{ID: uuid.New(), Expression: "2 + 2", Result: 4}
	third := models.CalculationResult{ID: uuid.New(), Expression: "3 + 3", Result: 6}

	store.Add(first)
	store.Add(second)
	store.Add(third)

	entries := store.GetAll()

	assert.Len(t, entries, 2, "expected 2 entries")
	assert.NotContains(t, entries, first, "expected first entry to be removed")
}

func TestMaxSize(t *testing.T) {
	store := NewHistoryStore(3)

	for i := range 5 {
		store.Add(models.CalculationResult{Expression: fmt.Sprintf("%d+%d", i, i)})
	}

	entries := store.GetAll()
	assert.Len(t, entries, 3, "expected 3 entries")
}

func TestClear(t *testing.T) {
	store := NewHistoryStore(5)
	store.Add(models.CalculationResult{Expression: "1+1"})
	store.Clear()

	entries := store.GetAll()
	assert.Empty(t, entries, "expected no entries")
}
