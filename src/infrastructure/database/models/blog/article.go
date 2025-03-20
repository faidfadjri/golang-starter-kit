package blog

import (
	"akastra-mobile-api/src/infrastructure/database/models/users"
	"time"
)

type Article struct {
	ID          uint64     			`gorm:"primaryKey;autoIncrement;column:id"`
	Slug        string     			`gorm:"type:varchar(255);not null;column:slug"`
	Title       string     			`gorm:"type:varchar(255);not null;column:title"`
	Description string     			`gorm:"type:text;not null;column:description"`
	Image       string     			`gorm:"type:varchar(255);not null;column:image"`
	Like        int        			`gorm:"type:int(10);not null;default:0;column:like"`
	Views       int        			`gorm:"type:int(10);not null;default:0;column:views"`

	WriterID    uint64     			`gorm:"type:bigint(20) unsigned;column:writer_id"`
	Writer      users.User  		`gorm:"foreignKey:WriterID;references:ID"`

	CategoryID  uint64     			`gorm:"type:bigint(20) unsigned;column:category_id"`
	Category    ArticleCategory  	`gorm:"foreignKey:CategoryID;references:ID"`

	UpdatedAt   *time.Time 			`gorm:"column:updated_at"`
}

func (Article) TableName() string {
	return "article"
}