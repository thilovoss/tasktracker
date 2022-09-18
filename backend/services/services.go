package services

import (
	"tasktracker/database"
	"tasktracker/models"
)

func GetAllTasks() (*[]models.Task, error) {
	var results []models.Task

	tx := database.Instance.Find(&results)

	return &results, tx.Error
}

func CreateTask(task *models.Task) (*string, error) {
	tx := database.Instance.Create(task)
	if tx.Error != nil {
		return &task.ID, tx.Error
	}
	return &task.ID, tx.Error
}

func DeleteTask(ID string) error {
	tx := database.Instance.Where("id = ?", ID).Delete(&models.Task{})
	return tx.Error
}
