package dal

import (
	"bistro/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type InventoryRepository struct {
	dataDir string
}

func NewInventoryRepository(dataDir string) *InventoryRepository {
	return &InventoryRepository{
		dataDir: dataDir,
	}
}

func (r *InventoryRepository) SaveItem(item models.InventoryItem) error {
	filepath := r.dataDir + "/inventory.json"

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
	filepath := r.dataDir + "/inventory.json"
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	inventoryItem := []models.InventoryItem{}
	err = json.Unmarshal(data, &inventoryItem)
	if err != nil {
		return nil, err
	}
	fmt.Println(inventoryItem)
	// TODO: прочитать файл
	// TODO: распарсить JSON
	// TODO: вернуть массив

	return inventoryItem, nil
}

func (r *InventoryRepository) GetItem(id string) (models.InventoryItem, error) {
	filePath := r.dataDir + "/inventory.json"
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
	filepath := r.dataDir + "/inventory.json"
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
	filepath := r.dataDir + "/inventory.json"
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
