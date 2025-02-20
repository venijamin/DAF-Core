package model

type Item struct {
	ItemUUID string `gorm:"primaryKey;type:uuid"`

	BoardUUID string `gorm:"type:uuid"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Quantity    int
	Tags        []string `gorm:"type:text[]"`
	Picture     string
	Barcode     string
	Fields      []string `gorm:"type:text[]"`

	Parents  []Item `gorm:"many2many:item_associations;joinForeignKey:ChildUUID;joinReferences:ParentUUID"`
	Children []Item `gorm:"many2many:item_associations;joinForeignKey:ParentUUID;joinReferences:ChildUUID"`
}

type ItemAssociation struct {
	ParentUUID string `gorm:"type:uuid;primaryKey"`
	ChildUUID  string `gorm:"type:uuid;primaryKey"`
}
