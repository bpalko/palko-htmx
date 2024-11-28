package main

import (
	"html/template"
	"net/http"

	"github.com/bpalko/palko-htmx/internal/db"

	"github.com/bpalko/palko-htmx/logger"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db.Initialize()

	r := mux.NewRouter()
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Routes
	r.HandleFunc("/", handleIndex(tmpl)).Methods("GET")
	r.HandleFunc("/create-build", handleCreateBuild(tmpl)).Methods("GET", "POST")

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	logger.GetLogger().Info("Server starting on http://localhost:8080")
	logger.GetLogger().Fatal(http.ListenAndServe(":8080", r))
}

// Handler for the root path
func handleIndex(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			logger.GetLogger().Errorf("%v", http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Handler for the create build page
func handleCreateBuild(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Retrieve form values directly from the request
			barrel := r.FormValue("barrel")
			grip := r.FormValue("grip")
			sight := r.FormValue("sight")

			// Create a new build
			build := db.Build{
				Barrel: barrel,
				Grip:   grip,
				Sight:  sight,
			}

			// Save the build to the database
			err := db.SaveBuild(build)
			if err != nil {
				http.Error(w, "Failed to save build", http.StatusInternalServerError)
				return
			}

			// Return a success message
			w.Write([]byte("Build created successfully!"))
			return
		}
		tmpl.ExecuteTemplate(w, "create-build.html", nil)
	}
}
