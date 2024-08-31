package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView Handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the Query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't

	w.Write([]byte("Display a specific snippet"))
}

// Add a snippetCreate Handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check wheather the request is using POST or not.
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// Use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servmux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server, we pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
