package server

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

type requested struct { 
    ID string `json:"id"` 
    Name string `json:"name"` 
    Resource string `json:"resource"` 
    Event string `json:"event"` 
    Filter string `json:"filter"` 
    OrgID string `json:"orgId"` 
    CreatedBy string `json:"createdBy"` 
    AppID string `json:"appId"` 
    OwnedBy string `json:"ownedBy"` 
    Status string `json:"status"` 
    ActorID string `json:"actorId"` 
    Data struct { 
        ID string `json:"id"` 
        RoomID string `json:"roomId"` 
        PersonID string `json:"personId"` 
        PersonEmail string `json:"personEmail"` 
        Created time.Time `json:"created"` 
    } `json:"data"` 
}

func test(rw http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    log.Println(req.Form)
    //LOG: map[{"test": "that"}:[]]
    var t requested
    decoder := json.NewDecoder(req.Body)
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    defer req.Body.Close()
    log.Println(t.Data.ID)
    //LOG: that
}

func Start() {
    http.HandleFunc("/", test)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
