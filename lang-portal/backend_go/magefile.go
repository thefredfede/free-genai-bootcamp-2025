// filepath: /Users/securingus/Downloads/BootCamp/free-genai-bootcamp-2025/lang-portal/backend_go/magefile.go
//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Build compiles the application
func Build() error {
	fmt.Println("Building the application...")
	cmd := exec.Command("go", "build", "-o", "lang-portal-backend", "./cmd/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Run starts the application
func Run() error {
	fmt.Println("Running the application...")
	cmd := exec.Command("./lang-portal-backend")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Migrate runs the database migrations
func Migrate() error {
	fmt.Println("Running database migrations...")
	dbPath := filepath.Join("internal", "db", "migrations", "001_create_tables.sql")
	cmd := exec.Command("sqlite3", "words.db", fmt.Sprintf(".read %s", dbPath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Seed seeds the database with initial data
func Seed() error {
	fmt.Println("Seeding the database...")
	seedPath := filepath.Join("internal", "db", "seeds", "seed_data.sql")
	cmd := exec.Command("sqlite3", "words.db", fmt.Sprintf(".read %s", seedPath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
