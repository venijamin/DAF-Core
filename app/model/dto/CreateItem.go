package dto

type CreateItem struct {
	ParentUUID string `json:"parent_uuid"`

	BoardUUID string `json:"board_uuid"`

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	Tags        []string `gorm:"type:text[]"`
	Picture     string   `json:"picture"`
	Barcode     string   `json:"barcode"`
	Fields      []string `gorm:"type:text[]"`
}
