package server

import (
    "encoding/json"
    "log"
    "net/http"
    "io/ioutil"
    "time"
)

const (
    apiURL = "https://api.ciscospark.com/v1/messages"
    testRoom = "Y2lzY29zcGFyazovL3VzL1JPT00vNWZmNGM1ZTAtZjNhMS0xMWU2LWJhOWYtOTUwN2UyMTZkOTRj"
    testToken = "MzVhMzc3NzYtNDNjYS00MWZkLWJjODgtN2JjMWIwNzgzYTY4YjMwZjE4MGMtNGFj"
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
    var t requested
    decoder := json.NewDecoder(req.Body)
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    defer req.Body.Close()
    //log.Printf("%+v\n", t)
    log.Println(t.ID)
    getMessage(t.Data.ID)
    //LOG: that
}

func getMessage(id string) string {
    
    log.Println("URL:>", apiURL + "/" + id)

    req, err := http.NewRequest("GET", apiURL, nil)
    req.Header.Set("Authorization", "Bearer " + testToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    log.Println("response Status:", resp.Status)
    log.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    log.Println("response Body:", string(body))
    
    return string(body)
}

func Start() {
    http.HandleFunc("/", test)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
