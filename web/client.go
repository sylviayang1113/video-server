package main

import (
	"log"
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	_"net/url"
	"encoding/json"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	switch b.Method {
		case http.MethodGet:
			req, _ := http.NewRequest("GET", b.Url, nil)
			req.Header = r.Header
			resp, err = httpClient.Do(req)

			if err != nil {
				log.Println(err)
				return
			}
			normalResponse(w, resp)
		case http.MethodPost:
			req, _ := http.NewRequest("Post", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
			resp, err = httpClient.Do(req)

			if err != nil {
				log.Println(err)
				return
			}
			normalResponse(w, resp)
		case http.MethodDelete:
			req, _ := http.NewRequest("Post", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
			resp, err = httpClient.Do(req)

			if err != nil {
				log.Println(err)
				return
			}
			normalResponse(w, resp)
		default:
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad api Request")
			return
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}


