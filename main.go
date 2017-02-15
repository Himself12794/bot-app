package main

import (
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "flag"
)

const (
    apiURL = "https://api.ciscospark.com/v1/messages"
    testRoom = "Y2lzY29zcGFyazovL3VzL1JPT00vNWZmNGM1ZTAtZjNhMS0xMWU2LWJhOWYtOTUwN2UyMTZkOTRj"
    testToken = "MzVhMzc3NzYtNDNjYS00MWZkLWJjODgtN2JjMWIwNzgzYTY4YjMwZjE4MGMtNGFj"
)

var message = flag.String("msg", "Test", "The message to send to Spark")

func main() {
    flag.Parse()
    sendTestMessage()
}

func sendTestMessage() {
    
    fmt.Println("URL:>", apiURL)

    var jsonStr = []byte(`{"markdown":"` + &message + `", "roomId":"` + testRoom + `"}`)
                     
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
