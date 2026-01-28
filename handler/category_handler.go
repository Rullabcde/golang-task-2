package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"category-api/entity"
	"category-api/service"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/categories", h.GetAll).Methods("GET")
	router.HandleFunc("/categories", h.Create).Methods("POST")
	router.HandleFunc("/categories/{id}", h.GetByID).Methods("GET")
	router.HandleFunc("/categories/{id}", h.Update).Methods("PUT")
	router.HandleFunc("/categories/{id}", h.Delete).Methods("DELETE")
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	writeSuccess(w, http.StatusOK, categories)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	category, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			writeError(w, http.StatusNotFound, "Category not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to fetch category")
		return
	}

	writeSuccess(w, http.StatusOK, category)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req entity.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	category, err := h.service.Create(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create category")
		return
	}

	writeSuccessWithMessage(w, http.StatusCreated, "Category created successfully", category)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req entity.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	category, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			writeError(w, http.StatusNotFound, "Category not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to update category")
		return
	}

	writeSuccessWithMessage(w, http.StatusOK, "Category updated successfully", category)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			writeError(w, http.StatusNotFound, "Category not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	writeSuccessMessage(w, http.StatusOK, "Category deleted successfully")
}

// Helper functions
func getIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}

func writeSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func writeSuccessWithMessage(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func writeSuccessMessage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": message,
	})
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": message,
	})
}
