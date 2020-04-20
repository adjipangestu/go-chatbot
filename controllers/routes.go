package controllers

import "wa-chattbot/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.GetData)).Methods("GET")
}