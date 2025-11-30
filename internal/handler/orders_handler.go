package handler

import (
	"bistro/internal/dal"
	"bistro/internal/service"
	"bistro/models"
	"encoding/json"
	"net/http"
)

func PostOrder(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		sendError(w, http.StatusBadRequest, "StatusBadRequest", err.Error())
		return
	}
	err = service.PostOrder(order, ordersRepo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)

}

func GetAllOrders(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository) {
	orders, err := service.GetAllOrders(ordersRepo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func GetOrderById(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository, id string) {
	order, err := service.GetOrderById(ordersRepo, id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	json.NewEncoder(w).Encode(order)
}

func UpdateOrderById(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository, id string) {
	var status models.OrderStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	err = service.UpdateOrderById(ordersRepo, id, status)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	order, err := service.GetOrderById(ordersRepo, id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository, id string) {

	err := service.DeleteOrder(id, ordersRepo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func CloseOrders(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository, id string) {
	err := service.CloseOrders(id, ordersRepo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
