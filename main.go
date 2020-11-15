package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ContohReq struct {
	Nama string `json:"name"`
}

func readRequest(s interface{}, r *http.Request) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, s)
}

func HandleUserPost(w http.ResponseWriter, r *http.Request) {
	var contoh ContohReq

	err := readRequest(&contoh, r)
	if err != nil {
		w.Write([]byte("error"))
	}

	msg := fmt.Sprintf("Nama User : %s", contoh.Nama)
	w.Write([]byte(msg))
}

func HandlePost(w http.ResponseWriter, r *http.Request) {

	var contoh ContohReq

	err := readRequest(&contoh, r)
	if err != nil {
		w.Write([]byte("error"))
	}

	msg := fmt.Sprintf("Nama : %s", contoh.Nama)
	w.Write([]byte(msg))
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/new", HandlePost)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
