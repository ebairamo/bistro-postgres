package dal

import (
	"bistro/models"
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
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
	query := `
	INSERT INTO menu_items (product_id, name, description, price) VALUES ($1,$2,$3,$4)
	`
	_, err := r.conn.Exec(query, menuItem.ID, menuItem.Name, menuItem.Description, menuItem.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *MenuRepository) GetMenuAllItems() ([]models.MenuItem, error) {

	query := `
		SELECT id,product_id, name, description, price FROM menu_items 
		`
	rows, err := r.conn.Query(query)
	if err != nil {
		slog.Error("Query error", "error", err)
		return []models.MenuItem{}, nil
	}
	defer rows.Close()
	var items []models.MenuItem
	var ingredients []models.MenuItemIngredient
	for rows.Next() {
		var item models.MenuItem
		var id int
		err := rows.Scan(&id, &item.ID, &item.Name, &item.Description, &item.Price)
		if err != nil {
			return []models.MenuItem{}, err
		}
		item.Ingredients = ingredients
		items = append(items, item)
	}

	return items, nil
}

func (r *MenuRepository) GetMenuItem(id string) (models.MenuItem, error) {
	var item models.MenuItem
	var idDb int
	query := `
	SELECT id,product_id, name, description, price FROM menu_items WHERE product_id = $1
	`
	row := r.conn.QueryRow(query, id)

	err := row.Scan(&idDb, &item.ID, &item.Name, &item.Description, &item.Price)
	if err != nil {
		return models.MenuItem{}, err
	}
	// SELECT ... WHERE product_id = $1
	// QueryRow() (не Query!)
	// Scan()
	// return item
	return item, nil
}

func (r *MenuRepository) UpdateMenuItem(id string, menuItem models.MenuItem) error {
	query := `
	UPDATE menu_items SET  name = $1, description = $2, price = $3 WHERE product_id = $4
	`
	_, err := r.conn.Exec(query, menuItem.Name, menuItem.Description, menuItem.Price, menuItem.ID)
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
