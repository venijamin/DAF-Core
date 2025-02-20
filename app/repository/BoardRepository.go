package repository

import (
	"DAF-Core/app/model"
	"DAF-Core/app/model/dto"
	"DAF-Core/app/util"
	"github.com/google/uuid"
)

type BoardRepository struct{}

func (r BoardRepository) GetAll() []model.Board {
	var boards []model.Board
	util.GetMainDB().Find(&boards)
	return boards
}

func (r BoardRepository) Get(uuid string) model.Board {
	var board model.Board
	util.GetMainDB().Where("board_uuid = ?", uuid).First(&board)
	return board
}
func (r BoardRepository) Create(dto dto.CreateBoard) string {
	boardUUID := uuid.NewString()
	board := model.Board{
		BoardUUID: boardUUID,
		ThemeUUID: dto.ThemeUUID,
		Name:      dto.Name,
	}
	util.GetMainDB().Create(&board)
	return boardUUID
}
func (r BoardRepository) Update(dto dto.CreateBoard, uuid string) {
	board := model.Board{
		BoardUUID: uuid,
		ThemeUUID: dto.ThemeUUID,
		Name:      dto.Name,
	}
	util.GetMainDB().Save(&board)
}
func (r BoardRepository) Delete(uuid string) {
	util.GetMainDB().Where("board_uuid = ?", uuid).Delete(&model.Board{})
}
