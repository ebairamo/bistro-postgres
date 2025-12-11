package service

import (
	"bistro/internal/dal"
	"bistro/models"
	"errors"
)

func PostOrder(order models.Order, ordersRepo *dal.OrdersRepository) error {
	if order.ID == "" {
		return errors.New("order id cannot be empty")
	}
	if order.CustomerName == "" {
		return errors.New("CustomerName cannot be empty")
	}
	if order.Status == "" {
		return errors.New("status cannot be empry")
	}
	for _, item := range order.Items {
		if item.ProductID == "" {
			return errors.New("item ProductId cannot be empty")
		}
		if item.Quantity <= 0 {
			return errors.New("item.Quantity cannot be <= 0")
		}
	}
	err := ordersRepo.PostOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func GetAllOrders(ordersRepo *dal.OrdersRepository) ([]models.Order, error) {
	orders, err := ordersRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderById(ordersRepo *dal.OrdersRepository, id string) (models.Order, error) {
	order, err := ordersRepo.GetOrderById(id)
	if err != nil {
		return models.Order{}, err
	}
	return order, err
}

func UpdateOrderById(ordersRepo *dal.OrdersRepository, id string, status models.OrderStatus) error {
	err := ordersRepo.UpdateOrderById(id, status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(id string, ordersRepo *dal.OrdersRepository) error {
	err := ordersRepo.DeleteOrder(id)
	if err != nil {
		return err
	}
	return nil
}

func CloseOrders(id string, ordersRepo *dal.OrdersRepository) error {
	err := ordersRepo.CloseOrders(id)
	if err != nil {
		return err
	}
	return nil
}

func NumberOfOrderedItems(startDate string, endDate string, ordersRepo *dal.OrdersRepository) ([]models.OrderItem, error) {
	items, err := ordersRepo.NumberOfOrderedItems(startDate, endDate)
	if err != nil {
		return []models.OrderItem{}, err
	}
	return items, nil
}
