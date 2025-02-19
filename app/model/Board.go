package model

type Board struct {
	BoardUUID string `json:"board_uuid" gorm:"primaryKey"`

	ThemeUUID string `json:"theme_uuid"`
	Theme     Theme  `gorm:"foreignKey:ThemeUUID;references:ThemeUUID"`

	Items []Item `gorm:"foreignKey:BoardUUID;references:BoardUUID"`

	Name string `json:"name"`
}
