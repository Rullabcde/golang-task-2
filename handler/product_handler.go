package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"category-api/entity"
	"category-api/service"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.GetAll).Methods("GET")
	router.HandleFunc("/products", h.Create).Methods("POST")
	router.HandleFunc("/products/{id}", h.GetByID).Methods("GET")
	router.HandleFunc("/products/{id}", h.Update).Methods("PUT")
	router.HandleFunc("/products/{id}", h.Delete).Methods("DELETE")
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	writeSuccess(w, http.StatusOK, products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	product, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			writeError(w, http.StatusNotFound, "Product not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to fetch product")
		return
	}

	writeSuccess(w, http.StatusOK, product)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req entity.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	product, err := h.service.Create(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	writeSuccessWithMessage(w, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req entity.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	product, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			writeError(w, http.StatusNotFound, "Product not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	writeSuccessWithMessage(w, http.StatusOK, "Product updated successfully", product)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			writeError(w, http.StatusNotFound, "Product not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	writeSuccessMessage(w, http.StatusOK, "Product deleted successfully")
}
