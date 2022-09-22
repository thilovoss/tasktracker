package models

import (
	"time"
)

type Category struct {
	ID       string  `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title    string  `gorm:"column:title" json:"title" gorm-patch:"patchable"`
	ParentID *string `json:"parentId" gorm-patch:"patchable"`
	Tasks    []Task  `gorm:"foreignKey:category_id"`
}

type Task struct {
	ID          string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"column:name" json:"name" gorm-patch:"patchable"`
	Description string    `gorm:"column:description" json:"description" gorm-patch:"patchable"`
	Start       time.Time `gorm:"column:start" json:"start" gorm-patch:"patchable"`
	End         time.Time `gorm:"column:end" json:"end" gorm-patch:"patchable"`
	CategoryID  string    `json:"categoryId" gorm-patch:"patchable"`
}
