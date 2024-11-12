package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// User represents user data
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Handler to return a hardcoded user
func getUser(w http.ResponseWriter, r *http.Request) {
    user := User{ID: 1, Name: "John Doe", Age: 30}

    // Set content type to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode user struct to JSON and write the response
    json.NewEncoder(w).Encode(user)
}

func main() {
    // Setup HTTP route
    http.HandleFunc("/user", getUser)

    // Start the server on port 8081
    fmt.Println("User service running on port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}