package main

import (
	"bistro/internal/dal"
	"bistro/internal/database"
	"bistro/internal/handler"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {

	flagPort := flag.Int("port", 8000, "Port number")
	flagHelp := flag.Bool("help", false, "help flag")
	flag.Parse()
	slog.Info("StartingBistro", "port", *flagPort)
	if *flagHelp {
		help()
		os.Exit(0)
	}
	slog.Info("Storage initialized")

	conn, err := database.NewDB(database.LoadConfig())
	if err != nil {
		log.Fatal(err)
	}
	repo := dal.NewInventoryRepository(conn)
	menuRepo := dal.NewMenuRepository(conn)
	ordersRepo := dal.NewOrdersRepository(conn)
	addr := fmt.Sprintf(":%d", *flagPort)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		inventoryHandler(w, r, repo, menuRepo, ordersRepo)
	})

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("Server failed", "error", err)
	}
}

func inventoryHandler(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository, menuRepo *dal.MenuRepository, ordersRepo *dal.OrdersRepository) {
	url := strings.Split(r.URL.Path, "/")
	switch url[1] {
	case "inventory":
		if len(url) == 2 {
			switch r.Method {
			case http.MethodPost:
				handler.AddInventoryItem(w, r, repo)
			case http.MethodGet:
				handler.GetAllItems(w, r, repo)
			}
		}
		if len(url) == 3 {
			switch r.Method {
			case http.MethodGet:
				if url[2] == "getLeftOvers" {
					handler.GetLeftOvers(w, r, repo)
				} else {
					handler.GetItem(w, r, repo)
				}

			case http.MethodPut:
				handler.UpdateInventoryItem(w, r, repo)
			case http.MethodDelete:
				handler.DeleteItem(w, r, repo)
			}
		}
	case "menu":
		if len(url) == 2 {
			switch r.Method {
			case http.MethodPost:
				handler.AddMenuItem(w, r, menuRepo)
			case http.MethodGet:
				handler.GetMenuAllItems(w, r, menuRepo)
			}
		}
		if len(url) == 3 {
			switch r.Method {
			case http.MethodGet:
				handler.GetMenuItem(w, r, menuRepo, url[2])
			case http.MethodPut:
				handler.UpdateMenuItem(w, r, menuRepo, url[2])
			case http.MethodDelete:
				handler.DeleteMenuItem(w, r, menuRepo, url[2])
			}
		}
	case "orders":
		if len(url) == 2 {
			switch r.Method {
			case http.MethodPost:
				handler.PostOrder(w, r, ordersRepo)
			case http.MethodGet:
				handler.GetAllOrders(w, r, ordersRepo)
			}
		}
		if len(url) == 3 {
			switch r.Method {
			case http.MethodGet:
				if url[2] == "numberOfOrderedItems" {
					handler.NumberOfOrderedItems(w, r, ordersRepo)
				} else {
					handler.GetOrderById(w, r, ordersRepo, url[2])
				}
			case http.MethodPut:
				handler.UpdateOrderById(w, r, ordersRepo, url[2])
			case http.MethodDelete:
				handler.DeleteOrder(w, r, ordersRepo, url[2])
			case http.MethodPost:
				handler.CloseOrders(w, r, ordersRepo, url[2])
			}
		}
	case "reports":
		if len(url) == 3 {
			switch r.Method {
			case http.MethodGet:
				if url[2] == "total-sales" {
					handler.GetTotalSales(w, r, ordersRepo)
				}
				if url[2] == "popular-items" {
					handler.GetPopularItems(w, r, ordersRepo)
				}
			}
		}
	}
}

func help() {
	fmt.Println(`$ ./bistro --help
Bistro Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.`)
}
