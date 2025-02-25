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
	boardUUID, _ := boardRepository.Create(boardDTO)
	itemDTO1 := dto.CreateItem{
		BoardUUID:   boardUUID,
		ParentUUIDs: nil,
		ChildUUIDs:  nil,
		Name:        "name1",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}

	item1UUID, _ := itemRepository.Create(itemDTO1)
	itemDTO2 := dto.CreateItem{
		BoardUUID:   boardUUID,
		ParentUUIDs: []string{item1UUID},
		ChildUUIDs:  nil,
		Name:        "name2",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}
	item2UUID, _ := itemRepository.Create(itemDTO2)
	itemDTO3 := dto.CreateItem{
		BoardUUID:   boardUUID,
		ParentUUIDs: []string{item2UUID},
		ChildUUIDs:  []string{item1UUID},
		Name:        "name3",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}
	itemRepository.Create(itemDTO3)
	itemDTO4 := dto.CreateItem{
		BoardUUID:   boardUUID,
		Name:        "name4",
		Description: "description",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}
	itemRepository.Create(itemDTO4)
}
