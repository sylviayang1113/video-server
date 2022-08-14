package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func registerHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)

	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)

	router.POST("/userhome", userHomeHandler)

	router.POST("/api", apiHandler)

	router.GET("/videos/:vid-id", proxyHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router

}

func main()  {
	r := registerHandler()
	http.ListenAndServe(":8080", r)
}

