// ---- Go Module & Build Commands with Meaning ----
// go mod init name
// Initializes a new Go module in the current directory with the given name (creates a go.mod file).
// go mod tidy
// Adds missing and removes unused modules (cleans up dependencies).
// go list -m all
// Lists all modules your code depends on (direct and indirect).
// go mod verify
// Verifies dependencies against their checksums to ensure integrity and security.
// go build
// Compiles the Go code in the current directory into a binary.
// go mod graph
// Shows a dependency graph of all modules your project uses.
// go mod edit -go 1.16
// Edits the Go version used in your module file (e.g., set Go version to 1.16).
// go list -m -versions github.com/gorilla/mux
// Lists all versions available for the `gorilla/mux` package.
// go mod vendor
// Copies all module dependencies to a `vendor/` directory (used for reproducible builds).
// ---- Actual Go Code Below ----
package main // Declares the main package; required for an executable Go program.

import (
	"fmt"      // For printing to the console
	"log"      // For logging errors
	"net/http" // For handling HTTP requests and responses

	"github.com/gorilla/mux" // Third-party router for handling routes
)

func main() {
	fmt.Println("Hello mod in go lang") // Prints a greeting message
	greeter()                           // Calls the greeter function

	// Initializes a new router using gorilla/mux
	r := mux.NewRouter()

	// Registers a handler function for the root route "/" using HTTP GET method
	r.HandleFunc("/", serveHome).Methods("GET")

	// Starts the HTTP server on port 4000. If it fails, logs the error and stops.
	log.Fatal(http.ListenAndServe(":4000", r))
}

// greeter prints a simple greeting to the console
func greeter() {
	fmt.Println("hey there mod users")
}

// serveHome writes a basic HTML response to the client
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang</h1>"))
}
