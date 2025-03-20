package blog

type ArticleCategory struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	CreatedAt string `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp;not null"`
}

func (ArticleCategory) TableName() string {
	return "article_category"
}