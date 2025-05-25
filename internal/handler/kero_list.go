package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *KeroHandler) ListKeros(c echo.Context) error {
	keros, err := h.lister.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, keros)
}
