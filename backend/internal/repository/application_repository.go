package repository

import (
	"job-application-app/backend/internal/model"

	"gorm.io/gorm"
)

type ApplicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

// ✅ CREATE APPLICATION
func (r *ApplicationRepository) Create(app *model.Application) error {
	return r.db.Create(app).Error
}

// ✅ GET BY USER ID
func (r *ApplicationRepository) FindByUserID(userID string) ([]model.Application, error) {
	var apps []model.Application

	err := r.db.
		Preload("Job").
		Preload("Job.Company").
		Where("user_id = ?", userID).
		Find(&apps).Error

	return apps, err
}

func (r *ApplicationRepository) UpdateStatus(id string, status string) error {
	return r.db.
		Model(&model.Application{}).
		Where("id = ?", id).
		Update("status", status).Error
}



