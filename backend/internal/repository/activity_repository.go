package repository

import (
	"database/sql"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

type ActivityRepository struct {
    db *sql.DB
}

func NewActivityRepository(db *sql.DB) *ActivityRepository {
    return &ActivityRepository{db: db}
}

func (r *ActivityRepository) Create(activity domain.Activity) error {
    _, err := r.db.Exec(
        "INSERT INTO activities (distance, start, finish, size, laps, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
         activity.Distance, activity.Start, activity.Finish, activity.Size, activity.Laps, activity.UserID,
	)
    return err
}

func (r *ActivityRepository) GetAll() ([]domain.Activity, error) {
    rows, err := r.db.Query("SELECT id, name, distance, start, finish, size, laps, user_id FROM activities")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var activities []domain.Activity
    for rows.Next() {
        var ex domain.Activity
        err = rows.Scan(&ex.ID,  &ex.Distance, &ex.Start, &ex.Finish, &ex.Size, &ex.Laps, &ex.UserID)
        if err != nil {
            return nil, err
        }
        activities = append(activities, ex)
    }
    return activities, nil
}