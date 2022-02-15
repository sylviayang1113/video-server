package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteStream(w, "Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := ps.ByName("user_name")
	io.WriteString(w, uname)
}