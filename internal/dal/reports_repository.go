package dal

import (
	"bistro/models"
	"encoding/json"
	"fmt"
	"os"
)

func (r *OrdersRepository) GetTotalSales() (models.TotalSales, error) {
	filepath := r.dataDir + "/orders.json"
	filepathMenu := r.dataDir + "/menu.json"
	var orders []models.Order
	var menus []models.MenuItem
	file, err := os.ReadFile(filepath)
	if err != nil {
		return models.TotalSales{}, err
	}
	err = json.Unmarshal(file, &orders)
	if err != nil {
		return models.TotalSales{}, err
	}

	fileMenu, err := os.ReadFile(filepathMenu)
	if err != nil {
		return models.TotalSales{}, err
	}
	err = json.Unmarshal(fileMenu, &menus)
	if err != nil {
		return models.TotalSales{}, err
	}

	var totalSales models.TotalSales
	var total float64
	for _, order := range orders {
		for _, orderItems := range order.Items {
			for _, menu := range menus {

				if orderItems.ProductID == menu.ID {
					total = total + (menu.Price * float64(orderItems.Quantity))

				}
			}
		}
	}
	totalSales.TotalSales = total
	return totalSales, nil
}

func (r *OrdersRepository) GetPopularItems() (map[string]int, error) {
	filepath := r.dataDir + "/orders.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return map[string]int{}, err
	}
	var orders []models.Order
	itemCounts := make(map[string]int)
	err = json.Unmarshal(file, &orders)
	if err != nil {
		return map[string]int{}, err
	}
	for _, order := range orders {
		for _, orderItems := range order.Items {
			itemCounts[orderItems.ProductID] += orderItems.Quantity
		}
	}
	fmt.Println(itemCounts)
	return itemCounts, nil
}
