package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type GenerateHash struct {
	Ok       bool   `json:"ok"`
	Msg      string `json:"msg"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
	Cost     int    `json:"cost"`
}

type CheckPassword struct {
	Ok       bool   `json:"ok"`
	Msg      string `json:"msg"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
	Cost     int    `json:"cost"`
}

type Err struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func checkErr(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func checkErrJson(w http.ResponseWriter, err error) bool {
	if err != nil {
		data := Err{
			false,
			err.Error(),
		}

		// send back some JSON
		json, err := json.Marshal(data)
		if checkErr(w, err) {
			return true
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		return true
	}
	return false
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/"+r.URL.Path[1:])
}

func apiGenerateHash(w http.ResponseWriter, r *http.Request) {
	// force the Method to be POST
	if strings.ToUpper(r.Method) == "POST" {
		r.Method = "POST"
	}

	err := r.ParseForm()
	if checkErr(w, err) {
		return
	}

	// get the password and cost
	password := r.FormValue("password")
	costStr := r.FormValue("cost")

	// check the password input
	if len(password) == 0 {
		// provide a password
		checkErrJson(w, errors.New("Provide a password"))
		return
	}

	// convert cost to an integer
	cost := 6
	if len(costStr) == 0 {
		cost = 6
	} else {
		cost, err = strconv.Atoi(costStr)
		if checkErrJson(w, err) {
			return
		}
	}

	// use bcrypt to hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if checkErrJson(w, err) {
		return
	}

	// create a datastructure to send back
	data := GenerateHash{
		true,
		"",
		password,
		string(hash),
		cost,
	}

	// send back some JSON
	json, err := json.Marshal(data)
	if checkErrJson(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func apiCheckPassword(w http.ResponseWriter, r *http.Request) {
	// get the password and cost
	password := r.FormValue("password")
	hash := r.FormValue("hash")

	// ToDo: check this hash looks syntactically correct

	data := CheckPassword{}

	// check this password is or isn't correct
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		// create a datastructure to send back
		data.Ok = true
		data.Password = password
		data.Hash = string(hash)
	} else {
		// see if the password/hash were different or something else entirely
		if err != bcrypt.ErrMismatchedHashAndPassword {
			if checkErrJson(w, err) {
				return
			}
		}

		// create a datastructure to send back
		data.Ok = false
		data.Msg = "Incorrect password for this hash"
		data.Password = password
		data.Hash = string(hash)
	}

	// get the cost of this hash
	cost, err := bcrypt.Cost([]byte(hash))
	if checkErrJson(w, err) {
		return
	}
	data.Cost = cost

	// send back some JSON
	json, err := json.Marshal(data)
	if checkErrJson(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	// the routers
	router := mux.NewRouter()

	// the handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/robots.txt", homeHandler)
	router.HandleFunc("/favicon.ico", homeHandler)
	router.PathPrefix("/s/").Handler(http.FileServer(http.Dir("static")))

	router.HandleFunc("/api/generate-hash.json", apiGenerateHash)
	router.HandleFunc("/api/check-password.json", apiCheckPassword)

	// the server
	http.Handle("/", router)

	// listen on
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
