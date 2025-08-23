package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// test CI/CD workflow 2
// test CI/CD workflow 3
// Struct untuk user
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func main() {
	// Data dummy manual (nggak pake DB)
	users := []User{
		{ID: 2, Name: "Amar", Email: "amar@example.com", CreatedAt: time.Now().Format(time.RFC3339)},
		{ID: 3, Name: "Maulana", Email: "maulana@example.com", CreatedAt: time.Now().Format(time.RFC3339)},
	}

	// Route root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format(time.RFC3339)
		fmt.Fprintf(w, "cangcut amar.maulana@domain.com - Time: %s", now)
	})

	// Route GET /users → ambil data manual
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
