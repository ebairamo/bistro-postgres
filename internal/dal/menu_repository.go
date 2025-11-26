package dal

import (
	"bistro/models"
	"database/sql"
	"log/slog"
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
	var ing models.MenuItemIngredient
	query := `
        SELECT 
        m.id,
        m.product_id, 
        m.name, 
        m.description, 
        m.price,
        i.ingredient_id,
        mi.quantity
        FROM menu_items m
        LEFT JOIN menu_item_ingredients mi ON m.id = mi.menu_item_id
        LEFT JOIN inventory i ON mi.ingredient_id = i.id
        ORDER BY m.id
		`
	rows, err := r.conn.Query(query)
	if err != nil {
		slog.Error("Query error", "error", err)
		return []models.MenuItem{}, nil
	}
	defer rows.Close()

	var items []models.MenuItem
	var preItem models.MenuItem

	for rows.Next() {
		var Item models.MenuItem
		var id int
		var quantity sql.NullFloat64
		var ingId sql.NullString

		err := rows.Scan(&id, &Item.ID, &Item.Name, &Item.Description, &Item.Price, &ingId, &quantity)
		if err != nil {
			return []models.MenuItem{}, err
		}

		if quantity.Valid {
			ing.Quantity = quantity.Float64
		}
		if ingId.Valid {
			ing.IngredientID = ingId.String
		}

		Item.Ingredients = append(Item.Ingredients, ing)

		if preItem.ID != Item.ID && preItem.ID != "" {
			items = append(items, preItem)
		}

		preItem = Item
	}

	if preItem.ID != "" {
		items = append(items, preItem)
	}

	return items, nil
}
func (r *MenuRepository) GetMenuItem(id string) (models.MenuItem, error) {
	query := `
SELECT 
    m.id, 
    m.product_id, 
    m.name, 
    m.description, 
    m.price, 
    i.ingredient_id,
    mi.quantity
FROM menu_items m
LEFT JOIN menu_item_ingredients mi ON m.id = mi.menu_item_id
LEFT JOIN inventory i ON mi.ingredient_id = i.id
WHERE m.product_id = $1
    `
	rows, err := r.conn.Query(query, id)
	if err != nil {
		return models.MenuItem{}, err
	}
	var menuItem models.MenuItem
	var ingredient models.MenuItemIngredient
	var menuItemId int

	for rows.Next() {

		rows.Scan(&menuItemId, &menuItem.ID, &menuItem.Name, &menuItem.Description, &menuItem.Price, &ingredient.IngredientID, &ingredient.Quantity)

		menuItem.Ingredients = append(menuItem.Ingredients, ingredient)

	}

	return menuItem, nil
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
	query := `
	DELETE FROM menu_items WHERE product_id = $1
	`
	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
	// r.conn.Exec()
	// check err
	// return err
}
