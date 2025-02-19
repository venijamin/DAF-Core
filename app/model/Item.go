package model

type Item struct {
	ItemUUID string `json:"item_uuid" gorm:"primaryKey"`

	ParentID *string
	Items    []Item `gorm:"foreignKey:ParentID;references:ItemUUID"`

	BoardUUID string `json:"board_uuid"`

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	Tags        []string `gorm:"type:text[]"`
	Picture     string   `json:"picture"`
	Barcode     string   `json:"barcode"`
	Fields      []string `gorm:"type:text[]"`
}
