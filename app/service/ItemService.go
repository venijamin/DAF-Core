package service

import (
	"DAF-Core/app/model"
	"DAF-Core/app/model/dto"
	"DAF-Core/app/repository"
	"github.com/google/uuid"
	"log"
)

type ItemService struct{}

var itemRepository repository.ItemRepository

func (s *ItemService) Get(uuid string) (*model.Item, error) {
	item, err := itemRepository.Get(uuid)
	if err != nil {
		log.Printf("Failed to retrieve item %s: %v", uuid, err)
		return nil, err
	}
	return item, nil
}

func (s *ItemService) GetAllByBoard(boardUUID string) ([]model.Item, error) {
	items, err := itemRepository.GetAllByBoard(boardUUID)
	if err != nil {
		log.Printf("Failed to retrieve items for board %s: %v", boardUUID, err)
		return nil, err
	}
	return items, nil
}

func (s *ItemService) Create(dto dto.CreateItem) (string, error) {
	item := model.Item{
		ItemUUID:    uuid.New().String(),
		BoardUUID:   dto.BoardUUID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Tags:        dto.Tags,
		Picture:     dto.Picture,
		Barcode:     dto.Barcode,
		Fields:      dto.Fields,
	}

	_, err := itemRepository.Create(item, dto.ParentUUIDs, dto.ChildUUIDs)
	if err != nil {
		log.Printf("Failed to create item %s: %v", item.ItemUUID, err)
		return "", err
	}

	return item.ItemUUID, nil
}

func (s *ItemService) Delete(uuid string) error {
	err := itemRepository.Delete(uuid)
	if err != nil {
		log.Printf("Failed to delete item %s: %v", uuid, err)
		return err
	}
	return nil
}

func (s *ItemService) Update(dto dto.CreateItem, uuid string) (string, error) {
	item := model.Item{
		ItemUUID:    uuid,
		BoardUUID:   dto.BoardUUID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Tags:        dto.Tags,
		Picture:     dto.Picture,
		Barcode:     dto.Barcode,
		Fields:      dto.Fields,
	}

	_, err := itemRepository.Update(item, dto.ParentUUIDs, dto.ChildUUIDs)
	if err != nil {
		log.Printf("Failed to create item %s: %v", item.ItemUUID, err)
	}

	return item.ItemUUID, nil
}
