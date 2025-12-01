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

func GetPopularItems(ordersRepo *dal.OrdersRepository) ([]models.PopularItems, error) {
	popularItems, err := ordersRepo.GetPopularItems()
	if err != nil {
		return []models.PopularItems{}, err
	}
	return popularItems, nil
}
