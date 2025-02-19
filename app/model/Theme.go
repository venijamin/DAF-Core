package model

type Theme struct {
	ThemeUUID string `json:"theme_uuid" gorm:"primaryKey"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Picture   string `json:"picture"`
}
