package persistence

import (
	"automatic-trade/backend/domain/model"
	"automatic-trade/backend/domain/repository"
	"automatic-trade/backend/infra/rdb/dto"

	"gorm.io/gorm"
)

type positionRepository struct {
	db *gorm.DB
}

func (p *positionRepository) Delete(orderID string) error {
	err := p.db.Where("order_id = ?", orderID).Delete(&dto.Position{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *positionRepository) Get(orderID string) (model.Position, error) {
	var position dto.Position

	err := p.db.Where("order_id = ?", orderID).First(&position).Error
	if err != nil {
		return model.Position{}, nil
	}

	result, err := position.ToModel()
	if err != nil {
		return model.Position{}, nil
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
