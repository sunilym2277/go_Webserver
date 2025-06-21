package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerRoot)
	mux.HandleFunc("POST /user", createUser)
	mux.HandleFunc("GET /user/{ID}", getUser)
	mux.HandleFunc("DELETE /user/{ID}", deleteUser)
	fmt.Printf("Server listening to 8088")
	http.ListenAndServe(":8088", mux)

}

type User struct {
	Name string `json:"name"`
}

var cacheMutex sync.RWMutex
var userCache = make(map[int]User)

func handlerRoot(
	w http.ResponseWriter, r *http.Request,
) {
	fmt.Fprintf(w, "Hello World!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
	}
	if user.Name == "" {
		http.Error(
			w,
			"Enter name",
			http.StatusBadRequest,
		)
	}
	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("ID"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest)
		return
	}
	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(
			w,
			"User Not found ",
			http.StatusNotFound)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
	w.Header().Set("Content-Type", "application/Json")
	w.Write(j)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("ID"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest)
		return
	}
	cacheMutex.RLock()
	_, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(
			w,
			"User Not found ",
			http.StatusNotFound)
		return
	}
	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
