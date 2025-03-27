package sqlite

import (
	"actions/internal/domain/models"
	"actions/internal/lib/logger/resp"
	"actions/internal/storage"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Storage) Board(
	ctx context.Context,
	BoardID string,
) (models.Board, error) {
	const op = "storage.sqlite.providers.Board"
	var board models.Board

	err := s.db.Model(&models.Board{}).Preload("OnBoardColumns.OnBoardActions.Action.Sprints").Preload("OnBoardColumns.Steps").Preload(clause.Associations).Preload("OnBoardColumns.OnBoardActions.Action.Fields").Preload(clause.Associations).First(&board).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Board{}, fmt.Errorf("%s: %w", op, storage.ErrBoardNotFound)
	}

	resp.PrintResp(board)

	return board, nil
}

func (s *Storage) FieldConfigByIDArray(
	ctx context.Context,
	IDs []string,
) ([]models.FieldConfig, error) {
	const op = "storage.sqlite.providers.FieldConfigByIDArray"
	var fieldConfigs []models.FieldConfig
	for _, fieldConfigID := range IDs {
		var fieldConfig models.FieldConfig
		err := s.db.Model(&models.FieldConfig{}).Preload("FieldType").First(&fieldConfig, "id = ?", fieldConfigID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.FieldConfig{}, fmt.Errorf("%s: %w", op, storage.ErrBoardNotFound)
		}
		fieldConfigs = append(fieldConfigs, fieldConfig)
	}

	return fieldConfigs, nil
}
