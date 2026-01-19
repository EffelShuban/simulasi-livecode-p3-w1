package handler

import (
	"context"
	"net/http"
	"sim-livecodep3w1/internal/model"
	"sim-livecodep3w1/internal/service"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameHandler struct {
	svc service.GameService
}

func NewGameHandler(svc service.GameService)GameHandler{
	return GameHandler{
		svc: svc,
	}
}

func (h *GameHandler) FindByID(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
    if err != nil {
       return c.JSON(http.StatusBadRequest, echo.Map{"error": "bad id query param format"})
    }
	game, err := h.svc.FindByID(c.Request().Context(), id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, game)
}

func (h *GameHandler) Find(c echo.Context) error {
	games, err := h.svc.FindAll(c.Request().Context())
	if err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, games)
}

func (h *GameHandler) Create(c echo.Context) error {
	var req model.GameCreateRequest
		if err := c.Bind(&req); err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()}) 
	}

	game, err := h.svc.Create(c.Request().Context(), req)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"})
	}

	return c.JSON(http.StatusCreated, game)
}

func (h *GameHandler) Update(c echo.Context) error{
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
    if err != nil {
       return c.JSON(http.StatusBadRequest, echo.Map{"error": "bad id query param format"})
    }
	
	var req model.GameUpdateRequest
	if err := c.Bind(&req); err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"}) 
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10 *time.Second)
	defer cancel()
	err = h.svc.Update(ctx, req, id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "updated"})
}

func (h *GameHandler) Delete(c echo.Context) error{
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
    if err != nil {
       return c.JSON(http.StatusBadRequest, echo.Map{"error": "bad id query param format"})
    }

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10 *time.Second)
	defer cancel()
	err = h.svc.Delete(ctx, id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}