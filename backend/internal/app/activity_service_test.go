package app

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockActivityRepository is a mock implementation of ActivityRepository
type MockActivityRepository struct {
	mock.Mock
}

func (m *MockActivityRepository) CreateActivity(activity domain.Activity) error {
	args := m.Called(activity)
	return args.Error(0)
}

func (m *MockActivityRepository) GetAllActivities() ([]domain.Activity, error) {
	args := m.Called()
	return args.Get(0).([]domain.Activity), args.Error(1)
}

func (m *MockActivityRepository) GetActivitiesByUser(userID uuid.UUID) ([]domain.Activity, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Activity), args.Error(1)
}

func (m *MockActivityRepository) GetActivityByID(activityID uuid.UUID) (domain.Activity, error) {
	args := m.Called(activityID)
	return args.Get(0).(domain.Activity), args.Error(1)
}

func (m *MockActivityRepository) UpdateActivity(activity domain.Activity) error {
	args := m.Called(activity)
	return args.Error(0)
}

func (m *MockActivityRepository) DeleteActivity(activityID uuid.UUID) error {
	args := m.Called(activityID)
	return args.Error(0)
}

type MockIntervalRepository struct {
	mock.Mock
}

func (m *MockIntervalRepository) CreateInterval(interval domain.Interval) error {
	args := m.Called(interval)
	return args.Error(0)
}

func (m *MockIntervalRepository) GetIntervalsByActivity(activityID uuid.UUID) ([]domain.Interval, error) {
	args := m.Called(activityID)
	return args.Get(0).([]domain.Interval), args.Error(1)
}

func TestCreateActivity(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activity := domain.Activity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Start:        time.Now(),
		Duration:     domain.DurationString("1h30m"),
		Distance:     4000,
		Laps:         160,
		PoolSize:     25,
		LocationType: domain.LocationPool,
		Notes:        "Test activity",
	}

	mockRepo.On("CreateActivity", activity).Return(nil)

	err := service.CreateActivity(activity)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateActivity_Error(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activity := domain.Activity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Start:        time.Now(),
		Duration:     domain.DurationString("1h30m"),
		Distance:     4000,
		Laps:         160,
		PoolSize:     25,
		LocationType: domain.LocationPool,
		Notes:        "Test activity",
	}

	mockRepo.On("CreateActivity", activity).Return(errors.New("db error"))

	err := service.CreateActivity(activity)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllActivities(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activities := []domain.Activity{
		{
			ID:           uuid.New(),
			UserID:       uuid.New(),
			Start:        time.Now(),
			Duration:     domain.DurationString("1h30m"),
			Distance:     4000,
			Laps:         160,
			PoolSize:     25,
			LocationType: domain.LocationPool,
			Notes:        "Test activity",
		},
		{
			ID:           uuid.New(),
			UserID:       uuid.New(),
			Start:        time.Now().Add(-time.Hour),
			Duration:     domain.DurationString("45m"),
			Distance:     2000,
			Laps:         80,
			PoolSize:     25,
			LocationType: domain.LocationPool,
			Notes:        "Another test activity",
		},
	}

	mockRepo.On("GetAllActivities").Return(activities, nil)

	result, err := service.GetAllActivities()
	assert.NoError(t, err)
	assert.Equal(t, activities, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllActivities_Error(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)

	mockRepo.On("GetAllActivities").Return([]domain.Activity{}, errors.New("db error"))

	result, err := service.GetAllActivities()
	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetActivitiesByUser(t *testing.T) {
	mockActivityRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)

	service := &activityService{
		repo:         mockActivityRepo,
		intervalRepo: mockIntervalRepo,
	}

	userID := uuid.New()
	activityID := uuid.New()
	activities := []domain.Activity{
		{
			ID:           activityID,
			UserID:       userID,
			Start:        time.Now(),
			Duration:     domain.DurationString("1h"),
			Distance:     1000,
			Laps:         40,
			PoolSize:     25,
			LocationType: domain.LocationPool,
			Notes:        "Test",
		},
	}
	intervals := []domain.Interval{
		{
			ID:         uuid.New(),
			ActivityID: activityID,
			Duration:   domain.DurationString("30m"),
			Distance:   500,
			Type:       domain.IntervalType("swim"),
			Stroke:     domain.StrokeType("freestyle"),
			Notes:      "Test interval",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockActivityRepo.On("GetActivitiesByUser", userID).Return(activities, nil)
		mockIntervalRepo.On("GetIntervalsByActivity", activityID).Return(intervals, nil)

		result, err := service.GetActivitiesByUser(userID)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
		mockActivityRepo.AssertExpectations(t)
		mockIntervalRepo.AssertExpectations(t)
	})

	t.Run("activity repo error", func(t *testing.T) {
		mockActivityRepo.ExpectedCalls = nil
		mockIntervalRepo.ExpectedCalls = nil
		mockActivityRepo.On("GetActivitiesByUser", userID).Return([]domain.Activity{}, errors.New("repo error"))

		result, err := service.GetActivitiesByUser(userID)
		assert.Error(t, err)
		assert.Empty(t, result)
		mockActivityRepo.AssertExpectations(t)
	})

	t.Run("interval repo error", func(t *testing.T) {
		mockActivityRepo.ExpectedCalls = nil
		mockIntervalRepo.ExpectedCalls = nil

		mockActivityRepo.On("GetActivitiesByUser", userID).Return(activities, nil)
		mockIntervalRepo.On("GetIntervalsByActivity", activityID).Return([]domain.Interval{}, errors.New("interval error"))

		result, err := service.GetActivitiesByUser(userID)
		assert.Error(t, err)
		assert.Empty(t, result)
		mockActivityRepo.AssertExpectations(t)
		mockIntervalRepo.AssertExpectations(t)
	})
}

func TestGetActivityByID(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activityID := uuid.New()
	activity := domain.Activity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Start:        time.Now(),
		Duration:     domain.DurationString("1h30m"),
		Distance:     4000,
		Laps:         160,
		PoolSize:     25,
		LocationType: domain.LocationPool,
		Notes:        "Test activity",
	}

	mockRepo.On("GetActivityByID", activityID).Return(activity, nil)

	result, err := service.GetActivityByID(activityID)
	assert.NoError(t, err)
	assert.Equal(t, activity, result)
	mockRepo.AssertExpectations(t)
}

