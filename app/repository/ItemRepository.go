package repository

import (
	"DAF-Core/app/model"
	"DAF-Core/app/model/dto"
	"DAF-Core/app/util"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemRepository struct{}

func (r ItemRepository) GetAllByBoard(boardUUID string) ([]model.Item, error) {
	if boardUUID == "" {
		return nil, fmt.Errorf("board UUID cannot be empty")
	}

	var items []model.Item
	db := util.GetMainDB()

	result := db.Where("board_uuid = ?", boardUUID).Find(&items)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve items for board %s: %w", boardUUID, result.Error)
	}

	// If no items found, return empty slice rather than nil
	if len(items) == 0 {
		return []model.Item{}, nil
	}

	return items, nil
}

func (r ItemRepository) Get(uuid string) (*model.Item, error) {
	var item model.Item
	db := util.GetMainDB()

	// Check if item exists
	result := db.Where("item_uuid = ?", uuid).First(&item)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("item with UUID %s not found", uuid)
		}
		return nil, fmt.Errorf("database error retrieving item: %w", result.Error)
	}

	// Load associations in a single transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		// Load parents
		if err := tx.Model(&item).Association("Parents").Find(&item.Parents); err != nil {
			return fmt.Errorf("error loading parents: %w", err)
		}

		// Load children
		if err := tx.Model(&item).Association("Children").Find(&item.Children); err != nil {
			return fmt.Errorf("error loading children: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r ItemRepository) Delete(uuid string) error {
	if uuid == "" {
		return fmt.Errorf("item UUID cannot be empty")
	}

	db := util.GetMainDB()

	// Find the item first to check if it exists
	var item model.Item
	result := db.Where("item_uuid = ?", uuid).First(&item)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("item with UUID %s not found", uuid)
		}
		return fmt.Errorf("failed to find item %s: %w", uuid, result.Error)
	}

	// Delete the item
	result = db.Delete(&item)
	if result.Error != nil {
		return fmt.Errorf("failed to delete item %s: %w", uuid, result.Error)
	}

	// Check if the item was actually deleted
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected when deleting item %s", uuid)
	}

	return nil
}

func (r ItemRepository) Create(dto dto.CreateItem) (string, error) {
	db := util.GetMainDB()
	item := model.Item{
		ItemUUID:    uuid.NewString(),
		BoardUUID:   dto.BoardUUID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Tags:        dto.Tags,
		Picture:     dto.Picture,
		Barcode:     dto.Barcode,
		Fields:      dto.Fields,
	}

	return item.ItemUUID, db.Transaction(func(tx *gorm.DB) error {
		// Create main item
		if err := tx.Create(&item).Error; err != nil {
			return err
		}

		// Handle parent relationships
		if len(dto.ParentUUIDs) > 0 {
			var parents []*model.Item
			if err := tx.Where("item_uuid IN ?", dto.ParentUUIDs).Find(&parents).Error; err != nil {
				return err
			}
			if len(parents) != len(dto.ParentUUIDs) {
				return fmt.Errorf("some parent items not found")
			}
			if err := tx.Model(&item).Association("Parents").Append(parents); err != nil {
				return err
			}
		}

		// Handle child relationships
		if len(dto.ChildUUIDs) > 0 {
			var children []*model.Item
			if err := tx.Where("item_uuid IN ?", dto.ChildUUIDs).Find(&children).Error; err != nil {
				return err
			}
			if len(children) != len(dto.ChildUUIDs) {
				return fmt.Errorf("some child items not found")
			}
			if err := tx.Model(&item).Association("Children").Append(children); err != nil {
				return err
			}
		}

		return nil
	})
}
