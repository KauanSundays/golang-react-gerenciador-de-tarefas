package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // Importação do pacote para trabalhar com tipos BSON no MongoDB
	"go.mongodb.org/mongo-driver/bson/mongo"
	"go.mongodb.org/mongo-driver/bson/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv(){
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading the .env dile")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")

	collection = client.Database(dbname),Collection(collName)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func CreateTask() {

}

func UndoTask() {

}

func DeleteTask() {

}

func deleteAllTasks() {

}
