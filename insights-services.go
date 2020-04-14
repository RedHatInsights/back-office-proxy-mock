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
    AccountNum string    `json:"account_number"`
}

type CountedUsers struct{
    UserCount int        `json:"userCount"`
    Users []User         `json:"users"`
}

func findUser(users []User, username string) int{
    for ind, user := range users {
        if user.Username == username {
            return ind
        }
    }
    return -1
}

func usersForAccount(w http.ResponseWriter, req *http.Request) {
    jsonFile, _ := os.Open("data/users.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    users := []User{}
    json.Unmarshal(byteValue, &users)
    var username = req.URL.Query().Get("username")
    if username != "" {
        res := findUser(users, username)
        if res == -1 {
            http.Error(w, "User not found.", 400)
            return
        } else {
            user := users[res]
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    json.NewEncoder(w).Encode(users)
}

func usersCountForAccount(w http.ResponseWriter, req *http.Request) {
    jsonFile, _ := os.Open("data/users.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    users := []User{}
    json.Unmarshal(byteValue, &users)

    var username = req.URL.Query().Get("username")
    if username != "" {
        res := findUser(users, username)
        if res == -1 {
            http.Error(w, "User not found.", 400)
            return
        } else {
            users = []User{users[res]}
        }
    }
    cu := CountedUsers{UserCount: len(users), Users: users}
    json.NewEncoder(w).Encode(cu)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/insights-services/v1/users", usersForAccount)
    router.HandleFunc("/insights-services/v1/accounts/{id}/users", usersForAccount)
    router.HandleFunc("/insights-services/v1/accounts/{id}/users", usersForAccount).Queries("username", "{name}")
    router.HandleFunc("/insights-services/v2/accounts/{id}/users", usersCountForAccount)
    router.HandleFunc("/insights-services/v2/accounts/{id}/users", usersCountForAccount).Queries("username", "{name}")

    log.Println("Listening on :8090")
    log.Fatal(http.ListenAndServe(":8090", router))
}
