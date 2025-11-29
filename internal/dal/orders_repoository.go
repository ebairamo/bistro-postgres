package dal

import (
	"bistro/models"
	"database/sql"
	"encoding/json"
	"errors"
	"os"
)

type OrdersRepository struct {
	conn *sql.DB
}

func NewOrdersRepository(conn *sql.DB) *OrdersRepository {
	return &OrdersRepository{
		conn: conn,
	}
}

func (r *OrdersRepository) PostOrder(order models.Order) error {
	filepathMenu := "/menu.json"
	filePathInventory := "/inventory.json"
	menuFile, err := os.ReadFile(filepathMenu)
	if err != nil {
		return err
	}
	inventoryFile, err := os.ReadFile(filePathInventory)
	if err != nil {
		return err
	}
	var MenuStruct []models.MenuItem
	var InventoryStruct []models.InventoryItem
	err = json.Unmarshal(menuFile, &MenuStruct)
	if err != nil {
		return err
	}

	err = json.Unmarshal(inventoryFile, &InventoryStruct)
	if err != nil {
		return err
	}
	var nedeedIngridients []models.MenuItemIngredient
	for _, orderItemsStrict := range order.Items {
		for _, menuStructPart := range MenuStruct {
			if orderItemsStrict.ProductID == menuStructPart.ID {
				// теперь у нас есть рецепт это menuStructPart
				// мне нужно теперь взять из него ингредиенты и умножить на колличество
				for _, menuIngridient := range menuStructPart.Ingredients {
					// мы получили меню айтем то есть конкретный рецепт
					// теперь нужно
					menuIngridient.Quantity = menuIngridient.Quantity * float64(orderItemsStrict.Quantity)
					nedeedIngridients = append(nedeedIngridients, menuIngridient)

				}

			}
		}
	}
	ingredientMap := make(map[string]float64)
	for _, ing := range nedeedIngridients {
		ingredientMap[ing.IngredientID] += ing.Quantity
	}

	nedeedIngridients = []models.MenuItemIngredient{}
	for ingID, qty := range ingredientMap {
		nedeedIngridients = append(nedeedIngridients, models.MenuItemIngredient{
			IngredientID: ingID,
			Quantity:     qty,
		})
	}
	var NewInventoryStruct []models.InventoryItem
	// я собрал все нужные мне ингридиенты в nedeedIngridients
	// теперь нужно удалить их из инвенторя
	for _, neededIngridientPart := range nedeedIngridients {
		for _, InventoryStructPart := range InventoryStruct {
			if InventoryStructPart.IngredientID == neededIngridientPart.IngredientID {

				InventoryStructPart.Quantity = InventoryStructPart.Quantity - neededIngridientPart.Quantity
				if InventoryStructPart.Quantity < 0 {
					return errors.New("not enough" + InventoryStructPart.IngredientID)
				}
				NewInventoryStruct = append(NewInventoryStruct, InventoryStructPart)
			}

		}
	}

	var inventoryStructToApend []models.InventoryItem
	// теперь мне нужно заполнить в структуру инвенторя обновленные данные
	for _, inventoryPart := range InventoryStruct {
		for _, newInventoryPart := range NewInventoryStruct {
			if inventoryPart.IngredientID == newInventoryPart.IngredientID {
				inventoryPart = models.InventoryItem{
					IngredientID: inventoryPart.IngredientID,
					Name:         inventoryPart.Name,
					Quantity:     newInventoryPart.Quantity,
					Unit:         inventoryPart.Unit,
				}

			}

		}
		inventoryStructToApend = append(inventoryStructToApend, inventoryPart)
	}

	inventoryFileMarshal, err := json.Marshal(inventoryStructToApend)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePathInventory, inventoryFileMarshal, 0666)

	if err != nil {
		return err
	}

	var orders []models.Order
	orderFile, err := os.ReadFile("/orders.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(orderFile, &orders)
	orders = append(orders, order)
	orderFileMarshal, err := json.Marshal(orders)
	if err != nil {
		return err
	}

	err = os.WriteFile("/orders.json", orderFileMarshal, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersRepository) GetAllOrders() ([]models.Order, error) {
	query := `
	SELECT 
	r.order_id, 
	r.customer_name, 
	r.status,
	r.created_at,
	ri.menu_item_id, 
	ri.quantity
	
	
	FROM orders r
	LEFT JOIN order_items ri ON r.id = ri.order_id
	ORDER BY r.order_id
	`

	rows, err := r.conn.Query(query)
	if err != nil {
		return []models.Order{}, err
	}

	var order models.Order

	var orders []models.Order
	var preOrder models.Order
	for rows.Next() {
		var orderItem models.OrderItem
		err := rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.CreatedAt, &orderItem.ProductID, &orderItem.Quantity)
		if err != nil {
			return []models.Order{}, err
		}
		if preOrder.ID == "" {
			preOrder = order
		}
		if order.ID != preOrder.ID && preOrder.ID != "" {
			orders = append(orders, preOrder)
			preOrder = order
		}
		preOrder.Items = append(preOrder.Items, orderItem)
	}
	if preOrder.ID != "" {
		orders = append(orders, preOrder)
	}
	return orders, nil
}

