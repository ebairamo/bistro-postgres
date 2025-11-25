package handler

import (
	"bistro/internal/dal"
	"bistro/internal/service"
	"encoding/json"
	"net/http"
)

func GetTotalSales(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository) {

	totalSales, err := service.GetTotalSales(ordersRepo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(totalSales)
}

func GetPopularItems(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository) {
	popularItems, err := service.GetPopularItems(ordersRepo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(popularItems)
}
