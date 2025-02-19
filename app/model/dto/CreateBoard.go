package dto

type CreateBoard struct {
	ThemeUUID string `json:"theme_uuid"`
	Name      string `json:"name"`
}
