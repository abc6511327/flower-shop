package handlers

import (
    "net/http"
)

// GetFlowers handles the request to get a list of flowers
func GetFlowers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Flowers fetched"))
}

// GetFlowerByID handles the request to get a single flower by its ID
func GetFlowerByID(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Flower details by ID"))
}
