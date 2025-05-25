package handler

import (
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type updateKeroRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *KeroHandler) UpdateKero(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}

	var req updateKeroRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if req.Title == "" || req.Content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "title and content are required")
	}

	kero := &entity.Kero{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.updater.Update(c.Request().Context(), kero); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
