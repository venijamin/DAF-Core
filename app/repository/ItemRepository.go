package repository

import (
	"DAF-Core/app/model"
	"DAF-Core/app/model/dto"
	"DAF-Core/app/util"
	"github.com/google/uuid"
)

type ItemRepository struct{}

func (r ItemRepository) Get(uuid string) model.Item {
	var item model.Item
	util.GetMainDB().Where("item_uuid = ?", uuid).First(&item)
	return item
}
func (r ItemRepository) Create(dto dto.CreateItem) {
	item := model.Item{
		ItemUUID:    uuid.NewString(),
		ParentUUID:  dto.ParentUUID,
		BoardUUID:   dto.BoardUUID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Tags:        dto.Tags,
		Picture:     dto.Picture,
		Barcode:     dto.Barcode,
		Fields:      dto.Fields,
	}
	util.GetMainDB().Create(&item)
}
func (r ItemRepository) Update(dto dto.CreateItem, uuid string) {
	item := model.Item{
		ItemUUID:    uuid,
		ParentUUID:  dto.ParentUUID,
		BoardUUID:   dto.BoardUUID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Tags:        dto.Tags,
		Picture:     dto.Picture,
		Barcode:     dto.Barcode,
		Fields:      dto.Fields,
	}
	util.GetMainDB().Save(&item)
}
func (r ItemRepository) Delete(uuid string) {
	util.GetMainDB().Where("item_uuid = ?", uuid).Delete(&model.Item{})
}
