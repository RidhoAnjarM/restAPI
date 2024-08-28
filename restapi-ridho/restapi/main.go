package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    initDB()

    r := mux.NewRouter()
    r.Use(TokenValidationMiddleware)

    // Definisikan route untuk setiap endpoint
    r.HandleFunc("/run/user", createUser).Methods("POST")
    r.HandleFunc("/list/users", listUsers).Methods("GET")
    r.HandleFunc("/detail/user/{id}", getUser).Methods("GET")
    r.HandleFunc("/run/user/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/run/user/{id}", deleteUser).Methods("DELETE")

    http.ListenAndServe(":8080", r)
}

