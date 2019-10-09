package batcave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	cleverBotURL = "https://www.cleverbot.com/getreply"
	apiKey2      = "CCCo6wW1Z5h8hyxmRx1vAXTC5BA"
	apiURL       = "https://api.ciscospark.com/v1/messages"
	botID        = "Y2lzY29zcGFyazovL3VzL1BFT1BMRS83NzU0YjYxYy04MjhlLTQ2MTItOWJjNy1lZmUyYWZhMDI3NGU"

	botRoomID = "Y2lzY29zcGFyazovL3VzL1JPT00vZTA3YWQyZGYtYjQwMC0zNzM0LThhNTAtMTJiYjJjMDEzMjgx"

	testRoom  = "Y2lzY29zcGFyazovL3VzL1JPT00vNWZmNGM1ZTAtZjNhMS0xMWU2LWJhOWYtOTUwN2UyMTZkOTRj"
	testToken = "MzVhMzc3NzYtNDNjYS00MWZkLWJjODgtN2JjMWIwNzgzYTY4YjMwZjE4MGMtNGFj"
	botID2    = "ZDk0YjlmYWMtNDkyNC00OTExLTkzNGUtMTM4YjdhMDU3ZjU2ZWQ4MmUzNWEtYWMz"

	apiUser = "BbPSrA2qmTRlb3E9"
	apiKey  = "EH4IdmsgN8I0CoFEJ2Fpf5gZTyxX7CAb"
)

type requested struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Resource  string `json:"resource"`
	Event     string `json:"event"`
	Filter    string `json:"filter"`
	OrgID     string `json:"orgId"`
	CreatedBy string `json:"createdBy"`
	AppID     string `json:"appId"`
	OwnedBy   string `json:"ownedBy"`
	Status    string `json:"status"`
	ActorID   string `json:"actorId"`
	Data      struct {
		ID          string    `json:"id"`
		RoomID      string    `json:"roomId"`
		PersonID    string    `json:"personId"`
		PersonEmail string    `json:"personEmail"`
		Created     time.Time `json:"created"`
	} `json:"data"`
}

type responsed struct {
	ID          string    `json:"id"`
	RoomID      string    `json:"roomId"`
	PersonID    string    `json:"personId"`
	PersonEmail string    `json:"personEmail"`
	Text        string    `json:"text"`
	Created     time.Time `json:"created"`
}

func getMessage(id string) string {
	url := apiURL + "/" + id

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
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
func Start() {

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		var t requested
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", t)
		defer req.Body.Close()

	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func sendTestMessage(message, room, token string) {

	var jsonStr = []byte(`{"markdown":"` + message + `", "roomId":"` + room + `"}`)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
