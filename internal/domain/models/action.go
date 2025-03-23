package models

//(!!) В центре всего Действие, которое добавляется на проекты, а уже проекты на доски.
//(!!) Доска - это просто сущность отображения (некий дашборд)

// Может находиться только на одной колонке. Если нужно отобразить задачи этой колонки, то ее нужно добавить на доску
// TODO: добавить воркфлоу
// Одна карточка может находиться только в 1 статуе-колонке. Этот статус-колока может быть уже добавлена на колонку на борде

type Action struct {
	Base
	ActionNum     uint
	Type          string
	ProjectID     string    `gorm:"type:uuid;"`
	CurrentStepID string    `gorm:"type:uuid;"` //`gorm:"many2many:action_steps;"`
	Sprints       []*Sprint `gorm:"many2many:action_sprints;"`
	Fields        []ActionField
}

// Value - строка, которую нужно будет парсить на фронте
type ActionField struct {
	Base
	Value         string
	ActionID      string      `gorm:"type:uuid;"`
	FieldConfigID string      `gorm:"type:uuid;"`
	Activities    []*Activity `gorm:"many2many:action_field_activities;"`
}

// TODO: Если карточка на нескольких бордах/проекта, то нужно настроить ValueSource (какие данные подтягиваются между полями для отображения)
type FieldConfig struct {
	Base
	Name            string
	Alias           string
	FieldTypeID     string `gorm:"type:uuid;"`
	FieldType       FieldType
	DefaultValue    string
	AvailableValues string
	IsPrimary       bool
	ActionFields    []ActionField
}
type FieldType struct {
	Base
	Name           string
	Alias          string
	IsCustom       bool
	AvailableSizes string
}
type Activity struct {
	Base
	Type         string
	ActionFields []*ActionField `gorm:"many2many:action_field_activities;"`
}
