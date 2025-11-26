package dal

import (
	"bistro/models"
	"database/sql"
	"fmt"
	"log/slog"
)

type InventoryRepository struct {
	conn *sql.DB
}

func NewInventoryRepository(conn *sql.DB) *InventoryRepository {
	return &InventoryRepository{
		conn: conn,
	}
}

func (r *InventoryRepository) SaveItem(item models.InventoryItem) error {
	query := `INSERT INTO inventory (ingredient_id,name,quantity,unit) VALUES($1,$2,$3,$4)`

	slog.Info("Saving item", "ingredient_id", item.IngredientID, "name", item.Name)

	result, err := r.conn.Exec(query, item.IngredientID, item.Name, item.Quantity, item.Unit)
	if err != nil {
		slog.Error("Save error", "error", err)
		return err
	}

	rows, _ := result.RowsAffected()
	slog.Info("Item saved", "rows_affected", rows)

	return nil
}
func (r *InventoryRepository) GetAllItems() ([]models.InventoryItem, error) {
	query := `SELECT id, ingredient_id, name, quantity, unit FROM inventory`
	rows, err := r.conn.Query(query)
	if err != nil {
		slog.Error("Query error", "error", err) // ← ДОБАВЬ ЭТО
		return nil, err
	}
	defer rows.Close()

	var items []models.InventoryItem

	for rows.Next() {
		var item models.InventoryItem
		var id int

		err := rows.Scan(&id, &item.IngredientID, &item.Name, &item.Quantity, &item.Unit)
		if err != nil {
			slog.Error("Scan error", "error", err) // ← ДОБАВЬ ЭТО
			return nil, err
		}
		items = append(items, item)
	}

	slog.Info("Items fetched", "count", len(items)) // ← ДОБАВЬ ЭТО

	if err = rows.Err(); err != nil {
		slog.Error("Rows error", "error", err) // ← ДОБАВЬ ЭТО
		return nil, err
	}

	return items, nil
}

func (r *InventoryRepository) GetItem(id string) (models.InventoryItem, error) {
	var item models.InventoryItem
	var idFromDb int
	quary := `
	SELECT id, ingredient_id, name, quantity, unit FROM inventory WHERE ingredient_id = $1
	`
	row := r.conn.QueryRow(quary, id)
	err := row.Scan(&idFromDb, &item.IngredientID, &item.Name, &item.Quantity, &item.Unit)
	if err != nil {
		return models.InventoryItem{}, err
	}
	fmt.Println(item)
	return item, nil
}

func (r *InventoryRepository) UpdateInventoryItem(id string, item models.InventoryItem) (models.InventoryItem, error) {
	query := `
	UPDATE inventory SET  name = $1, quantity =$2, unit = $3 WHERE ingredient_id  = $4
	`
	_, err := r.conn.Exec(query, item.Name, item.Quantity, item.Unit, item.IngredientID)
	if err != nil {
		return models.InventoryItem{}, err
	}
	return item, nil
}

func (r *InventoryRepository) DeleteItem(id string) error {
	query := `
	DELETE FROM inventory WHERE ingredient_id = $1
	`
	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
	// 1. Напиши SQL запрос DELETE WHERE ingredient_id
	// 2. Выполни r.conn.Exec(query, id)
	// 3. Проверь error и верни его
}
