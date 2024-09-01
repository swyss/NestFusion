package controllers

import (
	"encoding/json"
	"go-crud-api/internal/models"
	"go-crud-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskController struct {
	TaskService services.TaskServiceInterface
}

// ServeHTTP implementiert die http.Handler-Schnittstelle für TaskController.
func (c *TaskController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Festlegen des Content-Types auf JSON
	writer.Header().Set("Content-Type", "application/json")

	// Verwenden von Gorilla Mux zur Routenbestimmung
	router := mux.NewRouter()
	router.HandleFunc("/tasks", c.GetAllTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks", c.CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{id:[0-9]+}", c.UpdateTask).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{id:[0-9]+}", c.DeleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/tasks/{id:[0-9]+}/done", c.MarkTaskAsDone).Methods(http.MethodPost)

	// Mux-Router zur Bearbeitung der Anfrage verwenden
	router.ServeHTTP(writer, request)
}

func NewTaskController(service services.TaskServiceInterface) *TaskController {
	return &TaskController{TaskService: service}
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.TaskService.GetAllTasks()
	if err != nil {
		// If there is an error, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with the list of tasks and a 200 status code.
	c.handleJSONResponse(w, tasks, http.StatusOK)
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		// If decoding the request body fails, respond with a 400 status code.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}
	tasks, err := c.TaskService.CreateTask(&task)
	if err != nil {
		// If there is an error while creating the task, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with all the tasks and a 201 status code.
	c.handleJSONResponse(w, tasks, http.StatusCreated)

}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		// If decoding the request body fails, respond with a 400 status code.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}

	task.ID = uint(id)
	tasks, err := c.TaskService.UpdateTask(&task)
	if err != nil {
		// If there is an error while creating the task, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with all the tasks and a 201 status code.
	c.handleJSONResponse(w, tasks, http.StatusCreated)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}

	tasks, err := c.TaskService.DeleteTask(uint(id))
	if err != nil {
		// If there is an error while creating the task, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with all the tasks and a 201 status code.
	c.handleJSONResponse(w, tasks, http.StatusAccepted)
}
func (c *TaskController) MarkTaskAsDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Respond with a 400 status code if the ID is not a valid integer.
		c.handleError(w, err, http.StatusBadRequest)
		return
	}

	tasks, err := c.TaskService.MarkTaskAsDone(uint(id))

	if err != nil {
		c.handleError(w, err, http.StatusInternalServerError)
		return
	}

	c.handleJSONResponse(w, tasks, http.StatusOK)
}

func (c *TaskController) handleError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

func (c *TaskController) handleJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	jsonEncoder := json.NewEncoder(w)
	if err := jsonEncoder.Encode(data); err != nil {
		// If encoding the response fails, respond with a 500 status code.
		c.handleError(w, err, http.StatusInternalServerError)
	}
}
