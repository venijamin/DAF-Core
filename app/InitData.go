package main

import (
	"DAF-Core/app/model/dto"
	"DAF-Core/app/repository"
)

func InitData() {
	boardRepository := repository.BoardRepository{}
	itemRepository := repository.ItemRepository{}

	boardDTO := dto.CreateBoard{
		ThemeUUID: "f6e09753-03d4-4820-beff-6d713298589c",
		Name:      "name",
	}
	boardUUID := boardRepository.Create(boardDTO)
	itemDTO1 := dto.CreateItem{
		ParentUUID:  "",
		BoardUUID:   boardUUID,
		Name:        "name1",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}

	item1UUID := itemRepository.Create(itemDTO1)
	itemDTO2 := dto.CreateItem{
		ParentUUID:  item1UUID,
		BoardUUID:   boardUUID,
		Name:        "name2",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}
	item2UUID := itemRepository.Create(itemDTO2)
	itemDTO3 := dto.CreateItem{
		ParentUUID:  item2UUID,
		BoardUUID:   boardUUID,
		Name:        "name3",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
	}
	itemRepository.Create(itemDTO3)

}
