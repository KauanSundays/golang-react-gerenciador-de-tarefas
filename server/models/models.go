package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tarefas struct {
	ID		primitive.ObjectID		`json:"_id, omitempty" bson: "_id,omitempty"`
	Task	string					`json:"tarefa,omitempty"`
	Status 	bool					`json:"status, omitempty"`
}