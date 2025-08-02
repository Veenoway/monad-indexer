package handlers

import (
	"context"
	"encoding/json"
	"monad-indexer/internal/db"
	"monad-indexer/internal/models"
	"net/http"
	"time"
)

func CreateDev(w http.ResponseWriter, r *http.Request) {
	var dev models.Dev
	if err := json.NewDecoder(r.Body).Decode(&dev); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	dev.CreatedAt = time.Now()

	_, err := db.Conn.Exec(context.Background(), `
		INSERT INTO devs (id, username, profile_image, roles, address, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`,dev.ID, dev.Username, dev.ProfileImage, dev.Roles, dev.Address, dev.CreatedAt)

	if err != nil {
		http.Error(w, "DB error while inserting", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message":"Successfully created a dev"})
}

func GetAllDevs(w http.ResponseWriter, r *http.Request){
	rows,err := db.Conn.Query(context.Background(), `
		SELECT id, username, profile_image, roles, address, created_at FROM devs
	`)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var devs []models.Dev 

	for rows.Next() {
		var dev models.Dev
		rows.Scan(&dev.ID,  &dev.Username, &dev.ProfileImage, &dev.Roles, &dev.Address, &dev.CreatedAt)
		devs = append(devs, dev)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devs)
}

// func GetDevProjects(w http.ResponseWriter, r *http.Request) {
// 	creatorID := r.URL.Query().Get("creator_id")
// 	if creatorID == "" {
// 		http.Error(w, "Missing creator_id param", http.StatusBadRequest)
// 		return
// 	}

// 	rows, err := db.Conn.Query(context.Background(), `
// 		SELECT id, creator_id, name, image, categories, description, created_at
// 		FROM projects
// 		WHERE creator_id = $1 
// 	`)
// } 

// func GetDev(w http.ResponseWriter, r*http.Request) {
// 	address := r.URL.Query().Get("address")
// 	if address == "" {
// 		http.Error(w, "Missing address param", http.StatusBadRequest)
// 		return
// 	}
// 	if rows, err := db.Conn.Query(context.Background(), `
// 		SELECT id, creator_id, name, image, categories, description, created_at
// 		FROM projects
// 		WHERE address = $1 
// 	`, &address); err != nil {
// 		http.Error(w,"No dev found with this address", http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

	
	
// }
