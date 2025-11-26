package handler

import (
	"bistro/internal/dal"
	"bistro/internal/service"
	"bistro/models"
	"encoding/json"
	"net/http"
)

func AddMenuItem(w http.ResponseWriter, r *http.Request, menuRepo *dal.MenuRepository) {
	menu := models.MenuItem{}
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}

	err = service.AddMenuItem(menuRepo, menu)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Status Internal Server Error", err.Error())
		return
	}

	addedMenu, err := service.GetMenuItem(menuRepo, menu.ID)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedMenu)
}

func GetMenuAllItems(w http.ResponseWriter, r *http.Request, menuRepo *dal.MenuRepository) {

	menuItems, err := service.GetMenuAllItems(menuRepo)
	if err != nil {
		sendError(w, http.StatusNotFound, "StatusNotFound", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menuItems)
}

func GetMenuItem(w http.ResponseWriter, r *http.Request, menuRepo *dal.MenuRepository, id string) {

	menuItem, err := service.GetMenuItem(menuRepo, id)
	if err != nil {
		sendError(w, http.StatusNotFound, "StatusNotFound", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menuItem)

}

func UpdateMenuItem(w http.ResponseWriter, r *http.Request, menuRepo *dal.MenuRepository, id string) {
	var menuItem models.MenuItem

	err := json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {
		sendError(w, http.StatusBadRequest, "StatusBadRequest", err.Error())
		return
	}
	err = service.UpdateMenuItem(menuRepo, id, menuItem)
	if err != nil {
		sendError(w, http.StatusNotFound, "StatusNotFound", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menuItem)
}

func DeleteMenuItem(w http.ResponseWriter, r *http.Request, menuRepo *dal.MenuRepository, id string) {

	err := service.DeleteMenuItem(id, menuRepo)
	if err != nil {
		sendError(w, http.StatusNotFound, "StatusNotFound", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
