package handler

import (
	"net/http"

	"job-application-app/backend/internal/model"
	"job-application-app/backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApplicationHandler struct {
	repo *repository.ApplicationRepository
}

func NewApplicationHandler(repo *repository.ApplicationRepository) *ApplicationHandler {
	return &ApplicationHandler{repo: repo}
}

func (h *ApplicationHandler) ApplyJob(c *gin.Context) {
	var req struct {
		JobID        uuid.UUID `json:"job_id"`
		FullName     string    `json:"full_name"`
		BirthDate    string    `json:"birth_date"`
		Address      string    `json:"address"`
		Phone        string    `json:"phone"`
		CVURL        string    `json:"cv_url"`
		PortfolioURL string    `json:"portfolio_url"`
		CoverLetter  string    `json:"cover_letter"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id") // dari JWT

	app := model.Application{
		ID:           uuid.New(),
		JobID:        req.JobID,
		UserID:       uuid.MustParse(userID),
		FullName:     req.FullName,
		Address:      req.Address,
		Phone:        req.Phone,
		CVURL:        req.CVURL,
		PortfolioURL: req.PortfolioURL,
		CoverLetter:  req.CoverLetter,
		Status:       "pending",
	}

	if err := h.repo.Create(&app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "application submitted",
		"data":    app,
	})
}

