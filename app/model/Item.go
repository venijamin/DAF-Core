package model

type Item struct {
	ItemUUID string `gorm:"primaryKey;type:uuid"`

	ParentUUID string `gorm:"type:uuid"`
	Items      []Item `gorm:"foreignKey:ParentUUID;references:ItemUUID"`

	BoardUUID string `gorm:"type:uuid"`

	Name        string
	Description string
	Quantity    int
	Tags        []string `gorm:"type:text[]"`
	Picture     string
	Barcode     string
	Fields      []string `gorm:"type:text[]"`
}
