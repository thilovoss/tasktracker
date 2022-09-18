package models

import "time"

type Task struct {
	ID          string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Start       time.Time `gorm:"column:start" json:"start"`
	End         time.Time `gorm:"column:end" json:"end"`
	CategoryID  string    `gorm:"column:category_id" json:"categoryId"`
}

type Category struct {
	ID    string `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title string `gorm:"column:titke" json:"title"`
}
