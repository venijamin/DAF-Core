package repository

import (
	"DAF-Core/app/model"
	"DAF-Core/app/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BoardRepository struct{}

func (r BoardRepository) GetAll() ([]model.Board, error) {
	var boards []model.Board
	db := util.GetMainDB()

	result := db.Find(&boards)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve boards: %w", result.Error)
	}

	// If no boards found, return empty slice rather than nil
	if len(boards) == 0 {
		return []model.Board{}, nil
	}

	return boards, nil
}

func (r BoardRepository) Get(uuid string) (*model.Board, error) {
	if uuid == "" {
		return nil, fmt.Errorf("board UUID cannot be empty")
	}

	var board model.Board
	db := util.GetMainDB()

	result := db.Where("board_uuid = ?", uuid).First(&board)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("board with UUID %s not found", uuid)
		}
		return nil, fmt.Errorf("failed to retrieve board %s: %w", uuid, result.Error)
	}

	return &board, nil
}

func (r BoardRepository) Create(board model.Board) (string, error) {
	// Save to database
	db := util.GetMainDB()

	board.DateCreated = time.Now().String()
	board.DateLastModified = time.Now().String()

	result := db.Create(&board)
	if result.Error != nil {
		return "", fmt.Errorf("failed to create board: %w", result.Error)
	}

	// Verify board was created
	if result.RowsAffected == 0 {
		return "", fmt.Errorf("board was not created, no rows affected")
	}

	return board.BoardUUID, nil
}

func (r BoardRepository) Update(board model.Board, uuid string) error {
	db := util.GetMainDB()

	board.DateLastModified = time.Now().String()

	// Check if board exists
	var existingBoard model.Board
	result := db.Where("board_uuid = ?", uuid).First(&existingBoard)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("board with UUID %s not found", uuid)
		}
		return fmt.Errorf("failed to find board %s: %w", uuid, result.Error)
	}

	// Update board fields
	existingBoard.ThemeUUID = board.ThemeUUID
	existingBoard.Name = board.Name

	// Save updated board
	result = db.Save(&existingBoard)
	if result.Error != nil {
		return fmt.Errorf("failed to update board %s: %w", uuid, result.Error)
	}

	// Verify update was successful
	if result.RowsAffected == 0 {
		return fmt.Errorf("board %s was not updated, no rows affected", uuid)
	}

	return nil
}

func (r BoardRepository) Delete(uuid string) error {
	if uuid == "" {
		return fmt.Errorf("board UUID cannot be empty")
	}

	db := util.GetMainDB()

	// Find the board first to check if it exists
	var board model.Board
	result := db.Where("board_uuid = ?", uuid).First(&board)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("board with UUID %s not found", uuid)
		}
		return fmt.Errorf("failed to find board %s: %w", uuid, result.Error)
	}

	// Delete the board
	result = db.Delete(&board)
	if result.Error != nil {
		return fmt.Errorf("failed to delete board %s: %w", uuid, result.Error)
	}

	// Check if the board was actually deleted
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected when deleting board %s", uuid)
	}

	return nil
}
