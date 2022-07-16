package config

import (
	"context"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySql *gorm.DB
var Mongo *mongo.Client

func Connect() {
	mySqlConn := "geoquest:geoquest@tcp(localhost:3306)/geoquest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mySqlConn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{})
	MySql = db

}

func GetCollection(collection string) *mongo.Collection {
	mongoConn := "mongodb://geoquest:geoquest@localhost:27017/?authSource=admin&readPreference=primary&ssl=false"
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConn))

	//mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		panic(err)
	}
	Mongo = mongoClient

	return Mongo.Database("geoquest").Collection(collection)
}
