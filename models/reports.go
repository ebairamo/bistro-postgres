package models

type TotalSales struct {
	TotalSales float64 `json:"total_sales"`
}

type PopularItems struct {
	Name     string
	Quantity int
}
