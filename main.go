package main

import (
	"go-mongo-redis/configuration"
	ds "go-mongo-redis/domain/datasources"
	repo "go-mongo-redis/domain/repositories"
	gw "go-mongo-redis/src/gateways"
	"go-mongo-redis/src/middlewares"
	sv "go-mongo-redis/src/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())

	mongodb := ds.NewMongoDB(10)
	redisdb := ds.NewRedisConnection()

	userMongo := repo.NewUsersRepository(mongodb)
	userRedis := repo.NewRedisRepository(redisdb)

	sv0 := sv.NewUsersService(userMongo, userRedis)

	gw.NewHTTPGateway(app, sv0)

	PORT := os.Getenv("DB_PORT_LOGIN")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
