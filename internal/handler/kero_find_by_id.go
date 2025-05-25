package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *KeroHandler) FindKeroByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}

	kero, err := h.finder.FindByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if kero == nil {
		return echo.NewHTTPError(http.StatusNotFound, "kero not found")
	}

	return c.JSON(http.StatusOK, kero)
}
