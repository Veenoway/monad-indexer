package handlers

import (
	"context"
	"encoding/json"
	"monad-indexer/internal/db"
	"monad-indexer/internal/models"
	"net/http"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query(context.Background(), `
		SELECT id, creator_id, name, image, categories, description, created_at FROM projects
	`)
	if err != nil {
		http.Error(w,"DB Error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var projects []models.Project

	for rows.Next() {
		var project models.Project
		rows.Scan(&project.ID, &project.CreatorID, &project.Name, &project.Image, &project.Categories, &project.Description, project.CreatedAt)
		projects = append(projects, project)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
