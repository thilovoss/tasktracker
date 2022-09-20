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
		return nil, tx.Error
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

func DeleteTask(taskID string) error {
	tx := database.Instance.Where("id = ?", taskID).Delete(&models.Task{})
	return tx.Error
}

func GetAllCategories() (*[]models.Category, error) {
	var results []models.Category

	tx := database.Instance.Find(&results)

	return &results, tx.Error
}

func GetCategoryById(ID string) (*models.Category, error) {
	var result models.Category

	tx := database.Instance.Where("id = ?", ID).First(&result)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, tx.Error
}

func CreateCategory(category *models.Category) (*string, error) {
	tx := database.Instance.Create(category)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &category.ID, tx.Error
}

func PatchCategory(categoryID string, patches *[]gormpatch.JsonPatch) (*models.Category, error) {
	category, err := GetCategoryById(categoryID)
	if err != nil {
		return nil, err
	}

	err = database.Instance.Transaction(func(tx *gorm.DB) error {
		for index := range *patches {
			patch := (*patches)[index]
			err := gormpatch.ApplyPatch(&category, &patch)
			if err != nil {
				return err
			}
		}

		err := tx.First(&category).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return category, err
}

func DeleteCategory(categoryID string) error {
	tx := database.Instance.Where("id = ?", categoryID).Delete(&models.Category{})
	return tx.Error
}
