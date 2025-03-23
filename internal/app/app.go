package app

import (
	"actions/internal/storage/sqlite"

	"log/slog"

	grpc_app "actions/internal/app/grpc"
	"actions/internal/services/board"
)

type App struct {
	GRPCSrv *grpc_app.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	boardService := board.New(log, storage)

	grpcApp := grpc_app.New(log, boardService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
