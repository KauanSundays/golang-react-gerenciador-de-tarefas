package models // Declaração do pacote models

import "go.mongodb.org/mongo-driver/bson/primitive" // Importação do pacote para trabalhar com tipos BSON no MongoDB

// Definição da estrutura Tarefas para representar uma tarefa
type Tarefas struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` // Identificador único da tarefa no MongoDB
	Task   string             `json:"tarefa,omitempty"`                    // Descrição da tarefa
	Status bool               `json:"status,omitempty"`                    // Status da tarefa (completa ou incompleta)
}
