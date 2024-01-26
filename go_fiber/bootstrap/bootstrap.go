package bootstrap

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Logger             *log.Logger
	MongoConnection    *mongo.Database
	PostgresConnection *sql.DB
}

func NewInitializeBootstrap() Application {
	app := Application{}
	// app.MongoConnection = mongoConnection()
	app.PostgresConnection = postgresConnection()
	return app
}
