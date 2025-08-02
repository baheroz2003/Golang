package main // Go ka main package jahan se program start hota hai

import (
	"fmt"      // Print karne ke liye
	"io"       // Input/Output handle karne ke liye (like body read)
	"net/http" // HTTP server banane ke liye
)

func main() {
	// Route set kar rahe hain alag-alag endpoint ke liye
	http.HandleFunc("/get", handleGetRequest)        // Jab GET request aaye is path pe
	http.HandleFunc("/post", handlePostRequest)      // Jab POST request aaye is path pe (JSON ya raw body)
	http.HandleFunc("/postform", handlePostFormRequest) // Jab POST form data aaye is path pe

	// Server start kar rahe hain port 8000 pe
	fmt.Println("🚀 Server running at http://localhost:8000")
	http.ListenAndServe(":8000", nil) // Yeh function server run karta hai
}

// Yeh function GET requests ko handle karta hai
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Agar method GET nahi hai toh error return karo
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	// Console me print karo ki GET request aayi
	fmt.Println("✅ Received GET request")
	w.WriteHeader(http.StatusOK) // 200 OK status bhejo
	w.Write([]byte("✔️ This is a response to your GET request")) // Client ko response bhejo
}

// Yeh function POST request handle karta hai jisme raw body hoti hai (JSON, text, etc.)
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Agar POST nahi hai toh error bhejo
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Request body ko read karo
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Agar body read nahi ho payi toh error bhejo
		http.Error(w, "❌ Could not read body", http.StatusBadRequest)
		return
	}

	// Console me print karo jo body client ne bheji hai
	fmt.Println("✅ Received POST body:", string(body))
	w.WriteHeader(http.StatusOK) // 200 OK status bhejo
	w.Write([]byte("✔️ Post received successfully")) // Confirmation response bhejo
}

// Yeh function POST form data handle karta hai (like HTML form)
func handlePostFormRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Agar POST nahi hai toh error bhejo
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Form data ko parse karo (form fields ko extract karo)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "❌ Error parsing form", http.StatusBadRequest)
		return
	}

	// Console me har ek field ka key/value print karo
	fmt.Println("✅ Received POST Form data:")
	for key, values := range r.PostForm {
		for _, value := range values {
			fmt.Printf("📌 %s = %s\n", key, value) // Har key ke saath uski value print karo
		}
	}

	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("✔️ POST Form processed")) // Client ko response bhejo
}
