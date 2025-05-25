package handler

import (
	"github.com/kaerubo/kaeruashi/internal/usecase"
)

type KeroHandler struct {
	creator usecase.KeroCreator
	finder  usecase.KeroByIDFinder
	lister  usecase.KeroLister
	updater usecase.KeroUpdater
	deleter usecase.KeroDeleter
}

func NewKeroHandler(
	creator usecase.KeroCreator,
	finder usecase.KeroByIDFinder,
	lister usecase.KeroLister,
	updater usecase.KeroUpdater,
	deleter usecase.KeroDeleter,
) *KeroHandler {
	return &KeroHandler{
		creator: creator,
		finder:  finder,
		lister:  lister,
		updater: updater,
		deleter: deleter,
	}
}
