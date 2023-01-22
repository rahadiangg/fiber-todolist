package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rahadiangg/fiber-todolist/config"
	"github.com/rahadiangg/fiber-todolist/todo"
	"log"
)

type Server struct {
	Handler todo.IToDoHandler
}

func NewServer(handler todo.IToDoHandler) *Server {
	return &Server{
		Handler: handler,
	}
}

func (s *Server) initServer() {
	app := fiber.New()

	initRoute(app, s)
	log.Fatal(app.Listen(":3000"))
}

func initRoute(app *fiber.App, server *Server) {

	app.Get("/todo", server.Handler.GetAll)
	app.Post("/todo", server.Handler.Create)
	app.Get("/todo/:id", server.Handler.GetById)
	app.Patch("/todo/:id", server.Handler.Update)
	app.Delete("/todo/:id", server.Handler.Delete)
}

func main() {

	db := config.NewDB()
	validate := validator.New()

	repo := todo.NewToDoRepository(db)
	service := todo.NewToDoSevice(repo, validate)
	handler := todo.NewToDoHandler(service)

	httpServer := NewServer(
		handler,
	)

	httpServer.initServer()
}
