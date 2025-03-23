package main

import (
	"actions/cmd/actions/inits"
	"actions/internal/app"
	"actions/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	envLocal = "local"
	envStage = "stage"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()    //Загружаем конфиг
	log := setupLogger(cfg.Env) //Инициализация логгера
	initTestDB(cfg.StoragePath)

	//Запуск приложения
	log.Info("starting application", slog.Any("config", cfg))
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath)
	go application.GRPCSrv.MustRun()

	//Ожидаем сигнала завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	signal := <-stop

	//Мягко останавливаем работу приложения
	log.Info("stopping application", slog.String("signal", signal.String()))
	application.GRPCSrv.Stop()
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envStage:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func initTestDB(StoragePath string) {
	db, err := gorm.Open(sqlite.Open(StoragePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	inits.InitTestDB(db)
	inits.TestSeed(db)
}
