package dto

type CreateItem struct {
	BoardUUID   string   `json:"board_uuid"`
	ParentUUIDs []string `json:"parent_uuids"`
	ChildUUIDs  []string `json:"child_uuids"`

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	Tags        []string `gorm:"type:text[]"`
	Picture     string   `json:"picture"`
	Barcode     string   `json:"barcode"`
	Fields      []string `gorm:"type:text[]"`
}
