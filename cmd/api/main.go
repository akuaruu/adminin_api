package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akuaruu/adminin_api/internal/handler"
	"github.com/akuaruu/adminin_api/internal/service"
	"github.com/akuaruu/adminin_api/internal/repository"
	"github.com/akuaruu/adminin_api/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//DATABASE INITIALIZATION
	db, err := gorm.Open(sqlite.Open("adminin.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	db.AutoMigrate(&model.Admin{})

	//REPOSITORY INITIALIZATION
	adminRepo := repository.NewAdminRepository(db)

	//service initialization
	adminService := service.NewAdminService(adminRepo)
	
	//handler initialization
	adminHandler := handler.NewAdminHandler(adminService)


	//HTTP SERVER SETUP
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", adminHandler.Register)

	port := "8080"
	fmt.Println("Server is running on http://localhost:" + port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	
}