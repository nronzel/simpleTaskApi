package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// @Summary Create a new task
// @Description Add a new task to the list
// @Accept json
// @Produce json
// @Param task body parameters true "Task to add"
// @Success 201 {object} taskEntry
// @Failure 500 {object} string "invalid json payload"
// @Failure 400 {object} string "task content cannot be empty"
// @Router /tasks [post]
func (db *tempDB) handlerTasksPost(w http.ResponseWriter, r *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	id := db.nextID

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		db.logError(r, http.StatusInternalServerError, "invalid json payload", err)
		respondWithError(w, http.StatusInternalServerError, "invalid json payload")
		return
	}

	if params.Task == "" {
		db.logError(r, http.StatusBadRequest, "task content cannot be empty", nil)
		respondWithError(w, http.StatusBadRequest, "task content cannot be empty")
		return
	}

	newTask := taskEntry{
		Id:        id,
		Task:      params.Task,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	// store the task in the db
	db.tasks[id] = newTask
	db.nextID++

	respondWithJSON(w, http.StatusCreated, newTask)
}
