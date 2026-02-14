package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	articleapplication "github.com/williamu04/medium-clone/application/articles"
	bookmarkapplication "github.com/williamu04/medium-clone/application/bookmarks"
	commentapplication "github.com/williamu04/medium-clone/application/comments"
	followapplication "github.com/williamu04/medium-clone/application/follows"
	topicapplication "github.com/williamu04/medium-clone/application/topic"
	userapplication "github.com/williamu04/medium-clone/application/users"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/infrastructure"
	"github.com/williamu04/medium-clone/infrastructure/dbrepository"
	"github.com/williamu04/medium-clone/infrastructure/seeder"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port"
)

func main() {
	config := pkg.LoadConfig()

	log := pkg.NewLogger(config.LogLevel)

	log.Infof("Starting app with log level: B%s", config.LogLevel)

	db, err := infrastructure.InitDatabase(config.DBDSN, log)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Info("Drop table attempt")
	model.DropAllQuery(db)
	log.Info("Drop all table successfully, migrating new tables...")
	model.AutoMigrate(db)
	log.Info("Migrate all table successfully")

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	} else {
		defer sqlDB.Close()
	}

	passwordHasher := pkg.NewHasher()
	tokenGenerator, err := pkg.NewJWTGen(config.JWTSecret, config.JWTExpiry)
	sluger := pkg.NewSluger()

	if config.SeedData == "true" {
		dataSeedHandler := seeder.NewSeedHandler(db, log, passwordHasher, sluger)
		if err := dataSeedHandler.SeedAll(); err != nil {
			log.Errorf("Error seeding data: %v", err)
		}
		log.Info("Database seeded successfully")
	}

	userRepository := dbrepository.NewUserDatabaseRepository(db, log)
	articleRepository := dbrepository.NewArticleDatabaseRepository(db, log)
	topicRepository := dbrepository.NewTopicDatabaseRepository(db, log)
	commentRepository := dbrepository.NewCommentDatabaseRepository(db, log)
	bookmarkRepository := dbrepository.NewBookmarkDatabaseRepository(db, log)
	followRepository := dbrepository.NewFollowDatabaseRepository(db, log)

	if err != nil {
		panic(err)
	}

	userUseCase := userapplication.NewUserUseCase(userRepository, passwordHasher, tokenGenerator)
	articleUseCase := articleapplication.NewArticleUseCase(articleRepository, topicRepository, sluger)
	topicUseCase := topicapplication.NewTopicUseCase(topicRepository, sluger)
	commentUseCase := commentapplication.NewCommentUseCase(commentRepository)
	bookmarkUseCase := bookmarkapplication.NewBookmarkUseCase(bookmarkRepository)
	followUseCase := followapplication.NewFollowUseCase(followRepository)

	authMiddleware := middleware.NewAuthMiddleware(log, tokenGenerator)

	restAPIRouter := gin.Default()
	restAPIRouter.RedirectTrailingSlash = false
	restAPIRouter.Use(cors.Default())

	v1 := restAPIRouter.Group("/v1/api")

	restAPIHandler := port.NewRestAPIHandler(log, userUseCase, articleUseCase, topicUseCase, commentUseCase, bookmarkUseCase, followUseCase, authMiddleware)
	restAPIHandler.RegisterRoutes(v1)

	srv := &http.Server{
		Addr:    ":" + config.RestAPIPort,
		Handler: restAPIRouter,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start REST server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Info("Graceful shutdown...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Errorf("REST server forced to shutdown: %v", err)
	}

	log.Info("Servers exited properly")
}
