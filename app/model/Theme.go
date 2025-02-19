package model

type Theme struct {
	ThemeUUID string `gorm:"primaryKey;type:uuid"`
	Name      string
	Color     string
	Picture   string
}
