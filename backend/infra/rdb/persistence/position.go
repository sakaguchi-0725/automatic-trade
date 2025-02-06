package persistence

import (
	"automatic-trade/backend/core/apperr"
	"automatic-trade/backend/domain/model"
	"automatic-trade/backend/domain/repository"
	"automatic-trade/backend/infra/rdb/dto"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type positionRepository struct {
	db *gorm.DB
}

func (p *positionRepository) Delete(orderID string) error {
	result := p.db.Where("order_id = ?", orderID).Delete(&dto.Position{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("position with orderID %s not found", orderID)
	}

	return nil
}

func (p *positionRepository) Get(orderID string) (model.Position, error) {
	var position dto.Position

	err := p.db.Where("order_id = ?", orderID).First(&position).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Position{}, apperr.ErrDataNotFound
		}
		return model.Position{}, err
	}

	result, err := position.ToModel()
	if err != nil {
		return model.Position{}, err
	}

	return result, err
}

func (p *positionRepository) Store(position model.Position) error {
	data := dto.NewPosition(position)
	if err := p.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func NewPositionRepository(db *gorm.DB) repository.Position {
	return &positionRepository{db}
}
