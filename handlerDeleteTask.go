package main

import (
	"net/http"
	"strconv"
)

// @Summary Delete a task
// @Description Remove a task by ID
// @Produce json
// @Param taskID path int true "Task ID"
// @Success 200 {object} taskEntry
// @Failure 400 {object} string "taskID parameter not supplied or invalid"
// @Failure 404 {object} string "task not found"
// @Router /tasks/{taskID} [delete]
func (db *tempDB) handlerTasksDelete(w http.ResponseWriter, r *http.Request) {
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

	// ensure task exists
	deletedTask, exists := db.tasks[id]
	if !exists {
		db.logError(r, http.StatusNotFound, "task not found", nil)
		respondWithError(w, http.StatusNotFound, "task not found")
		return
	}

	delete(db.tasks, id)
	respondWithJSON(w, http.StatusOK, deletedTask)
}
