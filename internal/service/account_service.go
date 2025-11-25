package service

import (
	"github.com/zarishsphere/zarish-fin/internal/models"
	"github.com/zarishsphere/zarish-fin/internal/repository"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateJournalEntry(move *models.AccountMove) error {
	// Add business logic here (e.g., validation, auto-numbering)
	if move.State == "" {
		move.State = "draft"
	}
	return s.repo.CreateMove(move)
}

func (s *AccountService) GetJournalEntry(id uint) (*models.AccountMove, error) {
	return s.repo.GetMoveByID(id)
}

func (s *AccountService) ListJournalEntries(page, pageSize int) ([]models.AccountMove, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListMoves(offset, pageSize)
}
