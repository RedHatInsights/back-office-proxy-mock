package main

import (
    "net/http"
    "log"
    "encoding/json"
    "os"
    "io/ioutil"

    "github.com/gorilla/mux"
)

type User struct{
    Id int               `json:"id"`
    Username string      `json:"username"`
    FirstName string     `json:"first_name"`
    LastName string      `json:"last_name"`
    Email string         `json:"email"`
    IsActive bool        `json:"is_active"`
    Locale string        `json:"locale"`
}

func usersForAccount(w http.ResponseWriter, req *http.Request) {
    jsonFile, _ := os.Open("data/users.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    users := []User{}
    json.Unmarshal(byteValue, &users)
    json.NewEncoder(w).Encode(users)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/insights-services/v1/accounts/{id}/users", usersForAccount)

    log.Println("Listening on :8090")
    log.Fatal(http.ListenAndServe(":8090", router))
}
