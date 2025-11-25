package service

import (
	"bistro/internal/dal"
	"bistro/models"
	"errors"
	"log/slog"
)

func AddMenuItem(menu *dal.MenuRepository, menuItem models.MenuItem) error {
	if menuItem.ID == "" {
		return errors.New("ingredient_id cannot be empty")
	}
	if menuItem.Name == "" {
		return errors.New("item name  cannot be empty")
	}
	if menuItem.Price <= 0 {
		return errors.New("price can not be <= 0")
	}
	if menuItem.Description == "" {
		return errors.New("description cannot be empty")
	}
	if len(menuItem.Ingredients) == 0 {
		return errors.New("ingredients cannot be empty")
	}
	err := menu.AddMenuItem(menuItem)
	if err != nil {
		return err
	}
	slog.Info("Item added", "id", menuItem.ID, "name", menuItem.Name)
	return nil
}

func GetMenuAllItems(menu *dal.MenuRepository) ([]models.MenuItem, error) {

	return menu.GetMenuAllItems()
}

func GetMenuItem(menu *dal.MenuRepository, id string) (models.MenuItem, error) {

	return menu.GetMenuItem(id)
}

func UpdateMenuItem(menu *dal.MenuRepository, id string, menuItem models.MenuItem) error {
	if menuItem.ID == "" {
		return errors.New("ingredient_id cannot be empty")
	}
	if menuItem.Name == "" {
		return errors.New("item name  cannot be empty")
	}
	if menuItem.Price <= 0 {
		return errors.New("price can not be <= 0")
	}
	if menuItem.Description == "" {
		return errors.New("description cannot be empty")
	}
	if len(menuItem.Ingredients) == 0 {
		return errors.New("ingredients cannot be empty")
	}

	err := menu.UpdateMenuItem(id, menuItem)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMenuItem(id string, menuRepo *dal.MenuRepository) error {
	err := menuRepo.DeleteMenuItem(id)
	if err != nil {
		return err
	}
	return nil
}
