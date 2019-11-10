// no response 서버
// go run src\github.com\adonovan\gopl.io\ch1\fetchall\main.go http://localhost:8000/noresponse 으로 실험해보면 멈춤
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/noresponse", noresponse)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// noresonse never responds.
func noresponse(w http.ResponseWriter, r *http.Request) {
	for {
	}
}
