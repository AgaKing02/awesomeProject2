package main

import (
	"log"
	"net/http"
)

var ip string
var port int

func init()  {
	ip="0.0.0.0"
	port=8000

}
func main()  {
	mux:=http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello,your name"))
	})
    err:=http.ListenAndServe("8000",mux)
	log.Logger.Fatal(err)

}
