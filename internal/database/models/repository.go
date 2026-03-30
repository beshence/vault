package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var repositoryNamePattern = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

var (
	ErrInvalidRepositoryName   = errors.New("repository name must contain only English letters, numbers, '_' or '-' characters")
	ErrRepositoryOwnerRequired = errors.New("repository owner is required")
)

type Repository struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `gorm:"size:128;not null;uniqueIndex:idx_owner_name" json:"name"`
	OwnerID   string    `gorm:"column:owner_id;type:char(36);not null;index;uniqueIndex:idx_owner_name" json:"owner_id"`
	LastEvent *string   `gorm:"column:last_event_id;type:char(36);index" json:"last_event,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	Owner User `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (r *Repository) BeforeCreate(_ *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}

	return r.Validate()
}

func (r *Repository) Validate() error {
	if !repositoryNamePattern.MatchString(r.Name) {
		return ErrInvalidRepositoryName
	}

	if r.OwnerID == "" {
		return ErrRepositoryOwnerRequired
	}

	return nil
}
