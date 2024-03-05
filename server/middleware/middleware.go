package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv" // Importação do pacote para trabalhar com tipos BSON no MongoDB
	"go.mongodb.org/mongo-driver/bson/mongo"
	"go.mongodb.org/mongo-driver/bson/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()       // Inicializa o carregamento das variáveis de ambiente
	createDBInstance() // Inicializa a instância do banco de dados (ctrl+click)
}

func loadTheEnv() { // Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env dile") // Encerra o programa se ocorrer um erro ao carregar o arquivo .env
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")     // Obtém a URI do banco de dados do ambiente
	dbName := os.Getenv("DB_NAME")              // Obtém o nome do banco de dados do ambiente
	collName := os.Getenv("DB_COLLECTION_NAME") // Obtém o nome da collection

	clientOptions := options.Client().ApplyURI(connectionString) // Define as opções do cliente com a URI do banco de dados
	client, err := mongo.Connect(context.TODO(), clientOptions)  // Conecta ao banco de dados MongoDB

	if err != nil {
		log.Fatal(err) // Encerra o programa se ocorrer um erro na conexão
	}

	fmt.Println("connected to mongodb") // deu certo

	collection = client.Database(dbname), Collection(collName) // Define a coleção global para interagir com o banco de dados
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()           // Obtém todas as tarefas do banco de dados
	json.NewEncoder(w).Encode(payload) // Codifica as tarefas em JSON e as envia como resposta
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.Tarefas               // Declara uma variável para armazenar a nova tarefa
	json.NewEncoder(r.Body).Decode(&task) // Decodifica o corpo da solicitação JSON para obter os dados da tarefa
	insertOneTask(task)                   // Insere a nova tarefa no banco de dados
	json.NewEncoder(w).Encode(task)       // Codifica a tarefa em JSON e a envia como resposta
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	TaskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	UndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneTask(params["id"])
}

func DeleteAllTasks() {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTasks()
	json.NewEncoder(w).Encode(params["id"])
}

func getAllTasks() []primitive.M{
	cur, err := collection.Find(context.Background())

	if err!=nil {
		log.Fatal(err)
	}

	var results []primitive.Methods
	for cur.Next(context.Background()) {
		var result bson.M
		cur.Decode(&result)
		if e !=nil{
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results[:]
}
