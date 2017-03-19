package main

import (
	"log"
	"net/http"
	"os"

	"heroku-golang-echoServer/echo"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	wordStr := "helloworld"

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Echo": echo.Echo(wordStr)})
	}))
	log.Fatal(http.ListenAndServe(":"+port, api.MakeHandler()))
}
