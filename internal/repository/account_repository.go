package repository

import (
	"github.com/zarishsphere/zarish-fin/internal/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// CreateMove creates a new account move (journal entry)
func (r *AccountRepository) CreateMove(move *models.AccountMove) error {
	return r.db.Create(move).Error
}

// GetMoveByID retrieves a move by its ID with preloaded lines
func (r *AccountRepository) GetMoveByID(id uint) (*models.AccountMove, error) {
	var move models.AccountMove
	err := r.db.Preload("Lines").First(&move, id).Error
	return &move, err
}

// UpdateMove updates an existing move
func (r *AccountRepository) UpdateMove(move *models.AccountMove) error {
	return r.db.Save(move).Error
}

// ListMoves retrieves a list of moves with pagination
func (r *AccountRepository) ListMoves(offset, limit int) ([]models.AccountMove, error) {
	var moves []models.AccountMove
	err := r.db.Offset(offset).Limit(limit).Find(&moves).Error
	return moves, err
}
