package router

import (
	"./middleware"
	"github.com/gorilla/mux" //mux ajuda a fazer o roteamento
)

func Router() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks/{id}", middleware.deleteAllTasks).Methods("DELETE", "OPTIONS")
	
	return router
}