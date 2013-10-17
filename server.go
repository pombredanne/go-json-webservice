package main

import (
	"net/http"
   "io/ioutil"
   "encoding/json"
)

type UserCreateRequest struct {
   Email string `json:"email"`
   Password string `json:"password"`
}

func baseHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"hello world\"}"))
}

func userCreateHandler(w http.ResponseWriter, req *http.Request) {
   body, _ := ioutil.ReadAll(req.Body)
   ucr := UserCreateRequest{}
   json.Unmarshal(body, &ucr)
   CreateUser(ucr)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"success\"}"))
}

func main() {
	servemux := http.NewServeMux()
	servemux.HandleFunc("/", baseHandler)
	servemux.HandleFunc("/user-create", userCreateHandler)

	http.ListenAndServeTLS("localhost:4321", "cert.pem", "key.pem", servemux)
}