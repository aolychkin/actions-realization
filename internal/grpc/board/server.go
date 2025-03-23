package board

import (
	"context"

	board_v1 "github.com/aolychkin/actions-contract/gen/go/board"
	"google.golang.org/grpc"
)

//TODO: зачем мне отображать содержимое всех полей и визуализации, когда я получаю именно борду?

// Структура, реализующая функционал API
type serverAPI struct {
	board_v1.UnimplementedBoardServiceServer
	board Board
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
// Каркасы для RPC-методов, которые мы будем использовать (описание обработчиков запросов)
type Board interface {
	GetBoard(
		ctx context.Context,
		id string,
	) (board *board_v1.TBoard, err error)
}

// Регистрируем serverAPI в gRPC-сервере
func Register(gRPC *grpc.Server, board Board) {
	board_v1.RegisterBoardServiceServer(gRPC, &serverAPI{board: board})
}
