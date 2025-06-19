package models

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"` 
}

var Users = []User{
    {Username: "admin", Password: "admin123", Role: "admin"},
    {Username: "user", Password: "user123", Role: "user"},
}
