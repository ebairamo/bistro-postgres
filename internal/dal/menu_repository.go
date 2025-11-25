package dal

import (
	"bistro/models"
	"database/sql"
	"encoding/json"
	"errors"
	"os"
)

type MenuRepository struct {
	conn *sql.DB
}

func NewMenuRepository(conn *sql.DB) *MenuRepository {
	return &MenuRepository{
		conn: conn,
	}
}
func (r *MenuRepository) AddMenuItem(menuItem models.MenuItem) error {
	filepath := "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var menuItems []models.MenuItem
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return err
	}
	menuItems = append(menuItems, menuItem)

	data, err := json.Marshal(menuItems)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (r *MenuRepository) GetMenuAllItems() ([]models.MenuItem, error) {

	filepath := "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var menuItems []models.MenuItem
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (r *MenuRepository) GetMenuItem(id string) (models.MenuItem, error) {
	filepath := "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return models.MenuItem{}, err
	}
	menuItems := []models.MenuItem{}
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return models.MenuItem{}, err
	}
	for _, item := range menuItems {
		if item.ID == id {
			return item, nil
		}
	}
	return models.MenuItem{}, errors.New("menuItem not found")
}

func (r *MenuRepository) UpdateMenuItem(id string, menuItem models.MenuItem) error {
	filepath := "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var menuItems []models.MenuItem
	var newMenuItems []models.MenuItem
	isFound := false
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return err
	}
	for _, item := range menuItems {
		if item.ID == id {
			isFound = true
			newMenuItems = append(newMenuItems, menuItem)
			continue
		}
		newMenuItems = append(newMenuItems, item)
	}
	if !isFound {
		return errors.New("menu item not found")
	}
	f, err := json.Marshal(newMenuItems)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, f, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (r *MenuRepository) DeleteMenuItem(id string) error {
	filepath := "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var menuItems []models.MenuItem
	var newNenuItems []models.MenuItem
	isFound := false
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return err
	}
	for _, item := range menuItems {
		if item.ID == id {
			isFound = true
			continue
		}
		newNenuItems = append(newNenuItems, item)
	}
	if !isFound {
		return errors.New("item not found")
	}
	f, err := json.Marshal(newNenuItems)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, f, 0666)
	if err != nil {
		return err
	}
	return nil
}
