package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", index)
	log.Println("Started http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}


type RandomCredential struct {
	FirstName string `json:"first_name"`
	LastName string	 `json:"last_name"`
}

type JokeValue struct {
	Joke string `json:"joke"`
}

type Joke struct {
	Type  string `json:"type"`
	Value JokeValue `json:"value"`
}

func index(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	defer resp.Body.Close()

	var rc RandomCredential 

    body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.Unmarshal(body, &rc); err != nil {  
		json.NewEncoder(w).Encode(err)
		return
	}

	var j Joke

	url := fmt. Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=nerdy", rc.FirstName, rc.LastName)

	resp, err = http.Get(url)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return	
	}
	if err = json.Unmarshal(body, &j); err != nil {  
		json.NewEncoder(w).Encode(err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}
