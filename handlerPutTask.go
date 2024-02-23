package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// @Summary Update a task
// @Description Update an existing task by ID
// @Accept json
// @Produce json
// @Param taskID path int true "Task ID"
// @Param task body string true "Task information"
// @Success 200 {object} taskEntry
// @Failure 400 {object} string "taskID parameter not supplied or invalid"
// @Failure 404 {object} string "task not found"
// @Failure 500 {object} string "invalid json payload"
// @Router /tasks/{taskID} [put]
func (db *tempDB) handlerTasksUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("taskID")
	if idStr == "" {
		db.logError(r, http.StatusBadRequest, "taskID parameter not supplied", nil)
		respondWithError(w, http.StatusBadRequest, "taskID parameter not supplied")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		db.logError(r, http.StatusBadRequest, "invalid id", err)
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	// ensure task exists
	_, ok := db.tasks[id]
	if !ok {
		db.logError(r, http.StatusNotFound, "task not found", nil)
		respondWithError(w, http.StatusNotFound, "task not found")
		return
	}

	params := parameters{}

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		db.logError(r, http.StatusInternalServerError, "invalid json payload", err)
		respondWithError(w, http.StatusInternalServerError, "invalid json payload")
		return
	}

	// update task
	db.tasks[id] = taskEntry{
		Task:      params.Task,
		UpdatedAt: time.Now().UTC(),
	}

	respondWithJSON(w, http.StatusOK, db.tasks[id])
}
