package db

import (
	"context"
	"fmt"
	"log"
)

func Migrate() {
	schema := `
	CREATE TABLE IF NOT EXISTS devs (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL,
		profile_image TEXT,
		roles TEXT[],
		address TEXT NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS projects (
		id TEXT PRIMARY KEY,
		creator_id TEXT REFERENCES devs(id),
		name TEXT NOT NULL,
		image TEXT,
		categories TEXT[],
		description TEXT,
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

	CREATE TABLE IF NOT EXISTS mission_projects (
		id TEXT PRIMARY KEY DEFAULT gen_random_uuid(),
		mission_id TEXT REFERENCES missions(id) ON DELETE CASCADE,
		project_id TEXT REFERENCES projects(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS mission_winners (
		id TEXT PRIMARY KEY DEFAULT gen_random_uuid(),
		mission_id TEXT REFERENCES missions(id) ON DELETE CASCADE,
		dev_id TEXT REFERENCES devs(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS applications (
		id TEXT PRIMARY KEY DEFAULT gen_random_uuid(),
		mission_id TEXT REFERENCES missions(id) ON DELETE CASCADE,
		dev_id TEXT REFERENCES devs(id) ON DELETE CASCADE,
		message TEXT,
		status TEXT CHECK (status IN ('pending', 'accepted', 'rejected')) DEFAULT 'pending',
		submitted_at TIMESTAMPTZ DEFAULT NOW()
	);
	`

	_, err := Conn.Exec(context.Background(), schema)
	if err != nil {
		log.Fatal("❌ Erreur création des tables :", err)
	}

	fmt.Println("✅ Tables vérifiées / créées")
}