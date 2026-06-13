package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/db/sqlc"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/repository"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/service"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		Repo: repo,
	}
}

type UserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req UserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid dob format"})
	}

	err = h.Repo.CreateUser(c.Context(), sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"name": req.Name,
		"dob":  req.DOB,
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.Repo.GetUser(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  service.CalculateAge(user.Dob),
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.Repo.ListUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var result []fiber.Map

	for _, user := range users {
		result = append(result, fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Format("2006-01-02"),
			"age":  service.CalculateAge(user.Dob),
		})
	}

	return c.JSON(result)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req UserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid dob"})
	}

	err = h.Repo.UpdateUser(c.Context(), sqlc.UpdateUserParams{
		ID:   int32(id),
		Name: req.Name,
		Dob:  dob,
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"id":   id,
		"name": req.Name,
		"dob":  req.DOB,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := h.Repo.DeleteUser(c.Context(), int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}
