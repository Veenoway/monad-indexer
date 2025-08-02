package db

import (
	"context"
	"fmt"
	"log"
)

func Migrate() {
	// drop := `
	// DROP TABLE IF EXISTS project_devs CASCADE;
	// DROP TABLE IF EXISTS projects CASCADE;
	// DROP TABLE IF EXISTS missions CASCADE;
	// DROP TABLE IF EXISTS devs CASCADE;
	// `

	create := `
	CREATE TABLE IF NOT EXISTS devs (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL,
		roles TEXT[],
		profile_image TEXT,
		address TEXT UNIQUE NOT NULL,
		twitter TEXT,
		discord TEXT,
		created_at TIMESTAMPTZ DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS missions (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		start_time TIMESTAMPTZ,
		end_time TIMESTAMPTZ,
		round INT,
		created_at TIMESTAMPTZ DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS projects (
		id TEXT PRIMARY KEY,
		dev_id TEXT REFERENCES devs(id),
		mission_id TEXT REFERENCES missions(id),
		name TEXT NOT NULL,
		image TEXT,
		categories TEXT[],
		description TEXT,
		created_at TIMESTAMPTZ DEFAULT NOW()
	);
	
	CREATE TABLE IF NOT EXISTS project_devs (
		id TEXT PRIMARY KEY DEFAULT gen_random_uuid(),
		project_id TEXT REFERENCES projects(id) ON DELETE CASCADE,
		dev_id TEXT REFERENCES devs(id) ON DELETE CASCADE
	);
	`


	// _, err := Conn.Exec(context.Background(), drop)
	// if err != nil {
	// 	log.Fatal("❌ Error dropping tables:", err)
	// }

	_, err := Conn.Exec(context.Background(), create)
	if err != nil {
		log.Fatal("❌ Erreur création des tables :", err)
	}

	fmt.Println("✅ Schema dropped and recreated successfully")
}
