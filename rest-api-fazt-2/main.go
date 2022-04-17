package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	task{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
	task{
		ID:      2,
		Name:    "Task Two",
		Content: "Other content",
	},
}

// indexRoute is the default route
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my API")
}

// getTasks returns all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// getTaskById returns a task by its ID
func getTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, task := range tasks {
		if task.ID == taskID {

			json.NewEncoder(w).Encode((task))
			return
		}
	}
	json.NewEncoder(w).Encode(task{})

}

// createTask creates a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid Task")
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// updateTaskById updates a task by its ID
func updateTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	var updatedTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid tata")
	}
	json.Unmarshal(reqBody, &updatedTask)
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "Task with ID %v has been updated", taskID)
		}
	}
}

// deleteTaskById deletes a task by its ID
func deleteTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "Task with ID %v has been deleted", taskID)
			return
		}

	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTaskById).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTaskById).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTaskById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
