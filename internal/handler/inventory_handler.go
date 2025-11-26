package handler

import (
	"bistro/internal/dal"
	"bistro/internal/service"
	"bistro/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AddInventoryItem(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {

	var item models.InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Status Internal Server Error", err.Error())
		return
	}
	fmt.Println(item)
	err = service.SaveItem(item, repo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Status Internal Server Error", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
	// TODO: проверить что метод POST

	// TODO: прочитать и распарсить JSON

	// TODO: вызвать service для сохранения

	// TODO: вернуть успешный ответ
}

func GetAllItems(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	items, err := service.GetAllItems(repo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	url := strings.Split(r.URL.Path, "/")
	if len(url) == 3 {
		item, err := service.GetItem(url[2], repo)
		if err != nil {
			fmt.Println(err)
			sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
			return
		}
		json.NewEncoder(w).Encode(item)
	}
}

func UpdateInventoryItem(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	url := strings.Split(r.URL.Path, "/")
	var item models.InventoryItem
	json.NewDecoder(r.Body).Decode(&item)
	if len(url) == 3 {
		itemUpdated, err := service.UpdateInventoryItem(url[2], repo, item)
		if err != nil {
			sendError(w, http.StatusNotFound, "StatusNotFound", err.Error())
			return
		}
		err = json.NewEncoder(w).Encode(itemUpdated)
		if err != nil {
			sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
			return
		}
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	url := strings.Split(r.URL.Path, "/")
	if len(url) == 3 {
		err := service.DeleteItem(url[2], repo)
		if err != nil {
			sendError(w, http.StatusNotFound, "StatusNotFound", err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
