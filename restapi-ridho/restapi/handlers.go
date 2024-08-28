package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Response struct {
    Status  bool        `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// Handler untuk membuat data User
func createUser(w http.ResponseWriter, r *http.Request) {
    var user Users
    json.NewDecoder(r.Body).Decode(&user)
    db.Create(&user)
    json.NewEncoder(w).Encode(Response{Status: true, Message: "User created", Data: user})
}

// Handler untuk menampilkan semua data User
func listUsers(w http.ResponseWriter, r *http.Request) {
    var users []Users
    db.Find(&users)
    json.NewEncoder(w).Encode(Response{Status: true, Message: "List of users", Data: users})
}

// Handler untuk menampilkan detail data User berdasarkan ID
func getUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user Users
    result := db.First(&user, params["id"])
    if result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(Response{Status: true, Message: "User found", Data: user})
}

// Handler untuk mengupdate data User berdasarkan ID
func updateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user Users
    db.First(&user, params["id"])
    json.NewDecoder(r.Body).Decode(&user)
    db.Save(&user)
    json.NewEncoder(w).Encode(Response{Status: true, Message: "User updated", Data: user})
}

// Handler untuk menghapus data User berdasarkan ID
func deleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user Users
    db.Delete(&user, params["id"])
    json.NewEncoder(w).Encode(Response{Status: true, Message: "User deleted"})
}
