package app

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"

	"github.com/google/uuid"
)

type ActivityService interface {
	CreateActivity(activity domain.Activity) error
	GetAllActivities() ([]domain.Activity, error)
	GetActivitiesByUser(userID uuid.UUID) ([]domain.Activity, error)
	GetActivityByID(activityID uuid.UUID) (domain.Activity, error)
	UpdateActivity(activity domain.Activity) error
	DeleteActivity(activityID uuid.UUID) error
}

type activityService struct {
	repo repository.ActivityRepository
}

// NewActivityService creates a new ActivityService
func NewActivityService(r repository.ActivityRepository) *activityService {
	return &activityService{repo: r}
}

func (s *activityService) CreateActivity(activity domain.Activity) error {
	return s.repo.CreateActivity(activity)
}

func (s *activityService) GetAllActivities() ([]domain.Activity, error) {
	return s.repo.GetAllActivities()
}

func (s *activityService) GetActivitiesByUser(userID uuid.UUID) ([]domain.Activity, error) {
	return s.repo.GetActivitiesByUser(userID)
}

func (s *activityService) GetActivityByID(activityID uuid.UUID) (domain.Activity, error) {
	return s.repo.GetActivityByID(activityID)
}

func (s *activityService) UpdateActivity(activity domain.Activity) error {
	return s.repo.UpdateActivity(activity)
}

func (s *activityService) DeleteActivity(activityID uuid.UUID) error {
	return s.repo.DeleteActivity(activityID)
}
