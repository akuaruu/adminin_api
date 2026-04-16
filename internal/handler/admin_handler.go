package handler

import (
	"encoding/json"
	"net/http"

	"github.com/akuaruu/adminin_api/internal/model"
	"github.com/akuaruu/adminin_api/internal/service"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) Register(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	// Decode the JSON request body into the Admin struct
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	// Basic validation
	if err := h.adminService.Register(r.Context(), &admin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Admin registered successfully",
	})
}