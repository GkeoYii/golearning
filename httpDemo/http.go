package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// normal
func normal() {
	resp, err := http.Get("http://localhost:1080/user/abc")
	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	log.Println("body = ", string(body))
}

//complex

func getByQuery() {

	//data := map[string]string{"username": "abc", "password": "student"}

	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:1080/users", nil)
	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	//request.Header.Add("Token", "123")
	//request.Header.Add("Accept", "application/json")

	q := request.URL.Query()
	q.Add("name", "abc")
	q.Add("role", "student")
	request.URL.RawQuery = q.Encode()

	log.Println(" url = ", request.URL.String())
	client.Timeout = 1 * time.Minute
	resp, err := client.Do(request)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	log.Println("body = ", string(body))
}

func postByForm() {

	request, err := http.NewRequest("POST", "http://localhost:1080/form",
		strings.NewReader("password=1234&username=abc"))

	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	client := &http.Client{}
	client.Timeout = 1 * time.Minute
	resp, err := client.Do(request)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	log.Println("body = ", string(body))
}

func postByJson() {

	data := map[string]string{"username": "abc", "password": "1234"}
	dateJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal("error := ", err)
		return
	}

	request, err := http.NewRequest("POST", "http://localhost:1080/json",
		bytes.NewReader(dateJson))
	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Timeout = 1 * time.Minute
	resp, err := client.Do(request)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error := ", err)
		return
	}
	log.Println("body = ", string(body))
}

func main() {
	//getByQuery()
	//postByForm()
	postByJson()
}
