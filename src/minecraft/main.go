package main

import (
    "log"
    "net/http"
    "minecraft/mcwhitelist"
    "encoding/json"
)

type Response struct {
    Status int `json:status`
    Whitelisted bool `json:isWhitelisted`
}

func handler(w http.ResponseWriter, r *http.Request) {
    isWhitelisted := mcwhitelist.NewMCWhitelist(r.URL.Path[len("/whitelist/"):])
    w.Header().Set("Content-Type", "application/json")
    data := &Response{200, isWhitelisted}
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/whitelist/", handler)
    log.Fatal(http.ListenAndServe(":8082", nil))
}