func (r *OrdersRepository) GetOrderById(id string) (models.Order, error) {

	filepath := "/orders.json"
	var orders []models.Order
	file, err := os.ReadFile(filepath)
	if err != nil {
		return models.Order{}, err
	}
	err = json.Unmarshal(file, &orders)
	if err != nil {
		return models.Order{}, err
	}
	for _, order := range orders {
		if order.ID == id {
			return order, nil
		}
	}
	return models.Order{}, errors.New("order not found")

}

func (r *OrdersRepository) UpdateOrderById(id string, status models.OrderStatus) (models.Order, error) {
	filepath := "/orders.json"
	var orders []models.Order
	var newOrders []models.Order
	file, err := os.ReadFile(filepath)
	if err != nil {
		return models.Order{}, err
	}
	err = json.Unmarshal(file, &orders)
	if err != nil {
		return models.Order{}, err
	}
	isFound := false
	var orderToReturn models.Order
	for _, order := range orders {
		if order.ID == id {
			order.Status = status.Status
			newOrders = append(newOrders, order)
			isFound = true
			orderToReturn = order
			continue
		}
		newOrders = append(newOrders, order)
	}
	if !isFound {
		return models.Order{}, errors.New("order not found by id")
	}
	f, err := json.Marshal(newOrders)
	if err != nil {
		return models.Order{}, err
	}
	err = os.WriteFile(filepath, f, 0666)
	if err != nil {
		return models.Order{}, err
	}
	return orderToReturn, nil
}

func (r *OrdersRepository) DeleteOrder(id string) error {
	filepath := "/orders.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var orders []models.Order
	var newOrders []models.Order
	isFound := false
	err = json.Unmarshal(file, &orders)
	if err != nil {
		return err
	}
	for _, order := range orders {
		if order.ID == id {
			isFound = true
			continue
		}
		newOrders = append(newOrders, order)
	}
	if !isFound {
		return errors.New("order not found")
	}
	f, err := json.Marshal(newOrders)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, f, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersRepository) CloseOrders(id string) error {
	filepath := "/orders.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var orders []models.Order
	var newOrders []models.Order
	isExist := false
	err = json.Unmarshal(file, &orders)
	if err != nil {
		return err
	}
	for _, order := range orders {
		if order.ID == id {
			order.Status = "close"
			newOrders = append(newOrders, order)
			isExist = true
			continue
		}
		newOrders = append(newOrders, order)
	}
	if !isExist {
		return errors.New("order not found")
	}
	f, err := json.Marshal(newOrders)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, f, 0666)
	if err != nil {
		return err
	}
	return nil
}
