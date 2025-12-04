package service

import (
	"bistro/internal/dal"
	"bistro/models"
	"errors"
	"fmt"
	"log/slog"
)

func SaveItem(item models.InventoryItem, repo *dal.InventoryRepository) error {
	fmt.Println(item)
	if item.IngredientID == "" {
		return errors.New("ingredient_id cannot be empty")
	}
	if item.Name == "" {
		return errors.New("item name  cannot be empty")
	}
	if item.Quantity <= 0 {
		return errors.New("quantity can not be <= 0")
	}
	if item.Unit == "" {
		return errors.New("unit cannot be empty")
	}

	err := repo.SaveItem(item)
	if err != nil {
		return err
	}
	slog.Info("Item added", "id", item.IngredientID, "name", item.Name)
	return nil
}

func GetAllItems(repo *dal.InventoryRepository) ([]models.InventoryItem, error) {
	return repo.GetAllItems()
}
func GetItem(id string, repo *dal.InventoryRepository) (models.InventoryItem, error) {
	if id == "" {
		return models.InventoryItem{}, errors.New("id cannot be empty")
	}
	item, err := repo.GetItem(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

func UpdateInventoryItem(id string, repo *dal.InventoryRepository, item models.InventoryItem) (models.InventoryItem, error) {
	if id == "" {
		return models.InventoryItem{}, errors.New("id cannot be empty")
	}
	item, err := repo.UpdateInventoryItem(id, item)
	if err != nil {
		return item, err
	}
	slog.Info("item updated", "id", id)
	return item, err
}

func DeleteItem(id string, repo *dal.InventoryRepository) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	err := repo.DeleteItem(id)
	if err != nil {
		return err
	}
	slog.Info("Item deleted", "id", id)
	return nil
}

func GetLeftOvers(pageInt int, pageSizeInt int, repo *dal.InventoryRepository) ([]models.InventoryItem, error) {

	return []models.InventoryItem{}, nil
}
