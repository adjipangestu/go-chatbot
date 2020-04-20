package controllers
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)


type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize()  {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8000")
	handler := cors.Default().Handler(server.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}