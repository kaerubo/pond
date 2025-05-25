package router

import (
	"github.com/kaerubo/kaeruashi/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	keroHandler *handler.KeroHandler,
) {
	e.POST("/keros", keroHandler.CreateKero)
	e.GET("/keros", keroHandler.ListKeros)
	e.GET("/keros/:id", keroHandler.FindKeroByID)
	e.PUT("/keros/:id", keroHandler.UpdateKero)
	e.DELETE("/keros/:id", keroHandler.DeleteKero)
}
