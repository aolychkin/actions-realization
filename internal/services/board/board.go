package board

import (
	"actions/internal/domain/models"
	"context"
	"errors"
	"log/slog"
)

// Взаимодействие со storage
type Board struct {
	log *slog.Logger
	// boardSaver    BoardSaver
	boardProvider BoardProvider
}

// Получение всяческих значений борды
type BoardProvider interface {
	Board(
		ctx context.Context,
		BoardID string,
	) (board models.Board, err error)
	FieldConfigByIDArray(
		ctx context.Context,
		IDs []string,
	) (board []models.FieldConfig, err error)
}

// Переменные ошибок
var (
	ErrInvalidBoardID = errors.New("invalid board id")
)

// Конструктор для сервиса
func New(
	log *slog.Logger,
	boardProvider BoardProvider,
) *Board {
	return &Board{
		log:           log,
		boardProvider: boardProvider,
	}
}
