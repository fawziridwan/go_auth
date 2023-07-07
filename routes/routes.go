package controllers

import "go-auth/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/health-check", middleware.SetMiddlewareJson(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middleware.SetMiddlewareJson(s.Login)).Methods("POST")
}
