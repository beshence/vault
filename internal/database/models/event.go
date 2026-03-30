package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID           string    `gorm:"type:char(36);primaryKey" json:"id"`
	ParentID     *string   `gorm:"column:parent_id;type:char(36);index:idx_events_repository_parent,unique" json:"parent_id,omitempty"`
	RepositoryID string    `gorm:"column:repository_id;type:char(36);not null;index;index:idx_events_repository_parent,unique" json:"repository_id"`
	SessionID    string    `gorm:"column:session_id;type:char(36);not null;index" json:"session_id"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	CreatedAt    time.Time `json:"created_at"`

	Parent     *Event     `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Repository Repository `gorm:"foreignKey:RepositoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Session    Session    `gorm:"foreignKey:SessionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (e *Event) BeforeCreate(_ *gorm.DB) error {
	if e.ID == "" {
		e.ID = uuid.NewString()
	}

	return nil
}
