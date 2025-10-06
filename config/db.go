package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://muazahmad_db_user:0HMrWqakO4FMsZCr@employee.w6hjpis.mongodb.net/library_db?retryWrites=true&w=majority&appName=Library+Management+System"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("library_db")
	fmt.Println("✅ Connected to MongoDB!")
}


