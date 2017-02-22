package batcave

import (
    "encoding/json"
    "log"
    "net/http"
    "time"    
    "fmt"
    "bytes"
    "io/ioutil"
    "strings"
    "github.com/CleverbotIO/go-cleverbot.io"
)

const (
    apiURL = "https://api.ciscospark.com/v1/messages"
    botID = "Y2lzY29zcGFyazovL3VzL1BFT1BMRS83NzU0YjYxYy04MjhlLTQ2MTItOWJjNy1lZmUyYWZhMDI3NGU"
    testRoom = "Y2lzY29zcGFyazovL3VzL1JPT00vNWZmNGM1ZTAtZjNhMS0xMWU2LWJhOWYtOTUwN2UyMTZkOTRj"
    testToken = "MzVhMzc3NzYtNDNjYS00MWZkLWJjODgtN2JjMWIwNzgzYTY4YjMwZjE4MGMtNGFj"
    apiUser = "BbPSrA2qmTRlb3E9"
    apiKey = "EH4IdmsgN8I0CoFEJ2Fpf5gZTyxX7CAb"
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

type responsed struct {
	ID string `json:"id"`
	RoomID string `json:"roomId"`
	PersonID string `json:"personId"`
	PersonEmail string `json:"personEmail"`
	Text string `json:"text"`
	Created time.Time `json:"created"`
}

func getMessage(id string) string {
    url := apiURL + "/" + id
    log.Println("URL:>", url)

    req, err := http.NewRequest("GET", url, nil)
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
    
    var t responsed
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    
    str := strings.SplitN(t.Text, " ", 2)
    
    if len(str) > 1 {
        return str[1]
    } 
 
    return ""
}

// Start the request server with the specified database
func Start(taskDb TaskDatabase) {
    
    theBot, err := cleverbot.New(apiUser, apiKey, "")
    if err != nil {
    
        http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
            
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
            if t.Data.PersonID != botID {
                msg := theBot.Ask(getMessage(t.Data.ID))
                sendTestMessage(msg, t.Data.RoomID)
            }
        })
        log.Fatal(http.ListenAndServe(":8080", nil))
    }
}

func sendTestMessage(message string, room string) {
    
    fmt.Println("URL:>", apiURL)
    fmt.Println(message)
    var jsonStr = []byte(`{"markdown":"` + message + `", "roomId":"` + room + `"}`)
                     
    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", "Bearer " + testToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
