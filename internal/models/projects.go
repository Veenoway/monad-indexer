package models

import "time"

type Project struct {
	ID          string    `json:"id"`
	CreatorID   string    `json:"creator_id"`  
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Categories  []string  `json:"categories"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	DevID	    string    `json:"dev_id"`
	MissionID   string    `json:"mission_id"`
}