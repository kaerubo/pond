package handler

import (
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type KeroHandler struct {
	creator usecase.KeroCreator
}

func NewKeroHandler(creator usecase.KeroCreator) *KeroHandler {
	return &KeroHandler{creator: creator}
}

type createKeroRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *KeroHandler) CreateKero(c echo.Context) error {
	var req createKeroRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if req.Title == "" || req.Content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "title and content are required")
	}

	kero := &entity.Kero{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.creator.Create(c.Request().Context(), kero); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": kero.ID})
}
