package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	ArticleId uint      `json:"article_id"` //Refrence key of articles table.
	Name      string    `json:"name"`
	Content   string    `json:"content" gorm:"type:text"`
	Timestamp time.Time `json:"timestamp" gorm:"default:current_timestamp"`
	IsDeleted bool      `json:"is_deleted,omitempty" gorm:"default:false" `
}
