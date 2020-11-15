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

func HandlePost(w http.ResponseWriter, r *http.Request) {

	var contoh ContohReq

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error"))
		return
	}

	err = json.Unmarshal(b, &contoh)
	if err != nil {
		w.Write([]byte("error unmarshal"))
		return
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
