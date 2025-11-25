package dal

import (
	"bistro/models"
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
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
	filepath := "/inventory.json"

	// Шаг 1: Прочитать весь файл в []byte
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Шаг 2: JSON → Go структура (unmarshal)
	var items []models.InventoryItem
	err = json.Unmarshal(data, &items)
	if err != nil {
		return err
	}

	// Шаг 3: Добавить новый item в массив
	items = append(items, item)

	// Шаг 4: Go структура → JSON (marshal)
	newData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	// Шаг 5: Записать обратно в файл
	err = os.WriteFile(filepath, newData, 0644)
	if err != nil {
		return err
	}

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
	filePath := "/inventory.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return models.InventoryItem{}, err
	}
	var items []models.InventoryItem
	json.Unmarshal(data, &items)
	for _, item := range items {
		if id == item.IngredientID {
			return item, nil
		}
	}

	return models.InventoryItem{}, errors.New("item not found")
}

func (r *InventoryRepository) UpdateInventoryItem(id string, item models.InventoryItem) (models.InventoryItem, error) {
	filepath := "/inventory.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return item, err
	}
	var items []models.InventoryItem
	var newItems []models.InventoryItem
	isFound := false
	err = json.Unmarshal(file, &items)
	if err != nil {
		return item, err
	}

	for _, ite := range items {
		if ite.IngredientID == id {
			newItems = append(newItems, item)
			isFound = true
			continue
		}
		newItems = append(newItems, ite)

	}
	if isFound == false {
		return item, errors.New("id not found in inventory")
	}
	f, err := json.Marshal(newItems)
	if err != nil {
		return item, err
	}
	err = os.WriteFile(filepath, f, 0666)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *InventoryRepository) DeleteItem(id string) error {
	filepath := "/inventory.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var items []models.InventoryItem
	err = json.Unmarshal(file, &items)
	if err != nil {
		return err
	}
	var newItems []models.InventoryItem
	isExsist := false
	for _, item := range items {
		if item.IngredientID == id {
			isExsist = true
			continue
		}
		newItems = append(newItems, item)
	}
	if !isExsist {
		return errors.New("item is not Exists")
	}
	f, err := json.Marshal(newItems)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, f, 0666)
	if err != nil {
		return err
	}

	return nil
}
