package service

import (
	"DAF-Core/app/model"
	"DAF-Core/app/model/dto"
	"DAF-Core/app/repository"
	"github.com/google/uuid"
	"log"
)

type BoardService struct{}

var boardRepository repository.BoardRepository

func (s *BoardService) Get(uuid string) (*model.Board, error) {
	board, err := boardRepository.Get(uuid)
	if err != nil {
		log.Printf("Failed to retrieve board %s: %v", uuid, err)
		return nil, err
	}
	return board, nil
}

func (s *BoardService) GetAll() (*[]model.Board, error) {
	boards, err := boardRepository.GetAll()
	if err != nil {
		log.Printf("Failed to retrieve board list: %v", err)
		return nil, err
	}
	return &boards, nil
}

func (s *BoardService) Create(dto dto.CreateBoard) (string, error) {
	board := model.Board{
		BoardUUID: uuid.New().String(),
		ThemeUUID: dto.ThemeUUID,
		Name:      dto.Name,
	}

	boardUUID, err := boardRepository.Create(board)
	if err != nil {
		log.Printf("Failed to create board: %v", err)
		return "", err
	}
	return boardUUID, nil
}

func (s *BoardService) Delete(uuid string) error {
	err := boardRepository.Delete(uuid)
	if err != nil {
		log.Printf("Failed to delete board %s: %v", uuid, err)
		return err
	}
	return nil
}
func (s *BoardService) Update(dto dto.CreateBoard, uuid string) error {
	board := model.Board{
		BoardUUID: uuid,
		ThemeUUID: dto.ThemeUUID,
		Name:      dto.Name,
	}
	err := boardRepository.Update(board, uuid)
	if err != nil {
		log.Printf("Failed to update board %s: %v", uuid, err)
		return err
	}
	return nil
}
