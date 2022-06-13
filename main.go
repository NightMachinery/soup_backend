package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var _ = log.Println
var _ = fmt.Println

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/register/player", test)
	// http.HandleFunc("/headers", headers)

	http.ListenAndServe("localhost:9102", nil)
}

type register_req struct {
	GameID *string `json:"game_id"`

	Name *string `json:"name"`
}

type register_ans struct {
	Okay bool `json:"okay"`
}

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))

	var t register_req
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	log.Println("name: " + *t.Name)
	log.Println("game_id: " + *t.GameID)

	ans := register_ans{Okay: true};
	json.NewEncoder(rw).Encode(ans)
}
