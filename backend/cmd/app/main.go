package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gridfs-file-manager/internal/config"
	AuthController "gridfs-file-manager/internal/controller/auth"
	FileController "gridfs-file-manager/internal/controller/file"
	UserController "gridfs-file-manager/internal/controller/user"
	AuthRepo "gridfs-file-manager/internal/repository/auth"
	FileRepo "gridfs-file-manager/internal/repository/file"
	UserRepo "gridfs-file-manager/internal/repository/user"
	AuthService "gridfs-file-manager/internal/services/auth"
	FileService "gridfs-file-manager/internal/services/file"
	UserService "gridfs-file-manager/internal/services/user"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	if err := godotenv.Load("app.env"); err != nil {
		log.Print("No app.env file found")
	}
}

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Print("can`t load config")
	}

	clientOptions := options.Client().ApplyURI(cfg.MongoURL)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("successfully connected to mongo")
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := mongoClient.Database(cfg.DbName)

	userCollection := db.Collection(cfg.UserCollection)
	fileCollection := db.Collection(cfg.GridFSCollection + ".files")

	gridFsOpts := options.GridFSBucket().SetName(cfg.GridFSCollection)
	bucket, err := gridfs.NewBucket(db, gridFsOpts)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := UserRepo.NewUserRepository(userCollection)
	authRepo := AuthRepo.NewUserRepository(userCollection)
	fileRepo := FileRepo.NewFileRepository(fileCollection, userCollection, bucket)

	userService := UserService.NewService(UserService.Deps{UserRepo: userRepo})
	authService := AuthService.NewService(AuthService.Deps{AuthRepo: authRepo})
	fileService := FileService.NewService(FileService.Deps{FileRepo: fileRepo, UserRepo: userRepo})

	authController := AuthController.NewController(AuthController.AuthService{AuthManagement: authService})
	userController := UserController.NewController(UserController.UserService{UserManagement: userService})
	fileController := FileController.NewController(FileController.FileService{FileManagement: fileService})

	app := fiber.New(fiber.Config{
		BodyLimit: 32 * 1024 * 1024,
	})
	file, err := os.OpenFile("./app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} ${method} ${path} ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     file,
	}))

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length," +
			" Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin, Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	authController.SetupRoute(app)
	userController.SetupRoute(app)
	fileController.SetupRoute(app)

	go func() {
		if err := app.Listen(":4000"); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Fiber was successful shutdown.")
}
