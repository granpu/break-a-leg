package handlers

import (
	"encoding/json"
	"net/http"
	"restapi-tasks/internal/models"
	"restapi-tasks/internal/models/handlers/database"
	"strconv"
	"strings"
	"vendor/golang.org/x/net/idna"
)

type Handlers struct {
	store *database.TaskStore
}

func NewHandlers(store *database.TaskStore) *Handlers {
	return &Handlers{
		store: store,
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondWitherror(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, error := h.store.GetAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка получение задач")
		return
	}
}
func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	idStr := pathParts[0]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неккорекктный ID задач")
		return
	}
	task, err := h.store.GetByID(id)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, task)
}
func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input models.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWitherror(w, http.StatusBadRequest, "Некорекктно отправленные данные")
		return
	}
	if strings.TrimSpace(input.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "Заголов задачи должен пристуствовать")
		return
	}
	task, err := h.store.Create(input)

	if err != nil {
		respondWrithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, task)
}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/"),
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неккорекктный ID задач")
		return
	}
	task, err := h.store.GetByID(id)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	var input models.UpdateTaskInput
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "Некорректные данные")
		return
	}

	if input.Title != nil && strings.TrimSpace(*input.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "Заголовок обязателен")
		return
	}

	task, err := h.stare.Update(id, input)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
		respondWithError(w, http.StatusNotFound, err.Error())
	} else {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	return
}
respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/"),
	idStr := pathParts[0]
	id, err := strconv.Atoi(idStr)
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "Неккорекктный ID задач")
		return
	}

	err = h.store.Delete(id)
	if err != nil{
		if strings.Contains(err.Error(), "record not found") {
		respondWithError(w, http.StatusNotFound, err.Error())
	} else {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "succes"})
}