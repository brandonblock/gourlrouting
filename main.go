package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", HomeHandler)

	//Posts
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	//Post, individual
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :8008")
	http.ListenAndServe(":8080", r)
}
