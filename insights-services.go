package main

import (
    "net/http"
    "log"
    "encoding/json"
    "os"
    "io/ioutil"

    "github.com/gorilla/mux"
)

type Auth struct{
    User AuthUser         `json:"user"`
    Mechanism string      `json:"mechanism"`
}

type AuthUser struct{
    Id int                  `json:"id"`
    Username string         `json:"username"`
    Email string            `json:"email"`
    FirstName string        `json:"first_name"`
    LastName string         `json:"last_name"`
    AccountNumber string    `json:"account_number"`
    AddressString string    `json:"address_string"`
    IsActive bool           `json:"is_active"`
    IsOrgAdmin string       `json:"is_org_admin"`
    IsInternal string       `json:"is_internal"`
    Locale string           `json:"locale"`
    OrgId int               `json:"org_id"`
    DisplayName string      `json:"display_name"`
    Type string             `json:"type"`
}

type UserV2 struct{
    Users []UserV1       `json:"users"`
    UserCount int        `json:"userCount"`
}

type UserV1 struct{
    Id int               `json:"id"`
    Username string      `json:"username"`
    FirstName string     `json:"first_name"`
    LastName string      `json:"last_name"`
    Email string         `json:"email"`
    IsActive bool        `json:"is_active"`
    Locale string        `json:"locale"`
    IsOrgAdmin string    `json:"is_org_admin"`
    IsInternal string    `json:"is_internal"`
}

func status(w http.ResponseWriter, req *http.Request) {
    log.Println(req)
}

func auth(w http.ResponseWriter, req *http.Request) {
    log.Println(req)
    jsonFile, _ := os.Open("data/auth.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    users := Auth{}
    json.Unmarshal(byteValue, &users)
    json.NewEncoder(w).Encode(users)
}

func usersForAccountV1(w http.ResponseWriter, req *http.Request) {
    log.Println(req)
    jsonFile, _ := os.Open("data/usersv1.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    users := []UserV1{}
    json.Unmarshal(byteValue, &users)
    json.NewEncoder(w).Encode(users)
}

func usersForAccountV2(w http.ResponseWriter, req *http.Request) {
    log.Println(req)
    jsonFile, _ := os.Open("data/usersv2.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    users := UserV2{}
    json.Unmarshal(byteValue, &users)
    json.NewEncoder(w).Encode(users)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/insights-services/", status)
    router.HandleFunc("/insights-services/v1/accounts/{id}/users", usersForAccountV1)
    router.HandleFunc("/insights-services/v2/accounts/{id}/users", usersForAccountV2)
    router.HandleFunc("/insights-services/v1/auth", auth)

    log.Println("Listening on :8090")
    log.Fatal(http.ListenAndServe(":8090", router))
}
