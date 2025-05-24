package router

import (
	"github.com/kaerubo/kaeruashi/internal/handler"
	"github.com/kaerubo/kaeruashi/internal/usecase"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	creator usecase.KeroCreator,
	finder usecase.KeroByIDReader,
	lister usecase.KeroLister,
	updater usecase.KeroUpdater,
	deleter usecase.KeroDeleter,
) {
	h := handler.NewKeroHandler(
		creator,
		finder,
		lister,
		updater,
		deleter,
	)

	e.POST("/keros", h.CreateKero)
	e.GET("/keros", h.ListKeros)
	e.GET("/keros/:id", h.FindKeroByID)
	e.PUT("/keros/:id", h.UpdateKero)
	e.DELETE("/keros/:id", h.DeleteKero)
}
