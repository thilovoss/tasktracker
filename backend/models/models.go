package models

import "time"

type Task struct {
	ID          string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"column:name" json:"name" gorm-patch:"patchable"`
	Description string    `gorm:"column:description" json:"description" gorm-patch:"patchable"`
	Start       time.Time `gorm:"column:start" json:"start" gorm-patch:"patchable"`
	End         time.Time `gorm:"column:end" json:"end" gorm-patch:"patchable"`
	CategoryID  string    `gorm:"column:category_id" json:"categoryId" gorm-patch:"patchable"`
}

type Category struct {
	ID    string `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title string `gorm:"column:titke" json:"title" gorm-patch:"patchable"`
}
