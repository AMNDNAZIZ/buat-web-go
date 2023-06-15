package app

import (
	"fmt"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server * Server) Iniitialize() {
	fmt.Println( a...: "Welcome to GoToko")

	server.Router = mux.NewRouter()
}

func (server * Server) Run(addr string){
	fmt.Printf( format: "Listening to port: %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))

}
	

func Run() {
	var server = Server{}

	server.Iniitialize()
	server.Run(addr: ":8080")
}