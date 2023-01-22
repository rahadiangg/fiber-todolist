package todo

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ToDoHandler struct {
	service IToDoService
}

type IToDoHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewToDoHandler(Service IToDoService) IToDoHandler {
	return &ToDoHandler{
		service: Service,
	}
}

func (t *ToDoHandler) Create(c *fiber.Ctx) error {
	var data ToDoRequest

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("Error when parsing your request: %s", err.Error()),
		})
	}

	err = t.service.Create(data)
	if err != nil {

		var status int = fiber.StatusInternalServerError
		if err.Error() == "data not valid" {
			status = fiber.StatusBadRequest
		}

		return c.Status(status).JSON(&BasicResponse{
			Code:    status,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&BasicResponse{
		Code:    fiber.StatusOK,
		Message: "success create your data",
		Data:    data,
	})
}

func (t *ToDoHandler) GetAll(c *fiber.Ctx) error {

	all, err := t.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&BasicResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    all,
	})
}

func (t *ToDoHandler) GetById(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	data, err := t.service.GetById(uint32(paramsInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&BasicResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("success return your data with id %d", paramsInt),
		Data:    data,
	})
}

func (t *ToDoHandler) Update(c *fiber.Ctx) error {

	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	var data ToDoRequest

	err = c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("Error when parsing your request: %s", err.Error()),
		})
	}

	err = t.service.Update(uint32(paramsInt), data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&BasicResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("success update data %d", paramsInt),
	})
}

func (t *ToDoHandler) Delete(c *fiber.Ctx) error {

	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	err = t.service.Delete(uint32(paramsInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&BasicResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&BasicResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("success delete data %d", paramsInt),
	})
}
