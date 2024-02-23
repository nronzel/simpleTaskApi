package main

import (
	"net/http"
	"sort"
	"strconv"
)

// @Summary List all tasks
// @Description Get a list of all tasks sorted by ID
// @Produce json
// @Success 200 {array} taskEntry
// @Router /tasks [get]
func (db *tempDB) handlerTasksGetAll(w http.ResponseWriter, r *http.Request) {
	var tasks []taskEntry

	db.mu.RLock()
	for _, task := range db.tasks {
		tasks = append(tasks, task)
	}
	db.mu.RUnlock()

	// sort by id
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	respondWithJSON(w, http.StatusOK, tasks)
}

// @Summary Get a task by ID
// @Description Get details of a specific task by ID
// @Produce json
// @Param taskID path int true "Task ID"
// @Success 200 {object} taskEntry
// @Failure 400 {object} string "taskID parameter is missing or invalid"
// @Failure 404 {object} string "task does not exist"
// @Router /tasks/{taskID} [get]
func (db *tempDB) handlerTasksGet(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("taskID")
	if strId == "" {
		db.logError(r, http.StatusBadRequest, "taskID parameter is missing", nil)
		respondWithError(w, http.StatusBadRequest, "taskID parameter is missing")
		return
	}

	// convert parameter to int
	id, err := strconv.Atoi(strId)
	if err != nil {
		db.logError(r, http.StatusBadRequest, "invalid id", err)
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	db.mu.RLock()
	// check if task exists in db
	task, exists := db.tasks[id]
	if !exists {
		db.logError(r, http.StatusNotFound, "task does not exist", nil)
		respondWithError(w, http.StatusNotFound, "task does not exist")
		return
	}
	db.mu.RUnlock()

	respondWithJSON(w, http.StatusOK, task)
}
