package sqlite

import (
	"actions/internal/domain/models"
	"actions/internal/storage"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (s *Storage) Board(
	ctx context.Context,
	BoardID string,
) (models.Board, error) {
	const op = "storage.sqlite.providers.Board"
	var board models.Board

	err := s.db.Model(&models.Board{}).Preload("Board.Projects").Preload("Board.Columns").Preload("Board.Actions").Preload("Actions.Fields").Preload("Actions.Fields.Config").First(&board, BoardID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Board{}, fmt.Errorf("%s: %w", op, storage.ErrBoardNotFound)
	}

	return board, nil
}
