package app

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"
)

type IntervalService interface {
	CreateInterval(interval domain.Interval) error
}

// IntervalService provides interval-related operations
type intervalService struct {
	repo repository.IntervalRepository
}

// NewIntervalService creates a new IntervalService
func NewIntervalService(r repository.IntervalRepository) *intervalService {
	return &intervalService{repo: r}
}

func (s *intervalService) CreateInterval(interval domain.Interval) error {
	return s.repo.Create(interval)
}
