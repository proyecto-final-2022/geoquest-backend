package config

import (
	"context"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Mongo *mongo.Client

func Connect() {
	dsn := "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{})
	DB = db

}

func GetCollection(collection string) *mongo.Collection {
	mongoUri := "mongodb://geoquest:geoquest@localhost:27017/?authSource=admin&readPreference=primary&ssl=false"
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))

	//mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		panic(err)
	}
	Mongo = mongoClient

	return Mongo.Database("geoquest").Collection(collection)
}
