package batcave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	apiURL   = "https://api.ciscospark.com/v1"
	botToken = "OTg4NGEwOTMtYjkwOS00ZDM5LTg4NWEtM2Q4NmM0MGNlZTk3YjZlNjgwNDYtM2Mz_PF84_1eb65fdf-9643-417f-9974-ad72cae0e10f"
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
	url := apiURL + "/messages/" + id

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+botToken)
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

	if strings.HasPrefix(t.Text, "shizzlebot") {
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

		//msg := getMessage(t.Data.ID)
		//person := getPersonDetails(t.Data.PersonID)
		//resp := fmt.Sprintf("Hi <@personId:%s|%s>, this is what you sent me: '%s'", t.Data.PersonID, person.NickName, msg)
		//resp := fmt.Sprintf(card, t.Data.PersonID, person.NickName, msg)

		sendCard(t.Data.RoomID)
		//sendTestMessage(resp, t.Data.RoomID, t.Data.PersonID, botToken)

		defer req.Body.Close()

	})
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))

}

type person struct {
	NickName string `json:"nickName"`
}

func getPersonDetails(id string) person {

	req, err := http.NewRequest("GET", apiURL+"/people/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+botToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var p person

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&p)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return p
}

func sendCard(roomID string) {
	jsonStr := []byte(fmt.Sprintf(card, roomID))

	req, err := http.NewRequest("POST", apiURL+"/messages", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+botToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%+v\n", string(content))

	defer resp.Body.Close()
}

func sendTestMessage(message, room, personID, token string) {

	var jsonStr = []byte(`{"markdown":"` + message + `", "roomId":"` + room + `", "mentionedPeople": "` + personID + `"}`)

	req, err := http.NewRequest("POST", apiURL+"/messages", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%+v\n", string(content))

	defer resp.Body.Close()
}
