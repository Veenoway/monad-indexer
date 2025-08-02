package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"monad-indexer/internal/db"
	"monad-indexer/internal/models"
	"net/http"
	"time"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	creatorID := r.URL.Query().Get("creator_id")
	category := r.URL.Query().Get("categories")
	name := r.URL.Query().Get("name")
	

	query := `SELECT id, creator_id, name, image, categories, description, created_at FROM projects WHERE 1=1`
	args := []interface{}{}
	i := 1

	if creatorID != "" {
		query += ` AND creator_id = $` + fmt.Sprint(i)
		args = append(args, creatorID)
		i++
	}

	if category != "" {
		query += ` AND $` + fmt.Sprint(i) + `= ANY(categories)`
		args = append(args, category)
		i++
	}

	if name != "" {
		query += `AND name = $` + fmt.Sprint(i) 
		args = append(args, name)
		i++
	}

	rows, err := db.Conn.Query(context.Background(), query, args...)
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

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project 
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w,"Error while decoding body", http.StatusBadRequest)
		return
	}
	
	project.CreatedAt = time.Now()
	_, err := db.Conn.Exec(context.Background(),
		`INSERT INTO projects (name, creator_id, image, categories, description, created_at) VALUES ($1 $2 $3 $4 $5 $6)`,
		&project.Name, &project.Image, &project.Categories, &project.Description, &project.CreatedAt,
	)

	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message":"Successfully created a project"})
}

