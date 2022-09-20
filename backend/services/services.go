package services

import (
	"tasktracker/database"
	"tasktracker/gormpatch"
	"tasktracker/models"

	"gorm.io/gorm"
)

func GetAllTasks() (*[]models.Task, error) {
	var results []models.Task

	tx := database.Instance.Find(&results)

	return &results, tx.Error
}

func GetTaskByID(ID string) (*models.Task, error) {
	var result models.Task

	tx := database.Instance.Where("id = ?", ID).First(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, tx.Error
}

func CreateTask(task *models.Task) (*string, error) {
	tx := database.Instance.Create(task)
	if tx.Error != nil {
		return &task.ID, tx.Error
	}
	return &task.ID, tx.Error
}

func PatchTask(taskID string, patches *[]gormpatch.JsonPatch) (*models.Task, error) {
	task, err := GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	err = database.Instance.Transaction(func(tx *gorm.DB) error {
		for index := range *patches {
			patch := (*patches)[index]
			err := gormpatch.ApplyPatch(&task, &patch)
			if err != nil {
				return err
			}
		}

		err := tx.First(&task).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return task, err
}

func DeleteTask(ID string) error {
	tx := database.Instance.Where("id = ?", ID).Delete(&models.Task{})
	return tx.Error
}
