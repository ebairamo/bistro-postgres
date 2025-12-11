package models

type InventoryItem struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

type ResponseGetLeftOvers struct {
	CurrentPage int             `json:"currentPage"`
	HasNextPage bool            `json:"hasNextPage"`
	TotalPages  int             `json:"totalPages"`
	PageSize    int             `json:"pageSize"`
	Data        []InventoryItem `json:"data"`
}
