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
	var orderId int
	query := `
	INSERT INTO orders (order_id, customer_name, status) 
	VALUES($1,$2,$3)
	RETURNING id
	`
	id := r.conn.QueryRow(query, order.ID, order.CustomerName, order.Status)
	err := id.Scan(&orderId)
	if err != nil {
		return err
	}
	for _, item := range order.Items {
		query = `
	INSERT INTO order_items (order_id, menu_item_id, quantity)
	VALUES($1,$2,$3)
	`
		_, err := r.conn.Exec(query, orderId, item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
	}
	return nil

	// filepathMenu := "/menu.json"
	// filePathInventory := "/inventory.json"
	// menuFile, err := os.ReadFile(filepathMenu)
	// if err != nil {
	// 	return err
	// }
	// inventoryFile, err := os.ReadFile(filePathInventory)
	// if err != nil {
	// 	return err
	// }
	// var MenuStruct []models.MenuItem
	// var InventoryStruct []models.InventoryItem
	// err = json.Unmarshal(menuFile, &MenuStruct)
	// if err != nil {
	// 	return err
	// }

	// err = json.Unmarshal(inventoryFile, &InventoryStruct)
	// if err != nil {
	// 	return err
	// }
	// var nedeedIngridients []models.MenuItemIngredient
	// for _, orderItemsStrict := range order.Items {
	// 	for _, menuStructPart := range MenuStruct {
	// 		if orderItemsStrict.ProductID == menuStructPart.ID {
	// 			// теперь у нас есть рецепт это menuStructPart
	// 			// мне нужно теперь взять из него ингредиенты и умножить на колличество
	// 			for _, menuIngridient := range menuStructPart.Ingredients {
	// 				// мы получили меню айтем то есть конкретный рецепт
	// 				// теперь нужно
	// 				menuIngridient.Quantity = menuIngridient.Quantity * float64(orderItemsStrict.Quantity)
	// 				nedeedIngridients = append(nedeedIngridients, menuIngridient)

	// 			}

	// 		}
	// 	}
	// }
	// ingredientMap := make(map[string]float64)
	// for _, ing := range nedeedIngridients {
	// 	ingredientMap[ing.IngredientID] += ing.Quantity
	// }

	// nedeedIngridients = []models.MenuItemIngredient{}
	// for ingID, qty := range ingredientMap {
	// 	nedeedIngridients = append(nedeedIngridients, models.MenuItemIngredient{
	// 		IngredientID: ingID,
	// 		Quantity:     qty,
	// 	})
	// }
	// var NewInventoryStruct []models.InventoryItem
	// // я собрал все нужные мне ингридиенты в nedeedIngridients
	// // теперь нужно удалить их из инвенторя
	// for _, neededIngridientPart := range nedeedIngridients {
	// 	for _, InventoryStructPart := range InventoryStruct {
	// 		if InventoryStructPart.IngredientID == neededIngridientPart.IngredientID {

	// 			InventoryStructPart.Quantity = InventoryStructPart.Quantity - neededIngridientPart.Quantity
	// 			if InventoryStructPart.Quantity < 0 {
	// 				return errors.New("not enough" + InventoryStructPart.IngredientID)
	// 			}
	// 			NewInventoryStruct = append(NewInventoryStruct, InventoryStructPart)
	// 		}

	// 	}
	// }

	// var inventoryStructToApend []models.InventoryItem
	// // теперь мне нужно заполнить в структуру инвенторя обновленные данные
	// for _, inventoryPart := range InventoryStruct {
	// 	for _, newInventoryPart := range NewInventoryStruct {
	// 		if inventoryPart.IngredientID == newInventoryPart.IngredientID {
	// 			inventoryPart = models.InventoryItem{
	// 				IngredientID: inventoryPart.IngredientID,
	// 				Name:         inventoryPart.Name,
	// 				Quantity:     newInventoryPart.Quantity,
	// 				Unit:         inventoryPart.Unit,
	// 			}

	// 		}

	// 	}
	// 	inventoryStructToApend = append(inventoryStructToApend, inventoryPart)
	// }

	// inventoryFileMarshal, err := json.Marshal(inventoryStructToApend)
	// if err != nil {
	// 	return err
	// }
	// err = os.WriteFile(filePathInventory, inventoryFileMarshal, 0666)

	// if err != nil {
	// 	return err
	// }

	// var orders []models.Order
	// orderFile, err := os.ReadFile("/orders.json")
	// if err != nil {
	// 	return err
	// }
	// err = json.Unmarshal(orderFile, &orders)
	// orders = append(orders, order)
	// orderFileMarshal, err := json.Marshal(orders)
	// if err != nil {
	// 	return err
	// }

	// err = os.WriteFile("/orders.json", orderFileMarshal, 0666)
	// if err != nil {
	// 	return err
	// }
	// return nil
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
	JOIN order_items ri ON ri.order_id = r.id
	
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
WHERE r.order_id = $1
`
	rows, err := r.conn.Query(query, id)
	if err != nil {
		return models.Order{}, err
	}
	var order models.Order
	var orderItem models.OrderItem

	for rows.Next() {
		rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.CreatedAt, &orderItem.ProductID, &orderItem.Quantity)
		order.Items = append(order.Items, orderItem)

	}

	return order, nil
}

func (r *OrdersRepository) UpdateOrderById(id string, status models.OrderStatus) error {
	query := `
	UPDATE orders SET status = $1 
	WHERE order_id = $2
	`
	_, err := r.conn.Exec(query, status.Status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersRepository) DeleteOrder(id string) error {
	query := `
	DELETE FROM order_items
	WHERE order_id IN(
	SELECT id FROM orders WHERE order_id = $1
	)
	`
	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}
	query = `
	DELETE FROM orders
	WHERE order_id = $1
	`
	_, err = r.conn.Exec(query, id)
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
