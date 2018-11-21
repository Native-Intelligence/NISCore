package core

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
	// "os"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/joho/godotenv"
)

// var db *gorm.DB
type baseModel struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type key string

const (
	hostKey     = key("hostKey")
	usernameKey = key("usernameKey")
	passwordKey = key("passwordKey")
	databaseKey = key("databaseKey")
)

// func init() {

// 	e := godotenv.Load()
// 	if e != nil {
// 		fmt.Print(e)
// 	}

// 	username := os.Getenv("db_user")
// 	password := os.Getenv("db_pass")
// 	dbName := os.Getenv("db_name")
// 	dbHost := os.Getenv("db_host")

// 	// dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
// 	dbUri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=%s", username, password, dbName, dbHost)
// 	fmt.Println(dbUri)

// 	conn, err := gorm.Open("mysql", dbUri)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	db = conn
// 	db.Debug().AutoMigrate(&Account{}, &Contact{})
// }

// func GetDB() *gorm.DB {
// 	return db
// }

func db(ctx context.Context) (*mongo.Database, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value(usernameKey),
		ctx.Value(passwordKey),
		ctx.Value(hostKey),
		ctx.Value(databaseKey),
	)
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}
	todoDB := client.Database("todo")
	return todoDB, nil
}
