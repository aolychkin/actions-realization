package board

import (
	"actions/internal/lib/logger/sl"
	"actions/internal/storage"
	"context"
	"errors"
	"fmt"
	"log/slog"

	board_v1 "github.com/aolychkin/actions-contract/gen/go/board"
)

func (b *Board) GetBoard(ctx context.Context, id string) (*board_v1.TBoard, error) {
	//Иницифализируемм логгер, добавляем полезной инфой
	const op = "Board.GetBoard"
	log := b.log.With(
		slog.String("op", op),
		slog.String("BoardID", id),
	)
	log.Info("getting board data")

	// Получаем данные борды
	_, err := b.boardProvider.Board(ctx, id)

	if err != nil {
		if errors.Is(err, storage.ErrBoardNotFound) {
			log.Warn("invalid board id", sl.Err(err))

			// return &board_v1.GetBoardResponse{}, fmt.Errorf("%s: %w", op, ErrInvalidBoardID)
			return &board_v1.TBoard{}, fmt.Errorf("%s: %w", op, ErrInvalidBoardID)
		}

		log.Error("failed to getting board data", sl.Err(err))
		// return &board_v1.GetBoardResponse{}, fmt.Errorf("%s: %w", op, err)
		return &board_v1.TBoard{}, fmt.Errorf("%s: %w", op, err)
	}

	// var columns []*board_v1.TColumn
	// var cards []*board_v1.TCard
	// var configs []*board_v1.TFieldConfig
	// var types []*board_v1.TFieldType
	// for _, column := range board.Columns {
	// 	columns = append(columns, &board_v1.TColumn{
	// 		Id:    column.ID,
	// 		Title: column.Name,
	// 	})
	// 	for _, card := range board.Actions {
	// 		var cardFields []*board_v1.TCardField
	// 		for _, field := range card.Fields {
	// 			cardFields = append(cardFields, &board_v1.TCardField{
	// 				Id:       field.ID,
	// 				ConfigID: field.Config.ID,
	// 				Values:   field.Value,
	// 			})
	// 			configs = append(configs, &board_v1.TFieldConfig{
	// 				Id:              field.Config.ID,
	// 				Name:            field.Config.Name,
	// 				Alias:           field.Config.Alias,
	// 				ValueTypeID:     field.Config.Type.ID,
	// 				DefaultValue:    field.Config.DefaultValue,
	// 				ValueSource:     field.Config.ValueSource,
	// 				AvailableValues: field.Config.AvailableValues,
	// 			})

	// 		}

	// 		cards = append(cards, &board_v1.TCard{
	// 			Id: card.ID,
	// 			CardMeta: &board_v1.TCardMeta{
	// 				BoardID:   board.ID,
	// 				ColumnID:  column.ID,
	// 				Order:     int32(card.Order),
	// 				CreatedAt: card.CreatedAt.String(),
	// 				CreatedBy: "",
	// 				UpdatedAt: card.UpdatedAt.String(),
	// 				UpdatedBy: "",
	// 				Type:      card.Type,
	// 				Project:   "",
	// 				Sprint:    "",
	// 			},
	// 			CardFields: cardFields,
	// 		})
	// 	}
	// }

	log.Info("fund config added successfully")
	return &board_v1.TBoard{}, nil
	// return &board_v1.TBoard{
	// 	Id:           board.ID,
	// 	Key:          board.Key,
	// 	Columns:      columns,
	// 	Cards:        cards,
	// 	FieldConfigs: board,
	// 	FieldTypes:   board.ID,
	// }, err
	// return &board_v1.GetBoardResponse{
	// 	Id:           board.ID,
	// 	Columns:      board.Columns,
	// 	Cards:        board.Cards,
	// 	FieldConfigs: board.FieldConfigs,
	// 	FieldTypes:   board.FieldTypes,
	// }, nil
}
