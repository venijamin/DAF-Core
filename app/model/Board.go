package model

type Board struct {
	BoardUUID string `gorm:"primaryKey;type:uuid"`

	ThemeUUID string `gorm:"type:uuid"`
	Theme     Theme  `gorm:"foreignKey:ThemeUUID;references:ThemeUUID"`

	Items []Item `gorm:"foreignKey:BoardUUID;references:BoardUUID"`

	Name string
}
