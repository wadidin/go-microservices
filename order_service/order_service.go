package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// User represents user data (same as in the user microservice)
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Order represents an order
type Order struct {
    OrderID int    `json:"order_id"`
    User    User   `json:"user"`
    Item    string `json:"item"`
    Amount  int    `json:"amount"`
}

// Handler to create an order
func createOrder(w http.ResponseWriter, r *http.Request) {
    // Call the user service API to get user data
    resp, err := http.Get("http://localhost:8081/user")
    if err != nil {
        http.Error(w, "Unable to contact user service", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Read the user service response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Failed to read response from user service", http.StatusInternalServerError)
        return
    }

    // Parse user JSON data
    var user User
    if err := json.Unmarshal(body, &user); err != nil {
        http.Error(w, "Failed to parse user data", http.StatusInternalServerError)
        return
    }

    // Create an order using user data
    order := Order{
        OrderID: 1,
        User:    user,
        Item:    "Laptop",
        Amount:  1200,
    }

    // Set content type to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode order struct to JSON and write the response
    json.NewEncoder(w).Encode(order)
}

func main() {
    // Setup HTTP route
    http.HandleFunc("/order", createOrder)

    // Start the server on port 8082
    fmt.Println("Order service running on port 8082")
    if err := http.ListenAndServe(":8082", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}