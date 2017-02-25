package batcave

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "time"    
    "bytes"
    "strings"
)

const (
    cleverBotURL = "https://www.cleverbot.com/getreply"
    apiKey2 = "CCCo6wW1Z5h8hyxmRx1vAXTC5BA"
    apiURL = "https://api.ciscospark.com/v1/messages"
    botID = "Y2lzY29zcGFyazovL3VzL1BFT1BMRS83NzU0YjYxYy04MjhlLTQ2MTItOWJjNy1lZmUyYWZhMDI3NGU"
    
    botRoomID = "Y2lzY29zcGFyazovL3VzL1JPT00vZTA3YWQyZGYtYjQwMC0zNzM0LThhNTAtMTJiYjJjMDEzMjgx"
    
    testRoom = "Y2lzY29zcGFyazovL3VzL1JPT00vNWZmNGM1ZTAtZjNhMS0xMWU2LWJhOWYtOTUwN2UyMTZkOTRj"
    testToken = "MzVhMzc3NzYtNDNjYS00MWZkLWJjODgtN2JjMWIwNzgzYTY4YjMwZjE4MGMtNGFj"
    botID2 = "ZDk0YjlmYWMtNDkyNC00OTExLTkzNGUtMTM4YjdhMDU3ZjU2ZWQ4MmUzNWEtYWMz"
    
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

type cleverbotReponse struct {
	Cs string `json:"cs"`
	InteractionCount string `json:"interaction_count"`
	Input string `json:"input"`
	Output string `json:"output"`
	ConversationID string `json:"conversation_id"`
}

func sendCleverbotQuery(msg, cs string) cleverbotReponse {
    reqURL := cleverBotURL + "?key=" + apiKey2
    reqURL += "&input=" + url.QueryEscape(msg)
    
    if cs != "" {
        reqURL += "&cs=" + cs
    }
    
    req, err := http.NewRequest("GET", reqURL, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    //log.Println("response Status:", resp.Status)
    
    var r cleverbotReponse
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&r)
    if err != nil {
        panic(err)
    }
    
    return r
}

func getMessage(id string) string {
    url := apiURL + "/" + id

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer " + testToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var t responsed
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    
    if strings.HasPrefix(t.Text, "PhilBot") {
        str := strings.SplitN(t.Text, " ", 2)
        if len(str) > 1 {
            return str[1]
        } 
        return "" 
    }
 
    return t.Text
}

// Start the request server with the specified database
func Start(taskDb TaskDatabase) {
    
    //cs := ""
    startedYet := false
    http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        var t requested
        decoder := json.NewDecoder(req.Body)
        err := decoder.Decode(&t)
        if err != nil {
            panic(err)
        }
        defer req.Body.Close()
        
        if t.Data.PersonID != botID && !startedYet {
            
            retrievedMessage := getMessage(t.Data.ID)
            log.Println("Retrieved message from Spark:", retrievedMessage)
            startedYet = true
            go startBots(retrievedMessage, t.Data.RoomID)
        }
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
    
}

// StartBots starts two bots talking to each other
func startBots(retrievedMessage, id string) {
    cs1 := ""
    cs2 := "" 
    lastMessage := retrievedMessage
    for true {
        msg := sendCleverbotQuery(lastMessage, cs1)
        txt, err := url.QueryUnescape(msg.Output)
        if err != nil {
            log.Fatal(err)
        }
        time.Sleep(time.Second)
        log.Println("Retrieved message from Clever1:", txt)
        cs1 = msg.Cs
        sendTestMessage(txt, id, testToken)
        
        msg2 := sendCleverbotQuery(txt, cs2)
        txt2, err2 := url.QueryUnescape(msg2.Output)
        if err2 != nil {
            log.Fatal(err2)
        }
        time.Sleep(time.Second)
        log.Println("Retrieved message from Clever2:", txt2)
        cs2 = msg2.Cs
        sendTestMessage(txt2, id, botID2)
        lastMessage = txt2
    }
}

func sendTestMessage(message, room, token string) {
    
    var jsonStr = []byte(`{"markdown":"` + message + `", "roomId":"` + room + `"}`)
                     
    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", "Bearer " + token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    //fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    //body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))
}
