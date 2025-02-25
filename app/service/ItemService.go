package service

import (
	"DAF-Core/app/model"
	"DAF-Core/app/repository"
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
