package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/interesting1113/video_server/api/defs"
	"github.com/interesting1113/video_server/api/dbops"
	"github.com/interesting1113/video_server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(q, defs.ErrorRequestBodyParseFailed)
		return 
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse((w, defs.ErrorRequestBodyParseFailed)
		return
	}

	uname := p.ByName("username")
	log.Printf("Login url name: %s", uname)
	log.Printf("Login body name: %s", ubody.Username)
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	log.Printf("%s", ubody.Username)
	pwd, err := dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd: %s", pwd)
	log.Printf("Login body pwd: %s", ubody.Pwd)
	if err != nil || len(pwd) == 0  || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedIn{Success: true, SessionId: id}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !validateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}

	uname := p.ByName("username")
	u, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
	}

	ui := &defs.UserInfo{Id: u.Id}

	if resp, err := json.Marshal(ui); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}


func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}


func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}


func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}