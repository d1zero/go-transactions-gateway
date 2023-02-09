package app

import (
	"context"
	"fmt"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go-transactions-gateway/config"
	_ "go-transactions-gateway/docs"
	"go-transactions-gateway/internal/controller/http"
	"go-transactions-gateway/internal/domain/service"
	"go-transactions-gateway/internal/postgres"
	"go-transactions-gateway/pkg/govalidator"
	pgpkg "go-transactions-gateway/pkg/postgres"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	v := govalidator.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config loading error: %v", err)
	}

	err = v.Validate(context.Background(), cfg)
	if err != nil {
		log.Fatalf("config is not valid: %v", err)
	}

	atom := zap.NewAtomicLevel()
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		atom,
	)
	logger := zap.New(zapCore)
	defer logger.Sync()

	l := logger.Sugar()
	atom.SetLevel(zapcore.Level(*cfg.Logger.Level))
	l.Infof("logger initialized successfully")

	db, err := pgpkg.New(cfg.Postgres)
	if err != nil {
		l.Fatalf("database connection error: %s", err.Error())
	}
	defer func() {
		db.Close()
		l.Infof("database disconnected successfully")
	}()

	transactionRepository := postgres.NewTransactionRepository(db)

	transactionService := service.NewTransactionService(transactionRepository)

	transactionController := http.NewTransactionService(transactionService, *v)

	// @title Swagger Transaction Gateway API
	// @version 1.0
	// @description This is a transaction gateway server
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host 127.0.0.1:8080
	// @BasePath /api
	app := fiber.New()
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	apiGroup := app.Group("api")

	transactionGroup := apiGroup.Group("transaction")

	transactionController.RegisterTransactionRoutes(transactionGroup)

	go func() {
		l.Infof("starting server on port: %s:%d", cfg.Server.Host, cfg.Server.Port)
		if err := app.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)); err != nil {
			l.Fatalf("error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	err = app.Shutdown()
	if err != nil {
		l.Fatal(err)
	} else {
		l.Info("Fiber server exited properly")
	}
}
