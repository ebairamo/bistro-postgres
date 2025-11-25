package service

import (
	"bistro/internal/dal"
	"bistro/models"
)

func GetTotalSales(ordersRepo *dal.OrdersRepository) (models.TotalSales, error) {
	totalSales, err := ordersRepo.GetTotalSales()
	if err != nil {
		return models.TotalSales{}, err
	}
	return totalSales, nil
}

func GetPopularItems(ordersRepo *dal.OrdersRepository) (map[string]int, error) {
	popularItems, err := ordersRepo.GetPopularItems()
	if err != nil {
		return map[string]int{}, err
	}
	return popularItems, nil
}
