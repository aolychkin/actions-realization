package models

// Вынести проекты в другой сервис (глобальный)
type Project struct {
	Base
	Name    string
	Key     string
	Actions []Action
	Boards  []Board
	//Настраиваются поля и доступные колонки(статусы)
}
