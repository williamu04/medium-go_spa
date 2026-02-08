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
	article_application "github.com/williamu04/medium-clone/application/articles"
	bookmark_application "github.com/williamu04/medium-clone/application/bookmarks"
	comment_application "github.com/williamu04/medium-clone/application/comments"
	follow_application "github.com/williamu04/medium-clone/application/follows"
	topic_application "github.com/williamu04/medium-clone/application/topic"
	user_application "github.com/williamu04/medium-clone/application/users"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/infrastructure"
	"github.com/williamu04/medium-clone/infrastructure/db_repository"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port"
)

func main() {
	config := pkg.LoadConfig()

	log := pkg.NewLogger(config.LogLevel)

	log.Infof("Starting app with log level: %s", config.LogLevel)

	db, err := infrastructure.InitDatabase(config.DBDSN, log)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	model.AutoMigrate(db)

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	} else {
		defer sqlDB.Close()
	}

	passwordHasher := pkg.NewHasher()
	tokenGenerator, err := pkg.NewJWTGen(config.JWTSecret, config.JWTExpiry)
	sluger := pkg.NewSluger()

	userRepository := db_repository.NewUserDatabaseRepository(db, log)
	articleRepository := db_repository.NewArticleDatabaseRepository(db, log)
	topicRepository := db_repository.NewTopicDatabaseRepository(db, log)
	commentRepository := db_repository.NewCommentDatabaseRepository(db, log)
	bookmarkRepository := db_repository.NewBookmarkDatabaseRepository(db, log)
	followRepository := db_repository.NewFollowDatabaseRepository(db, log)

	if err != nil {
		panic(err)
	}

	userUseCase := user_application.NewUserUseCase(userRepository, passwordHasher, tokenGenerator)
	articleUseCase := article_application.NewArticleUseCase(articleRepository, topicRepository, sluger)
	topicUseCase := topic_application.NewTopicUseCase(topicRepository, sluger)
	commentUseCase := comment_application.NewCommentUseCase(commentRepository)
	bookmarkUseCase := bookmark_application.NewBookmarkUseCase(bookmarkRepository)
	followUseCase := follow_application.NewFollowUseCase(followRepository)

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
