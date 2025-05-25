package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *KeroHandler) DeleteKero(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}

	if err := h.deleter.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
