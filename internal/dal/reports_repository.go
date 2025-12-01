package dal

import (
	"bistro/models"
	"database/sql"
)

func (r *OrdersRepository) GetTotalSales() (models.TotalSales, error) {
	var total_amount sql.NullFloat64
	quary := `
	SELECT SUM(total_amount) as total_sales FROM orders WHERE status = 'close'
	`
	row := r.conn.QueryRow(quary)
	err := row.Scan(&total_amount)
	if err != nil {
		return models.TotalSales{}, err
	}

	return models.TotalSales{
		TotalSales: total_amount.Float64,
	}, nil
}

func (r *OrdersRepository) GetPopularItems() ([]models.PopularItems, error) {
	var itemCounts []models.PopularItems
	var itemCount models.PopularItems
	query := `
	SELECT m.name, SUM(oi.quantity) as total_quantity
	FROM order_items oi
	JOIN menu_items m ON oi.menu_item_id = m.id
	GROUP BY m.id, m.name
	ORDER BY total_quantity DESC
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		return []models.PopularItems{}, err
	}
	for rows.Next() {
		var name string
		var quantity int
		err = rows.Scan(&name, &quantity)
		if err != nil {
			return []models.PopularItems{}, err
		}
		itemCount.Name = name
		itemCount.Quantity = quantity
		itemCounts = append(itemCounts, itemCount)

	}
	return itemCounts, nil
}
