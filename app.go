// app.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// App - Application stuct
type App struct {
	Router *mux.Router
}

// Initialise - Initialise the application
func (a *App) Initialise() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.initialiseRoutes()
}

// Run - Run the application
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initialiseRoutes() {
	a.Router.HandleFunc("/api/skills", a.getSkills).Methods("GET")
	a.Router.HandleFunc("/api/skills/{id:[0-9]+}", a.getSkill).Methods("GET")
}

func (a *App) getSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := getSkills()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, skills)
}

func (a *App) getSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid skill ID")
		return
	}

	skill, err := getSkill(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, skill)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