func TestGetActivityByID_Error(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activityID := uuid.New()

	mockRepo.On("GetActivityByID", activityID).Return(domain.Activity{}, errors.New("not found"))

	result, err := service.GetActivityByID(activityID)
	assert.Error(t, err)
	assert.Equal(t, domain.Activity{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateActivity(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activity := domain.Activity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Start:        time.Now(),
		Duration:     domain.DurationString("1h30m"),
		Distance:     4000,
		Laps:         160,
		PoolSize:     25,
		LocationType: domain.LocationPool,
		Notes:        "Test activity",
	}

	mockRepo.On("UpdateActivity", activity).Return(nil)

	err := service.UpdateActivity(activity)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateActivity_Error(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activity := domain.Activity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Start:        time.Now(),
		Duration:     domain.DurationString("1h30m"),
		Distance:     4000,
		Laps:         160,
		PoolSize:     25,
		LocationType: domain.LocationPool,
		Notes:        "Test activity",
	}

	mockRepo.On("UpdateActivity", activity).Return(errors.New("update error"))

	err := service.UpdateActivity(activity)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteActivity(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activityID := uuid.New()

	mockRepo.On("DeleteActivity", activityID).Return(nil)

	err := service.DeleteActivity(activityID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteActivity_Error(t *testing.T) {
	mockRepo := new(MockActivityRepository)
	mockIntervalRepo := new(MockIntervalRepository)
	service := NewActivityService(mockRepo, mockIntervalRepo)
	activityID := uuid.New()

	mockRepo.On("DeleteActivity", activityID).Return(errors.New("delete error"))

	err := service.DeleteActivity(activityID)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
