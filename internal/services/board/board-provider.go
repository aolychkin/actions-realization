package board

import (
	"actions/internal/lib/logger/sl"
	"actions/internal/storage"
	"context"
	"errors"
	"fmt"
	"log/slog"

	board_v1 "github.com/aolychkin/actions-contract/gen/go/board"
	"github.com/mpvl/unique"
)

func (b *Board) GetBoard(ctx context.Context, id string) (*board_v1.TBoard, error) {
	//Иницифализируемм логгер, добавляем полезной инфой
	const op = "Services.Board.Provider.GetBoard"
	log := b.log.With(
		slog.String("op", op),
		slog.String("BoardID", id),
	)
	log.Info("getting board data")

	// Получаем данные борды
	board, err := b.boardProvider.Board(ctx, id)

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
	log.Info("board data get successfully")

	columns := []*board_v1.TColumn{}
	var sprint_names []string
	for _, column := range board.OnBoardColumns {
		steps := []*board_v1.TCurrentStep{}
		for _, step := range column.Steps {
			steps = append(steps, &board_v1.TCurrentStep{
				Id:             step.ID,
				Name:           step.Name,
				WorkflowStatus: step.WorkflowStatusID,
			})
		}

		cards := []*board_v1.TCard{}
		for _, card := range column.OnBoardActions {
			var sprints []string
			//TODO: поправить ссанину и путаницу со спринтами (прото, инициализация, получение данных)
			for _, sprint := range card.Action.Sprints {
				sprints = append(sprints, sprint.ID)
				sprint_names = append(sprint_names, sprint.Name)
			}

			fields := []*board_v1.TActionField{}
			for _, field := range card.Action.Fields {
				fields = append(fields, &board_v1.TActionField{
					Id:       field.ID,
					Value:    field.Value,
					ConfigId: field.FieldConfigID,
				})
			}

			cards = append(cards, &board_v1.TCard{
				Id:            card.Action.ID,
				Order:         uint32(card.Order),
				ColumnId:      card.OnBoardColumnID,
				ActionNum:     uint32(card.Action.ActionNum),
				CurrentStepId: card.Action.CurrentStepID,
				SprintIds:     sprints,
				Fields:        fields,
			})
		}

		columns = append(columns, &board_v1.TColumn{
			Id:            column.ID,
			Name:          column.Name,
			Steps:         steps,
			OnBoardAction: cards,
		})
	}

	//TODO: убрать задвоение спринтов (в инициализации бд)
	sprints := []*board_v1.TSprint{}
	unique.Strings(&sprint_names)
	for _, sprint := range sprint_names {
		sprints = append(sprints, &board_v1.TSprint{
			Name: sprint,
		})
	}

	cardConfigs := []*board_v1.TCardConfig{}
	var fieldConfigIDs []string
	for _, cardConfig := range board.CardConfigs {
		cardConfigs = append(cardConfigs, &board_v1.TCardConfig{
			Id:            cardConfig.ID,
			RowOrder:      uint32(cardConfig.RowOrder),
			ColumnOrder:   uint32(cardConfig.ColumnOrder),
			Size:          uint32(cardConfig.Size),
			FieldConfigId: cardConfig.FieldConfigID,
		})
		fieldConfigIDs = append(fieldConfigIDs, cardConfig.FieldConfigID)
	}

	//TODO: закинуть этот код в работу со стором)
	fieldConfigs := []*board_v1.TFieldConfig{}
	unique.Strings(&fieldConfigIDs)
	fConfigs, err := b.boardProvider.FieldConfigByIDArray(ctx, fieldConfigIDs)
	if err != nil {
		log.Error("failed to getting fieldConfigs", sl.Err(err))
		return &board_v1.TBoard{}, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("fieldConfigs get successfully")

	for _, fieldConfig := range fConfigs {
		fieldConfigs = append(fieldConfigs, &board_v1.TFieldConfig{
			Id:    fieldConfig.ID,
			Name:  fieldConfig.Name,
			Alias: fieldConfig.Alias,
			FieldType: &board_v1.TFieldType{
				Id:             fieldConfig.FieldType.ID,
				Name:           fieldConfig.FieldType.Name,
				Alias:          fieldConfig.FieldType.Alias,
				IsCustom:       fieldConfig.FieldType.IsCustom,
				AvailableSizes: []string{fieldConfig.FieldType.AvailableSizes},
			},
			DefaultValue:    fieldConfig.DefaultValue,
			AvailableValues: fieldConfig.AvailableValues,
		})

	}

	// response := &board_v1.TBoard{
	// 	Id: board.ID,
	// 	Key: "CS",
	// 	TColumn: []board_v1.TColumn{
	// 		{
	// 			Id: board.,
	// 		}
	// 	}
	// }

	return &board_v1.TBoard{
		Id:           board.ID,
		Key:          "CS",
		Columns:      columns,
		Sprints:      sprints,
		FieldConfigs: fieldConfigs,
		CardConfigs:  cardConfigs,
	}, nil
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
