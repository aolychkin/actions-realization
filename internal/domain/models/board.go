package models

// TODO: как сделаю норм модель, то можно объединить Gorm и proto файл
// Список доступных полей
type Board struct {
	Base
	ProjectID      string `gorm:"type:uuid;"`
	OnBoardColumns []OnBoardColumn
	Sprints        []Sprint
	CardConfigs    []CardConfig
}

type OnBoardColumn struct {
	Base
	Name           string
	BoardID        string         `gorm:"type:uuid;"`
	Steps          []*CurrentStep `gorm:"many2many:column_steps;"`
	OnBoardActions []OnBoardAction
}

// TODO: При перетаскивании карточки назначается первый в списке статус проекта карточки действия
type CurrentStep struct {
	Base
	Name             string
	WorkflowStatusID string
	Actions          []Action         //`gorm:"many2many:action_steps;"`
	OnBoardColumns   []*OnBoardColumn `gorm:"many2many:column_steps;"`
}
type WorkflowStatus struct {
	Base
	Name  string
	Steps []*CurrentStep
}

type OnBoardAction struct { // Создавать в БД в самом конце
	Base
	Order           uint
	OnBoardColumnID string `gorm:"type:uuid;"`
	ActionID        string `gorm:"type:uuid;"`
	Action          Action
}

// Ваще это настройка глобальная
// Спринт - это тупо фильтр по дате. Воспринимаю так
type Sprint struct {
	Base
	Name    string
	Actions []*Action `gorm:"many2many:action_sprints;"`
	BoardID string    `gorm:"type:uuid;"`
}

type CardConfig struct {
	Base
	RowOrder      uint
	ColumnOrder   uint
	Size          uint
	BoardID       string `gorm:"type:uuid;"`
	FieldConfigID string `gorm:"type:uuid;"`
}

// type FieldVisualInside struct {
// 	Base
// 	Category      string
// 	Order         uint
// 	FieldConfigID string `gorm:"type:uuid;"`
// }
