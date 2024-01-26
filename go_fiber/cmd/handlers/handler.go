package handlers

import (
	"go_fiber/bootstrap"
	"go_fiber/controller"

	"github.com/gofiber/fiber/v2"
)

// MyHandler represents the handler struct.
type MyHandler struct {
	Application bootstrap.Application
}

// NewServiceInitial creates a new instance of MyHandler.
func NewServiceInitial(app bootstrap.Application) MyHandler {
	return MyHandler{
		Application: app,
	}
}

// ServerInterfaceWrapper wraps controller interfaces.
type ServerInterfaceWrapper struct {
	CheckHandler controller.ICheckController
	BookHandler controller.IBookController
}

// CheckOperationId implements api.ServerInterface.
func (h *ServerInterfaceWrapper) CheckOperationId(c *fiber.Ctx) error {
	return h.BookHandler.Execute(c)
}

// Check implements api.ServerInterface.
func (h *ServerInterfaceWrapper) Check(c *fiber.Ctx) error {
	return h.CheckHandler.Execute(c)
}
