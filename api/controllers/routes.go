package controllers

import (
	"github.com/titoyudha/go_blog_api/api/middleware"
)

func (server *Server) initializeRoutes() {

	// mux := goji.NewMux()

	//Home Routes
	//With Goji
	// mux.HandleFunc(pat.Get("/api.v1.example/"), server.Home)
	//Without Goji / Plain
	server.Router.HandleFunc("/api.v1.example/", middleware.SetMiddleWareJSON(server.Home)).Methods("GET")

	//Login & SiginIn Routes
	// mux.HandleFunc(pat.Post("/api.v1.example/login"), server.Login)
	server.Router.HandleFunc("/api.v1.example/login", middleware.SetMiddleWareJSON(server.Login)).Methods("POST")

	//UserRoutes with goji
	// mux.HandleFunc(pat.Post("/api.v1.example/users"), server.CreateUser)
	// mux.HandleFunc(pat.Get("/api.v1.example/users"), server.GetUser)
	// mux.HandleFunc(pat.Get("/api.v1.example/users/{id}"), server.GetUserbyID)
	// mux.HandleFunc(pat.Put("/api.v1.example/users/{id}"), server.UpdateUser)
	// mux.HandleFunc(pat.Delete("/api.v1.example/users/{id}"), server.DeleteUser)

	// User routes withou goji
	server.Router.HandleFunc("/api.v1.example/user", middleware.SetMiddleWareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/api.v1.example/users", middleware.SetMiddleWareJSON(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/api.v1.example/users/{id}", middleware.SetMiddleWareJSON(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/api.v1.example/users/{id}", middleware.SetMiddleWareJSON(middleware.SetMiddlewareAuth(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/api.v1.example/users/{id}", middleware.SetMiddlewareAuth(server.DeleteUser)).Methods("DELETE")

	//PostRoutes with goji
	// mux.HandleFunc(pat.Post("/api.v1.example/post"), server.CreatePost)
	// mux.HandleFunc(pat.Get("/api.v1.example/post"), server.GetPosts)
	// mux.HandleFunc(pat.Get("/api.v1.example/post/{id}"), server.GetPost)
	// mux.HandleFunc(pat.Put("/api.v1.example/post/{id}"), server.UpdatePost)
	// mux.HandleFunc(pat.Delete("/api.v1.example/post/{id}"), server.DeletePost)

	// Post routes without goji
	server.Router.HandleFunc("/api.v1.example/post", middleware.SetMiddleWareJSON(server.CreatePost)).Methods("POST")
	server.Router.HandleFunc("/api.v1.example/posts", middleware.SetMiddleWareJSON(server.GetPosts)).Methods("GET")
	server.Router.HandleFunc("/api.v1.example/post/{id}", middleware.SetMiddleWareJSON(server.GetPost)).Methods("GET")
	server.Router.HandleFunc("/api.v1.example/post/{id}", middleware.SetMiddleWareJSON(middleware.SetMiddlewareAuth(server.UpdatePost))).Methods("PUT")
	server.Router.HandleFunc("/api.v1.example/post/{id}", middleware.SetMiddlewareAuth(server.DeletePost)).Methods("DELETE")

}
