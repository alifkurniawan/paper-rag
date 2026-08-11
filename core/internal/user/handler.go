package user

import (
	"core/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/users", h.List)
	r.Get("/users/{id}", h.GetByID)
	r.Post("/users", h.Create)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	u, err := h.service.GetByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, u)
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.service.Create(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, http.StatusOK, u)

}
