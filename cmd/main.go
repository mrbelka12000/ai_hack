package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/ai_hack/internal/client/ml"
	v1 "github.com/mrbelka12000/ai_hack/internal/delivery/http/v1"
	"github.com/mrbelka12000/ai_hack/internal/repo"
	"github.com/mrbelka12000/ai_hack/internal/usecase"
	"github.com/mrbelka12000/ai_hack/migrations"
	"github.com/mrbelka12000/ai_hack/pkg/config"
	"github.com/mrbelka12000/ai_hack/pkg/gorm/postgres"
	"github.com/mrbelka12000/ai_hack/pkg/redis"
	"github.com/mrbelka12000/ai_hack/pkg/server"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      aiport.mrbelka12000.com
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service_name", cfg.ServiceName)

	db, err := postgres.New(cfg.PGURL)
	if err != nil {
		log.With("error", err).Error("failed to connect to database")
		return
	}
	migrations.RunMigrations(db)

	repository := repo.New(db)
	rds, err := redis.New(cfg)
	if err != nil {
		log.With("error", err).Error("failed to connect to redis")
		return
	}

	mlClient := ml.NewClient(cfg.AISuflerAPIURL, log)

	uc := usecase.New(repository, log, rds, mlClient)

	if cfg.RunMBMigration {
		if err := uc.StartParseMB(cfg.CSVFile); err != nil {
			log.Error("failed to start parser", "error", err)
			return
		}
	}

	mx := mux.NewRouter()
	v1.Init(uc, mx, log)

	srv := server.New(mx, cfg.HTTPPort)

	srv.Start()
	log.With("port", cfg.HTTPPort).Info("Starting HTTP server")

	gs := make(chan os.Signal, 1)
	signal.Notify(gs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-gs:
		log.Info(fmt.Sprintf("Received signal: %d", sig))
		log.Info("Server stopped properly")
		srv.Stop()
		close(gs)
	case err := <-srv.Ch():
		log.With("error", err).Error("Server stopped")
	}
}
