package inits

import (
	"actions/internal/domain/models"
	"log/slog"

	"gorm.io/gorm"
)

func InitTestDB(db *gorm.DB) {
	const op = "cmd.actions.inits.InitTestDB"
	log := slog.With(
		slog.String("op", op),
	)
	log.Info("[START] Test DB initialization...")
	//https://gorm.io/docs/migration.html
	db.Migrator().DropTable(
		&models.Board{},
		&models.OnBoardColumn{},
		&models.OnBoardAction{},
		&models.Sprint{},
		&models.CurrentStep{},
		&models.WorkflowStatus{},
		&models.Action{},
		&models.ActionField{},
		&models.FieldConfig{},
		&models.FieldType{},
		&models.Activity{},
		&models.Project{},
		&models.CardConfig{},
	)
	db.AutoMigrate(
		&models.Board{},
		&models.OnBoardColumn{},
		&models.OnBoardAction{},
		&models.Sprint{},
		&models.CurrentStep{},
		&models.WorkflowStatus{},
		&models.Action{},
		&models.ActionField{},
		&models.FieldConfig{},
		&models.FieldType{},
		&models.Activity{},
		&models.Project{},
		&models.CardConfig{},
	)
	log.Info("[SUCCESS] Test DB initialization")
}

func TestSeed(db *gorm.DB) {
	const op = "cmd.actions.inits.InitTestDB"
	log := slog.With(
		slog.String("op", op),
	)
	log.Info("[START] Test Data seeding...")

	workflowStatus := []*models.WorkflowStatus{
		{Name: "Беклог"},
		{Name: "Начальный"},
		{Name: "В процессе"},
		{Name: "На паузе"},
		{Name: "Завершен"},
		{Name: "Отменен"},
	}
	db.Create(workflowStatus)

	project := models.Project{
		Name: "Create System",
		Key:  "CS",
	}
	db.Create(&project)

	board := models.Board{
		ProjectID: project.ID,
		OnBoardColumns: []models.OnBoardColumn{
			{
				Name: "Входящие",
				Steps: []*models.CurrentStep{
					{
						Name:             "Открытые",
						WorkflowStatusID: workflowStatus[0].ID,
					},
					{
						Name:             "Нераспределенные",
						WorkflowStatusID: workflowStatus[1].ID,
					},
				},
			},
			{
				Name: "В работе",
				Steps: []*models.CurrentStep{
					{
						Name:             "В работе",
						WorkflowStatusID: workflowStatus[2].ID,
					},
				},
			},
		},
		Sprints: []models.Sprint{
			{Name: "2025.01.01-2025.01.14"},
			{Name: "2025.01.15-2025.01.29"},
		},
	}
	db.Create(&board)

	fieldType := []*models.FieldType{
		{
			Name:           "text-inline",
			Alias:          "Строка",
			IsCustom:       false,
			AvailableSizes: "12",
		},
		{
			Name:           "text-multiline",
			Alias:          "Многострочный текст",
			IsCustom:       false,
			AvailableSizes: "12",
		},
	}
	db.Create(fieldType)

	fieldConfig := []*models.FieldConfig{
		{
			Name:            "summary",
			Alias:           "Заголовок",
			FieldTypeID:     fieldType[0].ID,
			DefaultValue:    "Укажите описание задачи",
			AvailableValues: "[a-Z]",
			IsPrimary:       true,
		},
		{
			Name:            "description",
			Alias:           "Описание",
			FieldTypeID:     fieldType[0].ID,
			DefaultValue:    "Укажите описание задачи",
			AvailableValues: "[a-Z]",
			IsPrimary:       true,
		},
	}
	db.Create(fieldConfig)

	actions := []*models.Action{
		{
			ActionNum:     1,
			Type:          "Задача",
			ProjectID:     project.ID,
			CurrentStepID: board.OnBoardColumns[0].Steps[0].ID,
			Sprints: []*models.Sprint{
				&board.Sprints[0],
				&board.Sprints[1],
			},
			Fields: []models.ActionField{
				{
					Value:         "Открытое действие 1",
					FieldConfigID: fieldConfig[0].ID,
				},
				{
					Value:         "It's description",
					FieldConfigID: fieldConfig[1].ID,
				},
			},
		},
		{
			ActionNum:     2,
			Type:          "Задача",
			ProjectID:     project.ID,
			CurrentStepID: board.OnBoardColumns[0].Steps[1].ID,
			Sprints: []*models.Sprint{
				&board.Sprints[0],
			},
			Fields: []models.ActionField{
				{
					Value:         "Нераспределенное действие 2",
					FieldConfigID: fieldConfig[0].ID,
				},
				{
					Value:         "Let's DO IT!",
					FieldConfigID: fieldConfig[1].ID,
				},
			},
		},
		{
			ActionNum:     3,
			Type:          "Задача",
			ProjectID:     project.ID,
			CurrentStepID: board.OnBoardColumns[1].Steps[0].ID,
			Sprints: []*models.Sprint{
				&board.Sprints[1],
			},
			Fields: []models.ActionField{
				{
					Value:         "В работе действие 2",
					FieldConfigID: fieldConfig[0].ID,
				},
				{
					Value:         "Simple task in work",
					FieldConfigID: fieldConfig[1].ID,
				},
			},
		},
	}
	db.Create(actions)

	onBoardActions := []*models.OnBoardAction{
		{
			Order:           1,
			OnBoardColumnID: board.OnBoardColumns[0].ID,
			ActionID:        actions[0].ID,
		},
		{
			Order:           1,
			OnBoardColumnID: board.OnBoardColumns[0].ID,
			ActionID:        actions[1].ID,
		},
		{
			Order:           1,
			OnBoardColumnID: board.OnBoardColumns[1].ID,
			ActionID:        actions[2].ID,
		},
	}
	db.Create(onBoardActions)

	cardConfigs := []*models.CardConfig{
		{
			RowOrder:      1,
			ColumnOrder:   1,
			Size:          12,
			BoardID:       board.ID,
			FieldConfigID: fieldConfig[0].ID,
		},
		{
			RowOrder:      2,
			ColumnOrder:   1,
			Size:          12,
			BoardID:       board.ID,
			FieldConfigID: fieldConfig[1].ID,
		},
	}
	db.Create(cardConfigs)

	log.Info("[SUCCESS] Test Data seeding")
}

// OnBoardActions: []models.OnBoardAction{
// 	{
// 		Order: 1,
// 		Action: models.Action{
// 			ActionNum: 1,
// 			Summary:   "Действие 1",
// 			Type:      "Задача",
// 			ProjectID: project.ID,
// 			CurrentStepID: ,
// 		},
// 	},
// },

// OnBoardColumns: []models.OnBoardColumn{
// 	{
// 		Name: "Входящие",
// 		Steps: []*models.CurrentStep{
// 			{
// 				Name:             "Входящие",
// 				WorkflowStatusID: workflowStatus[1].ID,
// 			},
// 		},
// 	},
// 	{
// 		Name: "В работе",
// 		Steps: []*models.CurrentStep{
// 			{
// 				Name:             "В работе",
// 				WorkflowStatusID: workflowStatus[2].ID,
// 			},
// 		},
// 	},
// },

// Actions: []models.Action{
// 	{
// 		ActionNum: 1,
// 		Summary: "Действие 1",
// 		Type: "Задача",
// 	},
// },
