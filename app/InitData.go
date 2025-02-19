package main

import (
	"DAF-Core/app/model/dto"
	"DAF-Core/app/repository"
)

func InitData() {
	//boardRepository := repository.BoardRepository{}
	itemRepository := repository.ItemRepository{}

	itemDTO1 := dto.CreateItem{
		ParentUUID:  "",
		BoardUUID:   "f6e09753-03d4-4820-beff-6d713298589c",
		Name:        "name1",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}
	itemDTO2 := dto.CreateItem{
		ParentUUID:  "f6e09753-03d4-4820-beff-6d713298589c",
		BoardUUID:   "f6e09753-03d4-4820-beff-6d713298589c",
		Name:        "name2",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}
	itemRepository.Create(itemDTO1)
	itemRepository.Create(itemDTO2)

}
