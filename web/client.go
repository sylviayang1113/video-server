package main

import (
	"log"
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"interesting1113/config"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {

}


