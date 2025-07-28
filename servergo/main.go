package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Route handlers
	http.HandleFunc("/get", handleGetRequest)
	http.HandleFunc("/post", handlePostRequest)
	http.HandleFunc("/postform", handlePostFormRequest)

	// Start the server
	fmt.Println("🚀 Server running at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

// Handle GET requests
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("✅ Received GET request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✔️ This is a response to your GET request"))
}

// Handle raw POST (JSON or other body)
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "❌ Could not read body", http.StatusBadRequest)
		return
	}

	fmt.Println("✅ Received POST body:", string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✔️ Post received successfully"))
}

// Handle POST form data (application/x-www-form-urlencoded)
func handlePostFormRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "❌ Error parsing form", http.StatusBadRequest)
		return
	}

	fmt.Println("✅ Received POST Form data:")
	for key, values := range r.PostForm {
		for _, value := range values {
			fmt.Printf("📌 %s = %s\n", key, value)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✔️ POST Form processed"))
}
